<template>
  <v-container>
    <v-form @submit.prevent="save()">
      <v-checkbox label="Tarefa sem ordem de serviço" v-model="withoutWO" :disabled="loading" />
      <v-autocomplete
        filled
        label="Ordem de serviço"
        v-model="task.work_order"
        :disabled="withoutWO || loading"
        :loading="loadingWO"
        :items="workOrders"
        append-outer-icon="mdi-file-eye-outline"
        @click:append-outer="woDetails()"
        no-data-text="Nenhuma ordem de serviço encontrada"
        :search-input.sync="searchWO"
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
      />
      <v-btn block color="primary" type="submit" large :loading="loading">
        <v-icon class="mr-2">mdi-check</v-icon>
        Salvar
      </v-btn>
    </v-form>
    <v-dialog fullscreen hide-overlay transition="dialog-bottom-transition" v-model="dialogWO">
      <work-order-form @close="dialogWO = false" />
    </v-dialog>
  </v-container>
</template>

<script>
import _ from 'lodash';
import WorkOrderForm from '../components/WorkOrderForm.vue';

export default {
  components: { WorkOrderForm },
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
      dialogWO: false,
      searchWO: null,
      debouncedSearch: _.debounce((vm, val) => {
        vm.loadWorkOrders(val);
      }, 500),
    };
  },
  mounted() {
    this.$store.commit('ui/setToolbar', { title: 'Nova Tarefa' });
    document.title = 'Nova Tarefa - Controle de Tarefas';

    this.loadWorkOrders();
  },
  methods: {
    save() {
      return false;
    },
    async loadWorkOrders(prefix) {
      try {
        this.loadingWO = true;
        const wo = await this.$store.dispatch('workOrders/loadWorkOrders', prefix);
        this.workOrders = wo.map((i) => i.number);
      } catch (err) {
        console.log(err);
      } finally {
        this.loadingWO = false;
      }
    },
    woDetails() {
      //
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
  },
};
</script>
