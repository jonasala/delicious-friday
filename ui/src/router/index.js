import Vue from 'vue';
import VueRouter from 'vue-router';
import MainToolbar from '../toolbars/MainToolbar.vue';
import TaskFormToolbar from '../toolbars/TaskFormToolbar.vue';
import MainDrawer from '../drawers/MainDrawer.vue';
import MyTasks from '../views/MyTasks.vue';
import LoginView from '../views/LoginView.vue';
import AllTasks from '../views/AllTasks.vue';
import TaskForm from '../views/TaskForm.vue';

Vue.use(VueRouter);

const routes = [
  {
    path: '/minhas-tarefas',
    name: 'my-tasks',
    components: {
      drawer: MainDrawer,
      toolbar: MainToolbar,
      default: MyTasks,
    },
  },
  {
    path: '/todas-tarefas',
    name: 'all-tasks',
    components: {
      drawer: MainDrawer,
      toolbar: MainToolbar,
      default: AllTasks,
    },
  },
  {
    path: '/nova-tarefa',
    name: 'create-task',
    components: {
      drawer: null,
      toolbar: TaskFormToolbar,
      default: TaskForm,
    },
  },
  {
    path: '/login',
    name: 'login',
    components: {
      toolbar: null,
      drawer: null,
      default: LoginView,
    },
  },
  { path: '*', redirect: { name: 'my-tasks' } },
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

router.beforeEach((to, from, next) => {
  if (to.name === 'login') {
    return next();
  }
  if (!localStorage.getItem('authToken')) {
    return next({ name: 'login' });
  }
  return next();
});

export default router;
