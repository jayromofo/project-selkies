import { createRouter, createWebHistory } from 'vue-router';
import type {RouteRecordRaw} from 'vue-router';

import RecipeModule from '@/modules/recipes/RecipeModule.vue';
import recipeRoute from '@/modules/recipes/router';
import MainLayout from '@/layouts/MainLayout.vue';
import budgetRoute from '@/modules/budget/router';
import UHome from '@/pages/UHome.vue';


const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'home',
    component: MainLayout,   
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (About.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
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
];


const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: routes.concat(recipeRoute, budgetRoute) /* Add each module routes to here */ 
});



export default router
