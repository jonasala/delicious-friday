import Vue from 'vue';
import Vuex from 'vuex';

import ui from './ui';
import users from './users';

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    ui,
    users,
  },
});
