<template>
  <v-app-bar app :color="toolbar.color || 'primary'" dark>
    <v-app-bar-nav-icon @click="drawer = !drawer" />
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
    drawer: {
      get() {
        return this.$store.state.ui.drawer;
      },
      set(open) {
        this.$store.commit('ui/setDrawer', open);
      },
    },
  },
};
</script>
