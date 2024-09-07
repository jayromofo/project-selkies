<template>
   <div>
      <h1 class="text-primary text-center">{{ recipe?.name }} </h1>
      <p>{{ recipe?.description }}</p>
      <div class="flex flex-col items-center">
         <span class="col-auto">
            <img src="/src/assets/imgs/Throne.jpg" :alt="recipe?.name">
         </span>
         <span class="col-auto">
            <img src="/src/assets/imgs/Throne.jpg" :alt="recipe?.name">

         </span>
      </div>
   
      <VaList>
         <VaListLable>Instructions</VaListLable>
         <VaListItem
            v-for="(instruction, index) in recipe?.Instructions"
            :key="index"
            class="list__item"
            >
            <VaListItemSection>
               <VaListItemLabel>
                  {{ instruction.line_num }} : {{ instruction.instruction }}
               </VaListItemLabel>
            </VaListItemSection>
         </VaListItem>
      </VaList>


   </div>

</template>

<script setup lang="ts">
import {ref, onMounted} from 'vue';
import {useRoute} from 'vue-router';
import axios from 'axios';
import type {Recipe, RecipeInstruction, MetaData} from "@/types/recipeTypes";
import { useRecipeStore } from '../stores/index';
import { VaListItem, VaListItemSection } from 'vuestic-ui';


var recJSON = {};

const route = useRoute();
const recipe = ref<Recipe | null>(null);
const metadata = ref<MetaData | null>(null);
const instructions = ref<RecipeInstruction[]>([]);

const store = useRecipeStore();

interface BackEndResponse {
   recipe: Recipe, 
   instruction: RecipeInstruction,
   metaData: MetaData
}

onMounted(async () => {
   const recipeId = route.params.id;

   console.log('id: '+ recipeId);
   try{
      const response = await axios.get<Recipe>(`http://localhost:4444/api/recipe/${recipeId}`); 
   
      recipe.value = response.data;    

      if (recipe.value)
         console.log(recipe.value);
      
   } catch (error) {
      console.error("Error fetching recipe data", error);
   }

   store.$reset()

   store.$patch((state) => {
      state.recipe = recipe.value
   })

});

</script>