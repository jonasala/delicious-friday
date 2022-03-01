export default {
  namespaced: true,
  actions: {
    async list(_, payload) {
      try {
        const response = await window.securedAxios.get(
          `/api/restricted/tasks/?search=${payload.search}&page=${payload.page}&items_per_page=${
            payload.itemsPerPage
          }&all=${payload.all || ''}`,
        );
        return response.data;
      } catch (err) {
        return Promise.reject(err);
      }
    },
    async create(_, payload) {
      try {
        const response = await window.securedAxios.post('/api/restricted/tasks', payload);
        return response.data;
      } catch (err) {
        return Promise.reject(err);
      }
    },
    async delete(_, id) {
      try {
        const response = await window.securedAxios.delete(`/api/restricted/tasks/${id}`);
        return response.data;
      } catch (err) {
        return Promise.reject(err);
      }
    },
  },
};
