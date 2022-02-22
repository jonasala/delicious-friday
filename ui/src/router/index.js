import Vue from 'vue';
import VueRouter from 'vue-router';
import HomeView from '../views/HomeView.vue';
import LoginView from '../views/LoginView.vue';
import MainToolbar from '../toolbars/MainToolbar.vue';
import MainDrawer from '../drawers/MainDrawer.vue';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'home',
    components: {
      drawer: MainDrawer,
      toolbar: MainToolbar,
      default: HomeView,
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
  { path: '*', redirect: { name: 'home' } },
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
