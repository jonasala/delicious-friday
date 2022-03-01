<template>
  <v-container fluid>
    <v-form @submit.prevent="save()">
      <v-checkbox label="Tarefa sem ordem de serviço" v-model="withoutWO" :disabled="loading" />
      <v-autocomplete
        filled
        label="Ordem de serviço"
        v-model="task.work_order"
        :disabled="withoutWO || loading || loadingWO"
        :loading="loadingWO"
        :items="workOrders"
        item-text="number"
        item-value="id"
        append-outer-icon="mdi-file-eye-outline"
        @click:append-outer="woDetails()"
        no-data-text="Nenhuma ordem de serviço encontrada"
        :search-input.sync="searchWO"
        clearable
        @blur="$v.task.work_order && $v.task.work_order.$touch()"
        @input="$v.task.work_order && $v.task.work_order.$touch()"
        :error-messages="woErrors"
      >
        <template v-slot:append-item>
          <v-list-item>
            <v-list-item-content>
              <v-btn block depressed @click="dialogWO = true">
                <v-icon class="mr-2">mdi-plus</v-icon>
                Nova ordem de serviço
              </v-btn>
            </v-list-item-content>
          </v-list-item>
        </template>
      </v-autocomplete>
      <v-textarea
        filled
        auto-grow
        label="Descrição"
        v-model="task.description"
        :disabled="loading"
        @blur="$v.task.description && $v.task.description.$touch()"
        @input="$v.task.description && $v.task.description.$touch()"
        :error-messages="descriptionErrors"
      />
      <v-btn block color="primary" type="submit" large :loading="loading">
        <v-icon class="mr-2">mdi-check</v-icon>
        Salvar
      </v-btn>
    </v-form>
    <v-dialog fullscreen hide-overlay transition="dialog-bottom-transition" v-model="dialogWO">
      <work-order-form @close="dialogWO = false" @create="createdWO" />
    </v-dialog>
    <v-dialog
      fullscreen
      hide-overlay
      transition="dialog-bottom-transition"
      v-model="dialogWODetail"
    >
      <work-order-details @close="dialogWODetail = false" v-if="selectedWO" :wo="selectedWO" />
    </v-dialog>
  </v-container>
</template>

<script>
import _ from 'lodash';
import { required } from 'vuelidate/lib/validators';
import WorkOrderForm from '../components/WorkOrderForm.vue';
import WorkOrderDetails from '../components/WorkOrderDetails.vue';

export default {
  components: { WorkOrderForm, WorkOrderDetails },
  data() {
    return {
      task: {
        work_order: '',
        description: '',
      },
      withoutWO: false,
      loading: false,
      loadingWO: false,
      workOrders: [],
      dialogWODetail: false,
      dialogWO: false,
      searchWO: null,
      debouncedSearch: _.debounce((vm, val) => {
        vm.loadWorkOrders(val);
      }, 500),
      selectedWO: null,
    };
  },
  validations() {
    if (this.withoutWO) {
      return { task: { description: { required } } };
    }
    return { task: { work_order: { required } } };
  },
  computed: {
    woErrors() {
      const errors = [];
      if (this.withoutWO || !this.$v.task.work_order.$dirty) {
        return errors;
      }

      if (!this.$v.task.work_order.required) {
        errors.push('Informe a ordem de serviço.');
      }
      return errors;
    },
    descriptionErrors() {
      const errors = [];
      if (!this.withoutWO || !this.$v.task.description.$dirty) {
        return errors;
      }

      if (!this.$v.task.description.required) {
        errors.push('A descrição é obrigatória quando não é informada a ordem de serviço.');
      }
      return errors;
    },
  },
  mounted() {
    this.$store.commit('ui/setToolbar', { title: 'Nova Tarefa' });
    document.title = 'Nova Tarefa - Controle de Tarefas';

    this.loadWorkOrders();
  },
  methods: {
    async save() {
      this.$v.$touch();
      if (this.$v.$invalid) {
        return;
      }
      try {
        this.loading = true;
        await this.$store.dispatch('tasks/create', this.task);
        this.$store.commit('ui/setSnackbar', {
          text: 'Tarefa criada com sucesso',
          color: 'success',
        });
        this.$router.back();
      } catch (err) {
        console.log(err);
      } finally {
        this.loading = false;
      }
    },
    async loadWorkOrders(prefix) {
      try {
        this.loadingWO = true;
        this.workOrders = await this.$store.dispatch('workOrders/loadWorkOrders', prefix);
      } catch (err) {
        console.log(err);
      } finally {
        this.loadingWO = false;
      }
    },
    async woDetails() {
      if (!this.task.work_order) {
        return;
      }
      try {
        this.loadingWO = true;
        this.selectedWO = await this.$store.dispatch(
          'workOrders/getWorkOrder',
          this.task.work_order,
        );
        this.dialogWODetail = true;
      } catch (err) {
        console.log(err);
      } finally {
        this.loadingWO = false;
      }
    },
    createdWO(wo) {
      this.workOrders = [wo];
      this.task.work_order = wo.id;
      this.dialogWO = false;
    },
  },
  watch: {
    withoutWO(val) {
      if (val) {
        this.$set(this.task, 'work_order', '');
      }
    },
    searchWO(val) {
      this.debouncedSearch(this, val);
    },
    task() {
      this.$v.$reset();
    },
  },
};
</script>
