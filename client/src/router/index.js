import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";
import Calendar from "../views/Calendar";
import RecordList from "../views/RecordList";
import Signup from "../components/Signup";
import Signin from "../components/Signin";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    component: Home
  },
  {
    path: '/signup',
    component: Signup
  },
  {
    path: '/signin',
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

export default router;
