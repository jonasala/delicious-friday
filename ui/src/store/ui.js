export default {
  namespaced: true,
  state: {
    snackbar: {
      open: false,
      color: '',
      text: '',
    },
    toolbar: {
      title: '',
      search: false,
    },
    drawer: null,
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
    setToolbar(state, toolbar) {
      state.toolbar = toolbar;
    },
    setDrawer(state, open) {
      state.drawer = open;
    },
  },
  actions: {},
};
