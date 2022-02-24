<template>
  <v-card>
    <v-card-title>
      Detalhes da tarefa
      <v-spacer />
      <v-btn icon @click="$emit('close')">
        <v-icon>mdi-close</v-icon>
      </v-btn>
    </v-card-title>
    <v-card-text class="px-0">
      <v-list two-line>
        <v-list-item>
          <v-list-item-content>
            <v-list-item-subtitle>Ordem de serviço</v-list-item-subtitle>
            <v-list-item-title>
              {{ task.work_order ? `Ordem ${task.work_order.number}` : 'Sem ordem de serviço' }}
            </v-list-item-title>
          </v-list-item-content>
          <v-list-item-action v-if="task.work_order">
            <v-btn icon @click="dialogWODetails = true">
              <v-icon>mdi-file-eye-outline</v-icon>
            </v-btn>
          </v-list-item-action>
        </v-list-item>
        <v-list-item>
          <v-list-item-content>
            <v-list-item-subtitle>Criação</v-list-item-subtitle>
            <v-list-item-title>
              {{ task.created_by }} em {{ format(task.created_at) }}
            </v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-list-item three-line>
          <v-list-item-content>
            <v-list-item-subtitle>Descrição da tarefa</v-list-item-subtitle>
            <v-list-item-title class="description">{{
              task.description ? task.description : '-'
            }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-card-text>
    <v-card-actions>
      <v-spacer />
      <v-btn depressed color="error"> <v-icon left>mdi-delete</v-icon> Excluir </v-btn>
      <v-btn depressed color="primary"><v-icon left>mdi-pencil</v-icon> Editar</v-btn>
    </v-card-actions>
    <v-dialog
      v-if="task.work_order"
      fullscreen
      hide-overlay
      transition="dialog-bottom-transition"
      v-model="dialogWODetails"
    >
      <work-order-details @close="dialogWODetails = false" :wo="task.work_order" />
    </v-dialog>
  </v-card>
</template>

<script>
import { DateTime } from 'luxon';
import WorkOrderDetails from './WorkOrderDetails.vue';

export default {
  components: { WorkOrderDetails },
  props: ['task'],
  data() {
    return {
      dialogWODetails: false,
    };
  },
  methods: {
    format(datetime) {
      return DateTime.fromISO(datetime).toLocaleString(DateTime.DATETIME_MED);
    },
  },
};
</script>

<style scoped>
.description {
  white-space: normal;
}
</style>
