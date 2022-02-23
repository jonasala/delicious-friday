import Vue from 'vue';
import Vuex from 'vuex';

import ui from './ui';
import users from './users';
import workOrders from './work-orders';

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    ui,
    users,
    workOrders,
  },
});
