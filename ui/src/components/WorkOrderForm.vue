<template>
  <v-card>
    <v-form @submit.prevent="save()">
      <v-toolbar dark color="primary">
        <v-btn icon @click="close()" dark :disabled="loading">
          <v-icon>mdi-close</v-icon>
        </v-btn>
        <v-toolbar-title>Criar Ordem de Serviço</v-toolbar-title>
        <v-spacer />
        <v-toolbar-items>
          <v-btn dark text type="submit" :loading="loading" :disabled="$v.$invalid">
            <v-icon class="mr-2">mdi-check</v-icon>SALVAR
          </v-btn>
        </v-toolbar-items>
      </v-toolbar>
      <v-card-text>
        <v-text-field
          filled
          label="Número"
          v-model="wo.number"
          :disabled="loading"
          @blur="$v.wo.number.$touch()"
          @input="$v.wo.number.$touch()"
          :error-messages="numberErrors"
        />
        <div class="text-center">
          <v-btn color="accent" @click="takePicture()" :disabled="loading">
            <v-icon class="mr-2">mdi-camera</v-icon>TIRAR FOTO
          </v-btn>
        </div>

        <div class="preview mt-5">
          <panZoom v-if="wo.file" class="text-center">
            <img :src="wo.file" alt="Foto da ordem de serviço" />
          </panZoom>
        </div>

        <v-file-input
          accept="image/*"
          capture="camera"
          ref="fileInput"
          class="d-none"
          v-model="wo.fileField"
          @change="preview"
        />
      </v-card-text>
    </v-form>
  </v-card>
</template>

<script>
import { required } from 'vuelidate/lib/validators';

export default {
  data() {
    return {
      loading: false,
      wo: {
        number: '',
        file: '',
      },
    };
  },
  validations: {
    wo: {
      number: { required },
      file: { required },
    },
  },
  computed: {
    numberErrors() {
      const errors = [];
      if (!this.$v.wo.number.$dirty) {
        return errors;
      }
      if (!this.$v.wo.number.required) {
        errors.push('O número da ordem de serviço é obrigatório');
      }
      return errors;
    },
  },
  methods: {
    async save() {
      try {
        this.$v.$touch();
        if (!this.$v.$invalid) {
          this.loading = true;
          const formData = new FormData();
          formData.append('number', this.wo.number);
          formData.append('document', this.wo.fileField);
          const created = await this.$store.dispatch('workOrders/createWorkOrder', formData);
          this.$emit('create', created);
        }
      } catch (err) {
        console.log(err);
      } finally {
        this.loading = false;
      }
    },
    takePicture() {
      this.$refs.fileInput.$el.querySelector('input').click();
    },
    preview(file) {
      if (file) {
        const reader = new FileReader();
        reader.onload = (e) => {
          this.wo.file = e.target.result;
        };
        reader.readAsDataURL(file);
      }
    },
    close() {
      this.wo = {
        number: '',
        file: '',
      };
      this.$v.$reset();
      this.$emit('close');
    },
  },
};
</script>

<style scoped>
.preview {
  width: 100%;
  max-height: 60vh;
  height: 60vh;
  overflow: hidden;
}
</style>
