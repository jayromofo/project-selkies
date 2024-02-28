import type { RouteRecordRaw } from 'vue-router'
const Module = () => import('../RecipeModule.vue');
const Home = () => import('../views/Home.vue');
const Recipe = () => import('../views/Recipes.vue');

const recipeRoute: Array<RouteRecordRaw> = [
   {
      path: '/recipes',
      name: 'recipes',
      components: {
        default: Module,
        mainContent: Module, 
   },
   children: [
      {
         path: '/',
         component: Home
      },

      {
         path: ':id',
         component: Recipe
      }
   ]
   
}];

export default recipeRoute

