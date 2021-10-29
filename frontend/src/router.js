import Vue from "vue";
import Router from "vue-router";

Vue.use(Router);

export default new Router({
  mode: "history",
  routes: [
    {
      path: "/",
      alias: "/accounts",
      name: "accounts",
      component: () => import("./components/AccountCompany.vue")
    },
    {
      path: "/accounts/:id",
      name: "accounts-detail",
      component: () => import("./components/AccountCompany.vue")
    },
    {
      path: "/contact",
      name: "contact",
      component: () => import("./components/Contact.vue")
    }
  ]
});