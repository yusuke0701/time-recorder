import Vue from "vue";
import VueRouter from "vue-router";
import firebase from "firebase";

import Home from "../views/Home.vue";
import Calendar from "../views/Calendar";
import RecordList from "../views/RecordList";
import Signup from "../components/Signup";
import Signin from "../components/Signin";

Vue.use(VueRouter);

const routes = [
  {
    path: "*",
    redirect: "signin"
  },
  {
    path: "/",
    component: Home
  },
  {
    path: "/signup",
    component: Signup
  },
  {
    path: "/signin",
    component: Signin
  },
  {
    path: "/calendar",
    component: Calendar
  },
  {
    path: "/list/:selectedDate",
    component: RecordList
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

router.beforeEach((to, from, next) => {
  let requiresAuth = to.matched.some(record => record.meta.requiresAuth);
  let currentUser = firebase.auth().currentUser;
  if (requiresAuth) {
    // このルートはログインされているかどうか認証が必要です。
    // もしされていないならば、ログインページにリダイレクトします。
    if (!currentUser) {
      next({
        path: "/signin",
        query: { redirect: to.fullPath }
      });
    } else {
      next();
    }
  } else {
    next(); // next() を常に呼び出すようにしてください!
  }
});

export default router;
