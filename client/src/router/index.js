import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";
import Calendar from "../views/Calendar";
import RecordList from "../views/RecordList";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    component: Home
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
