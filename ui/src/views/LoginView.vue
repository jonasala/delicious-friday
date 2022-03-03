<template>
  <v-container fill-height align-center justify-center>
    <v-card width="350" elevation="10">
      <v-card-text class="text-center">
        <img src="../assets/logo.png" width="120" alt="d3v-logo" />
        <h2 class="font-weight-light my-5">Controle de Tarefas</h2>
        <v-form @submit.prevent="login()">
          <v-text-field
            v-model="user.username"
            type="text"
            label="Usuário"
            :error-messages="usernameErrors"
            prepend-icon="mdi-account"
            @blur="$v.user.username.$touch()"
            @input="$v.user.username.$touch()"
            :disabled="loading"
          />
          <v-text-field
            v-model="user.password"
            type="password"
            label="Senha"
            :error-messages="passwordErrors"
            prepend-icon="mdi-lock"
            @blur="$v.user.password.$touch()"
            @input="$v.user.password.$touch()"
            :disabled="loading"
          />
          <v-btn type="submit" color="accent" class="mt-2" block depressed :loading="loading">
            <v-icon class="mr-2">mdi-login</v-icon> Entrar
          </v-btn>
        </v-form>
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script>
import { required } from 'vuelidate/lib/validators';

export default {
  data() {
    return {
      loading: false,
      user: {
        username: '',
        password: '',
      },
    };
  },
  validations: {
    user: {
      username: { required },
      password: { required },
    },
  },
  computed: {
    usernameErrors() {
      const errors = [];
      if (!this.$v.user.username.$dirty) {
        return errors;
      }
      if (!this.$v.user.username.required) {
        errors.push('O nome de usuário é obrigatório');
      }
      return errors;
    },
    passwordErrors() {
      const errors = [];
      if (!this.$v.user.password.$dirty) {
        return errors;
      }
      if (!this.$v.user.password.required) {
        errors.push('A senha é obrigatória');
      }
      return errors;
    },
  },
  mounted() {
    document.title = 'Login - D3V Controle de Tarefas';
  },
  methods: {
    async login() {
      try {
        this.$v.$touch();
        if (!this.$v.$invalid) {
          this.loading = true;
          await this.$store.dispatch('users/login', this.user);
          this.$router.replace({ name: 'my-tasks' });
        }
      } catch (err) {
        if (err.response && err.response.status === 401) {
          this.$store.commit('ui/setSnackbar', {
            text: 'Usuário e senha não conferem',
            color: 'warning',
          });
        }
      } finally {
        this.loading = false;
      }
    },
  },
};
</script>

<style scoped>
.container {
  background-color: #435967;
  width: 100% !important;
  max-width: 100% !important;
}
</style>
