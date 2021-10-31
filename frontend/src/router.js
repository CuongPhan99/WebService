import Vue from "vue";
import Router from "vue-router";

Vue.use(Router);

export default new Router({
  mode: "history",
  routes: [
    {
      path: "/:name",
      name: "home",
      component: () => import("./App.vue")
    },
    {
      path: "/:name/:id",
      name: "accountDetail",
      component: () => import("./components/AccountCompany.vue")
    },
    {
      path: "/contact",
      name: "contact",
      component: () => import("./components/Contact.vue")
    }
  ]
});