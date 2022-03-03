<template>
  <v-card>
    <v-toolbar dark color="primary">
      <v-btn icon @click="$emit('close')" dark>
        <v-icon>mdi-close</v-icon>
      </v-btn>
      <v-toolbar-title>Ordem {{ wo.number }}</v-toolbar-title>
    </v-toolbar>
    <v-card-text class="px-0">
      <v-list two-line>
        <v-list-item>
          <v-list-item-content>
            <v-list-item-subtitle>Criado por</v-list-item-subtitle>
            <v-list-item-title>{{ wo.created_by }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-list-item>
          <v-list-item-content>
            <v-list-item-subtitle>Criado em</v-list-item-subtitle>
            <v-list-item-title>{{ format(wo.created_at) }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
      <div class="preview">
        <panZoom v-if="wo.file_url" class="text-center">
          <img :src="absoluteURL" alt="Foto da ordem de serviço" />
        </panZoom>
      </div>
    </v-card-text>
  </v-card>
</template>

<script>
import { DateTime } from 'luxon';

export default {
  props: ['wo'],
  computed: {
    absoluteURL() {
      return `${process.env.VUE_APP_SERVER_URL}/${this.wo.file_url}`;
    },
  },
  methods: {
    format(date) {
      return DateTime.fromISO(date).toLocaleString(DateTime.DATETIME_MED);
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
