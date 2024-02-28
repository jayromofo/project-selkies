const Module = () => import('../Module.vue');
const Budget = () => import('../views/Budget.vue')
const Home = () => import('../views/Home.vue')


const moduleRoute = {
   path: '/budget',
   components: {
      default: Module,
      mainContent: Module,
   },
   children: [
      {
         path: '/',
         components: {
            default: Home,
            mainContent: Home,
         },
      },

      {
         path: ':id',
         component: Budget
      }
   ]
}


export default moduleRoute
