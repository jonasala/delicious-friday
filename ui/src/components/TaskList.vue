<template>
  <div>
    <v-list :two-line="all == 0" :three-line="all == 1">
      <v-skeleton-loader :type="skel" v-if="loading" />
      <v-list-item v-if="taskLength === 0">
        <v-list-item-content>
          <v-list-item-title class="text-center">Nenhuma tarefa encontrada</v-list-item-title>
        </v-list-item-content>
      </v-list-item>
      <template v-for="(tasks, date) in tasksByDate">
        <v-subheader :key="`subheader=${date}`">{{ format(date) }}</v-subheader>
        <v-list-item
          v-for="task in tasks"
          :key="`task-${date}.${task.id}`"
          @click="activeTask = task"
        >
          <v-list-item-content>
            <v-list-item-subtitle v-if="all == 1">
              Criado por {{ task.created_by }}
            </v-list-item-subtitle>
            <v-list-item-title>
              {{ task.work_order ? `Ordem ${task.work_order.number}` : 'Sem ordem de serviço' }}
            </v-list-item-title>
            <v-list-item-subtitle
              v-html="task.description ? task.description : '<i>Sem descrição</i>'"
            ></v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
      </template>
    </v-list>
    <v-dialog v-model="dialogTask" width="500">
      <task-details v-if="dialogTask" :task="activeTask" @close="dialogTask = null" />
    </v-dialog>
  </div>
</template>

<script>
import { mapState } from 'vuex';
import { DateTime } from 'luxon';
import _ from 'lodash';
import TaskDetails from './TaskDetails.vue';

export default {
  components: { TaskDetails },
  props: ['all'],
  data() {
    return {
      loading: false,
      skel: 'list-item-three-line, list-item-three-line, list-item-three-line, list-item-three-line, list-item-three-line, list-item-three-line, list-item-three-line, list-item-three-line',
      tasksByDate: {},
      page: 1,
      activeTask: null,
    };
  },
  computed: {
    ...mapState({
      search: (state) => state.ui.toolbar.searchText,
    }),
    taskLength() {
      return Object.keys(this.tasksByDate).length;
    },
    dialogTask: {
      get() {
        return this.activeTask !== null;
      },
      set(val) {
        if (!val) {
          this.activeTask = null;
        }
      },
    },
  },
  mounted() {
    this.loadTasks();
  },
  methods: {
    async loadTasks() {
      try {
        this.loading = true;
        const tasks = await this.$store.dispatch('tasks/list', {
          search: this.search || '',
          page: this.page,
          itemsPerPage: 10,
          all: this.all,
        });
        this.tasksByDate = {};
        tasks.data.map((task) => {
          const d = DateTime.fromISO(task.created_at).toISODate();
          if (!this.tasksByDate[d]) {
            this.tasksByDate[d] = [];
          }
          this.tasksByDate[d].push(task);
          return task;
        });
      } catch (err) {
        console.log(err);
      } finally {
        this.loading = false;
      }
    },
    format(datetime) {
      return DateTime.fromISO(datetime).toLocaleString(DateTime.DATE_MED);
    },
    debouncedSearch: _.debounce((vm) => {
      vm.loadTasks();
    }, 500),
  },
  watch: {
    search() {
      this.debouncedSearch(this);
    },
  },
};
</script>
