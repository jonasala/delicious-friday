import Vue from 'vue';
import Vuetify from 'vuetify/lib/framework';
import pt from 'vuetify/lib/locale/pt';
import colors from 'vuetify/es5/util/colors';

Vue.use(Vuetify);

export default new Vuetify({
  theme: {
    themes: {
      light: {
        primary: '#435967',
        secondary: colors.blue.darken2,
        accent: colors.pink,
        error: colors.red.darken2,
        info: colors.lightBlue,
        success: colors.green.darken1,
        warning: colors.yellow.darken4,
      },
    },
  },
  lang: {
    locales: { pt },
    current: 'pt',
  },
});
