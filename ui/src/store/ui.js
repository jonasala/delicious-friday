export default {
  namespaced: true,
  state: {
    snackbar: {
      open: false,
      color: '',
      text: '',
    },
  },
  mutations: {
    setSnackbar(state, options) {
      state.snackbar.text = options.text || '';
      state.snackbar.open = options.open || true;
      state.snackbar.color = options.color || 'error';
    },
    closeSnackbar(state) {
      state.snackbar.open = false;
    },
  },
  actions: {},
};
