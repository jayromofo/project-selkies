const Module = () => import('../BudgetModule.vue');
const Budget = () => import('../views/BudgetView.vue')
const Home = () => import('../views/HomeView.vue')
const Sidebar = () => import('../views/BudgetSidebar.vue')


const moduleRoute = {
   path: '/budget',
   components: {
      default: Module,
      budgetContent: Module,
      budgetSidebar: Sidebar,
   },
   children: [
      {
         path: '',
         components: {            
            budgetContent: Home,
            budgetSidebar: Sidebar,
            
         },
      },

      {
         path: ':id',
         components: {
            budgetContent: Budget,
            budgetSidebar: Sidebar,
         }
      }
   ]
}


export default moduleRoute
