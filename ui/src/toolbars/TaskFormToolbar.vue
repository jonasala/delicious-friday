<template>
  <v-app-bar app color="primary" dark>
    <v-btn icon @click="$router.back()">
      <v-icon>mdi-arrow-left</v-icon>
    </v-btn>
    <v-toolbar-title>{{ toolbar.title }}</v-toolbar-title>
  </v-app-bar>
</template>

<script>
import { mapState } from 'vuex';

export default {
  computed: {
    ...mapState({
      toolbar: (state) => state.ui.toolbar,
    }),
  },
};
</script>
