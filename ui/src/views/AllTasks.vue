<template>
  <v-skeleton-loader type="list-two-line" />
</template>

<script>
export default {
  mounted() {
    this.$store.commit('ui/setToolbar', { title: 'Todas as Tarefas' });
    document.title = 'Todas as Tarefas - Controle de Tarefas';
  },
};
</script>
