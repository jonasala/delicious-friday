<template>
  <v-app-bar
    app
    :color="toolbar.color || 'primary'"
    :dark="!toolbar.search"
    :light="toolbar.search"
  >
    <v-app-bar-nav-icon @click="drawer = !drawer" v-if="!toolbar.search" />
    <v-btn icon v-if="toolbar.search" @click="closeSearch()" class="mr-5">
      <v-icon>mdi-arrow-left</v-icon>
    </v-btn>
    <v-toolbar-title v-if="!toolbar.search">{{ toolbar.title }}</v-toolbar-title>
    <v-text-field
      v-if="toolbar.search"
      solo
      flat
      single-line
      hide-details
      clearable
      placeholder="Digite qualquer coisa a respeito da tarefa"
      v-model="searchText"
    />
    <v-spacer v-if="!toolbar.search"></v-spacer>
    <v-btn icon @click="initSearch()" v-if="!toolbar.search">
      <v-icon>mdi-magnify</v-icon>
    </v-btn>
  </v-app-bar>
</template>

<script>
import { mapState } from 'vuex';

export default {
  data() {
    return {
      prev: null,
    };
  },
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
    searchText: {
      get() {
        return this.$store.state.ui.toolbar.searchText;
      },
      set(text) {
        this.$store.commit('ui/setSearchText', text);
      },
    },
  },
  methods: {
    initSearch() {
      this.prev = this.toolbar;
      this.$store.commit('ui/setToolbar', {
        search: true,
        color: 'default',
      });
    },
    closeSearch() {
      this.$store.commit('ui/setToolbar', this.prev);
    },
  },
};
</script>
