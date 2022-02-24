<template>
  <v-container fluid class="px-0">
    <task-list all="1" />
    <v-btn color="accent" dark fixed bottom right fab :to="{ name: 'create-task' }">
      <v-icon>mdi-plus</v-icon>
    </v-btn>
  </v-container>
</template>

<script>
import TaskList from '../components/TaskList.vue';

export default {
  components: { TaskList },
  mounted() {
    this.$store.commit('ui/setToolbar', { title: 'Todas as Tarefas' });
    document.title = 'Todas as Tarefas - Controle de Tarefas';
  },
};
</script>
