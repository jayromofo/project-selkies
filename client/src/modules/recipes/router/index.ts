import type { RouteRecordRaw } from 'vue-router'
import CreateRecipe from '../views/CreateRecipe.vue';
const Module = () => import('../RecipeModule.vue');
const Home = () => import('../views/HomeView.vue');
const Recipe = () => import('../views/RecipesView.vue');

const recipeRoute: Array<RouteRecordRaw> = [
   {
      path: '/recipes',
      components: {
         default: Module,
         recipeContent: Module
      },
      children: [
         {
            path: '',
            name: 'home',
            components: {
               default: Home,
               recipeContent: Home
            }
         },
         {
            path: ':id',
            components: {
               default: Recipe,
               recipeContent: Recipe
            }
         },
         {
         path: '/create',
            name: 'create',
            components: {
               default: CreateRecipe,
               recipeContent: CreateRecipe
            }
         }

      ]
   },
   
   
];

export default recipeRoute

