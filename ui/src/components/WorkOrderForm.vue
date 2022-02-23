<template>
  <v-card>
    <v-form @submit.prevent="save()">
      <v-toolbar dark color="primary">
        <v-btn icon @click="$emit('close')" dark :disabled="loading">
          <v-icon>mdi-close</v-icon>
        </v-btn>
        <v-toolbar-title>Criar Ordem de Serviço</v-toolbar-title>
        <v-spacer />
        <v-toolbar-items>
          <v-btn dark text type="submit" :loading="loading">
            <v-icon class="mr-2">mdi-check</v-icon>SALVAR
          </v-btn>
        </v-toolbar-items>
      </v-toolbar>
      <v-card-text>
        <v-text-field filled label="Número" v-model="wo.number" :disabled="loading" />
        <div class="text-center">
          <v-btn color="accent" @click="takePicture()">
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
          @change="preview"
        />
      </v-card-text>
    </v-form>
  </v-card>
</template>

<script>
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
  methods: {
    save() {
      //
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
