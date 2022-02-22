<template>
  <div class="home">
    <HelloWorld msg="Welcome to Your Vue.js App" />
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue';

export default {
  name: 'HomeView',
  components: {
    HelloWorld,
  },
  mounted() {
    this.$store.commit('ui/setToolbar', { title: 'Controle de Tarefas' });
    document.title = 'Home - Controle de Tarefas';
  },
};
</script>
