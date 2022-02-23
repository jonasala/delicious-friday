export default {
  namespaced: true,
  actions: {
    async loadWorkOrders(_, prefix) {
      try {
        const response = await window.securedAxios.get(
          `/api/restricted/work-orders/?prefix=${prefix ?? ''}`,
        );
        return response.data;
      } catch (err) {
        return Promise.reject(err);
      }
    },
  },
};
