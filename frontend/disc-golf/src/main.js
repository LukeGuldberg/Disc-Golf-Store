import { createApp } from 'vue';
import App from './App.vue';
import '@mdi/font/css/materialdesignicons.css';
import 'vuetify/styles';
import { createVuetify } from 'vuetify';
import * as components from 'vuetify/components';
import * as directives from 'vuetify/directives';
import { mdi } from 'vuetify/iconsets/mdi';
import { createRouter, createWebHistory } from 'vue-router';

// Importing your components
import Home from './Home.vue';

// Define your routes here
const routes = [
  { path: '/', component: Home }
];

// Create the router instance
const router = createRouter({
  history: createWebHistory(),
  routes,
});

// Create the Vuetify instance
const vuetify = createVuetify({
  components,
  directives,
  icons: {
    defaultSet: 'mdi',
    sets: {
      mdi,
    },
  },
});

// Create and mount the app
createApp(App)
  .use(router)  // Make sure to use the router
  .use(vuetify)
  .mount('#app');
