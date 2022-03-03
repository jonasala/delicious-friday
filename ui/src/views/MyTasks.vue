<template>
  <v-container fluid class="px-0">
    <task-list />
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
    this.$store.commit('ui/setToolbar', { title: 'Minhas Tarefas' });
    document.title = 'Minhas Tarefas - Controle de Tarefas';
  },
};
</script>
