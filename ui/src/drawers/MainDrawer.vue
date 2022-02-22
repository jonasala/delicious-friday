<template>
  <v-navigation-drawer v-model="drawer" app>
    <v-list two-line>
      <v-list-item>
        <v-list-item-avatar color="grey">
          <v-icon>mdi-account</v-icon>
        </v-list-item-avatar>
        <v-list-item-content>
          <v-list-item-title>Jonas Sala</v-list-item-title>
          <v-list-item-subtitle>jonas.sala</v-list-item-subtitle>
        </v-list-item-content>
      </v-list-item>
    </v-list>
    <v-divider />
    <v-list nav dense>
      <v-list-item-group v-model="selectedItem" color="primary">
        <v-list-item value="my-tasks">
          <v-list-item-icon>
            <v-icon>mdi-format-list-bulleted</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>Minhas tarefas</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-list-item value="all-tasks">
          <v-list-item-icon>
            <v-icon>mdi-format-list-bulleted-type</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>Todas as tarefas</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list-item-group>
      <v-divider />
      <v-list-item @click="$store.dispatch('users/logout')">
        <v-list-item-icon>
          <v-icon>mdi-logout</v-icon>
        </v-list-item-icon>
        <v-list-item-content>
          <v-list-item-title>Sair</v-list-item-title>
        </v-list-item-content>
      </v-list-item>
    </v-list>
  </v-navigation-drawer>
</template>

<script>
export default {
  data() {
    return {
      selectedItem: this.$route.name,
    };
  },
  computed: {
    drawer: {
      get() {
        return this.$store.state.ui.drawer;
      },
      set(open) {
        this.$store.commit('ui/setDrawer', open);
      },
    },
  },
  watch: {
    selectedItem(to) {
      this.$router.push({ name: to });
    },
  },
};
</script>
