import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { Recipe } from '@/types/recipeTypes'



export const useRecipeStore = defineStore('recipe', () => {
   const recipe = ref<Recipe | null>()
   const recipeList = ref<Recipe[]>()   

   function $reset() {
      recipe.value = null
   }

   return {
      recipe,
      recipeList,    
      $reset
   }
})