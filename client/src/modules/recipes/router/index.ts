import type { RouteRecordRaw } from 'vue-router'
// const CreateRecipe = () => import('../views/CreateRecipe.vue');
const Module = () => import('../RecipeModule.vue');
const Home = () => import('../views/HomeView.vue');
const Recipe = () => import('../views/RecipesView.vue');

const recipeRoute: Array<RouteRecordRaw> = [
   {
      // Base path
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
         path: '/create',
            name: 'create',
            component: () => import('../views/CreateRecipe.vue')
            // components: {
            //    default: CreateRecipe,
            //    recipeContent: CreateRecipe
            // }
         },
         {
            path: ':id',
            name: 'RecipeDetail',
            components: {
               default: Recipe,
               recipeContent: Recipe
            },
            props: true
         },

      ]
   },
   
   
];

export default recipeRoute

