import Vue from 'vue';
import Vuelidate from 'vuelidate';
import panZoom from 'vue-panzoom';
import App from './App.vue';
import router from './router';
import store from './store';
import vuetify from './plugins/vuetify';
import './lib/axios';

Vue.use(Vuelidate);
Vue.use(panZoom);
Vue.config.productionTip = false;

new Vue({
  router,
  store,
  vuetify,
  render: (h) => h(App),
}).$mount('#app');
