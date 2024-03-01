import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import PrimeVue from 'primevue/config' 



/* Import Modules */
/* Not sure if this is actually needed
import recipeModule from "./modules/recipes/RecipeModule.vue"
import budgetModule from "./modules/budget/BudgetModule.vue"

import { registerModules } from "./register-modules"

registerModules({
   budget: budgetModule,
   recipes: recipeModule,
});
*/

const app = createApp(App)

app.use(PrimeVue, {
   unstyled: true,
})

app.use(router)

app.mount('#app')
