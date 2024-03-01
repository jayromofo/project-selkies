import { createRouter, createWebHistory } from 'vue-router';
import type {RouteRecordRaw} from 'vue-router';

/* Import Routes */
import recipeRoute from '@/modules/recipes/router';
import budgetRoute from '@/modules/budget/router';

/* Import Pages */
import UHome from '@/pages/UHome.vue';


/* Create Routes*/
const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    redirect: "/home",
    children: [
      {
        path: '',
        component: UHome
      },
      {
        path: '/home',
        name: 'home',
        component: UHome
      },
    ],
  },
  {
    path: '/about',
    name: 'about',
    component: () => import('../views/AboutView.vue'),      
  },
  {
    path: '/contact-us',
    name: 'contactus',
    component: () => import('../views/AboutView.vue')
  },
  {
    path: '/register',
    name: 'register',
    component: () => import('../views/AboutView.vue')
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('../views/AboutView.vue')
  },
  {
    path: '/dashboard',
    name: 'dashboard',
    component: () => import('../views/AboutView.vue')
  },
  {
    path: '/:pathMatch(.*)',       
    component: () => import('../pages/404-not-found.vue')
  },
];


const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: routes.concat(recipeRoute, budgetRoute) /* Add each module routes to here */ 
});



export default router
