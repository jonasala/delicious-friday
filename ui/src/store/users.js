export default {
  namespaced: true,
  state: {
    activeUser: JSON.parse(localStorage.getItem('activeUser')),
  },
  mutations: {
    setActiveUser(state, activeUser) {
      localStorage.setItem('activeUser', JSON.stringify(activeUser));
      state.activeUser = activeUser;
    },
  },
  actions: {
    async login({ commit }, payload) {
      try {
        const response = await window.publicAxios.post('/api/auth/login', payload);
        localStorage.setItem('authToken', response.data.token);
        commit('setActiveUser', response.data);
        return Promise.resolve(true);
      } catch (err) {
        return Promise.reject(err);
      }
    },
    logout({ commit }) {
      commit('setActiveUser', null);
      localStorage.removeItem('authToken');
      localStorage.removeItem('activeUser');
      window.location.reload();
    },
  },
};
