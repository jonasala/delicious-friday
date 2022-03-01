import axios from 'axios';
import store from '../store';

const pAxios = axios.create({
  baseURL: process.env.VUE_APP_SERVER_URL,
});

pAxios.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response === undefined) {
      store.commit('ui/setSnackbar', {
        text: 'Não foi possível estabelecer conexão com o servidor. Verifique sua conexão com a internet.',
      });
    } else if (error.response.status === 500) {
      store.commit('ui/setSnackbar', {
        text: 'Problemas de conexão com o servidor. Tente novamente mais tarde.',
      });
    }
    return Promise.reject(error);
  },
);

const sAxios = axios.create({
  baseURL: process.env.VUE_APP_SERVER_URL,
});

sAxios.interceptors.request.use(
  (cfg) => {
    const config = cfg;
    config.headers.Authorization = `Bearer ${localStorage.getItem('authToken')}`;
    return config;
  },
  (error) => Promise.reject(error),
);

sAxios.interceptors.response.use(
  (response) => response,
  async (error) => {
    if (error.response === undefined) {
      store.commit('ui/setSnackbar', {
        text: 'Não foi possível estabelecer conexão com o servidor. Verifique sua conexão com a internet.',
      });
    } else if (error.response.status === 500) {
      store.commit('ui/setSnackbar', {
        text: 'Problemas de conexão com o servidor. Tente novamente mais tarde.',
      });
    } else if (error.response.status === 401) {
      store.dispatch('users/logout');
      return Promise.reject(error);
    }
    return Promise.reject(error);
  },
);

window.publicAxios = pAxios;
window.securedAxios = sAxios;
