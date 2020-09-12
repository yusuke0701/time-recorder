import Vue from 'vue';
import App from './App';

import { BootstrapVue, IconsPlugin } from 'bootstrap-vue';
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';

import router from './router';

global.browser = require('webextension-polyfill');
Vue.prototype.$browser = global.browser;

Vue.use(BootstrapVue);
Vue.use(IconsPlugin);

/* eslint-disable no-new */
new Vue({
  el: '#app',

  router,
  render: h => h(App),
});
