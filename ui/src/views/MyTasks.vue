<template>
  <v-container fill-height>
    <v-btn color="accent" dark fixed bottom right fab :to="{ name: 'create-task' }">
      <v-icon>mdi-plus</v-icon>
    </v-btn>
  </v-container>
</template>

<script>
export default {
  mounted() {
    this.$store.commit('ui/setToolbar', { title: 'Minhas Tarefas' });
    document.title = 'Minhas Tarefas - Controle de Tarefas';
  },
};
</script>
