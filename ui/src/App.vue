<template>
  <v-app>
    <router-view name="drawer" />
    <router-view name="toolbar" />
    <v-main>
      <v-slide-x-transition mode="out-in">
        <router-view />
      </v-slide-x-transition>
    </v-main>
    <v-snackbar v-model="snackbarOpen" :color="snackbar.color" :timeout="4000" top right multi-line>
      <v-icon class="mr-2">{{ snackbarIcon() }}</v-icon>
      {{ snackbar.text }}
      <template v-slot:action="{ attrs }">
        <v-btn :color="snackbar.color" depressed v-bind="attrs" @click="snackbarOpen = false">
          OK
        </v-btn>
      </template>
    </v-snackbar>
  </v-app>
</template>

<script>
import { mapState } from 'vuex';

export default {
  computed: {
    ...mapState({ snackbar: (state) => state.ui.snackbar }),
    snackbarOpen: {
      get() {
        return this.$store.state.ui.snackbar.open;
      },
      set() {
        this.$store.commit('ui/closeSnackbar');
      },
    },
  },
  methods: {
    snackbarIcon() {
      switch (this.snackbar.color) {
        case 'error':
          return 'mdi-alert-circle-outline';
        case 'warning':
          return 'mdi-alert-outline';
        case 'info':
          return 'mdi-information-outline';
        case 'success':
          return 'mdi-check-circle-outline';
        default:
          return '';
      }
    },
  },
};
</script>
