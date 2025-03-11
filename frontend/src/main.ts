import { createApp } from "vue";
import App from "./App.vue";
import { createPinia } from "pinia";

import { Quasar } from "quasar";
import quasarIconSet from "quasar/icon-set/svg-mdi-v7";

import "@quasar/extras/roboto-font/roboto-font.css";
import "@quasar/extras/mdi-v7/mdi-v7.css";
import "quasar/dist/quasar.css";

const app = createApp(App);

app.use(createPinia());
app.use(Quasar, {
  plugins: {},
  iconSet: quasarIconSet,
});

app.mount("#app");
