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
    async getWorkOrder(_, number) {
      try {
        const response = await window.securedAxios.get(`/api/restricted/work-orders/${number}`);
        return response.data;
      } catch (err) {
        return Promise.reject(err);
      }
    },
    async createWorkOrder(_, formData) {
      try {
        const response = await window.securedAxios.post('/api/restricted/work-orders', formData, {
          headers: { 'Content-Type': 'multipart/form-data' },
        });
        return response.data;
      } catch (err) {
        return Promise.reject(err);
      }
    },
  },
};
