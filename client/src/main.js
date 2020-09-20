import Vue from "vue";
import App from "./App.vue";
import router from "./router";

import * as firebase from "firebase/app";
import "firebase/auth";
import "firebase/firestore";

import { BootstrapVue, IconsPlugin } from "bootstrap-vue";
import "bootstrap/dist/css/bootstrap.css";
import "bootstrap-vue/dist/bootstrap-vue.css";

Vue.config.productionTip = false;
Vue.use(BootstrapVue);
Vue.use(IconsPlugin);

const firebaseConfig = {
  apiKey: "AIzaSyCEdlcKinO_em8f_ymWrE3_qAkaMLftNms",
  authDomain: "hoge-hoge-123456789.firebaseapp.com",
  databaseURL: "https://hoge-hoge-123456789.firebaseio.com",
  projectId: "hoge-hoge-123456789",
  storageBucket: "hoge-hoge-123456789.appspot.com",
  messagingSenderId: "725316375063",
  appId: "1:725316375063:web:dc4c17bddbf092a846c252"
};
firebase.initializeApp(firebaseConfig);

new Vue({
  router,
  render: h => h(App)
}).$mount("#app");
