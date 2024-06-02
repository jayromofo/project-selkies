<template>
   <div>
      <h1>{{ recipe?.name }} </h1>
      <p>{{ recipe?.description }}</p>
      <h2>Instructions</h2>      
      <ul>
         <li v-for="instruction in recipe?.Instructions" :key="instruction.line_num">{{ instruction.line_num }} : {{ instruction.instruction }}</li>
      </ul>


   </div>

</template>

<script setup lang="ts">
import {ref, onMounted} from 'vue';
import {useRoute} from 'vue-router';
import axios from 'axios';
import type {Recipe, RecipeInstruction, MetaData} from "../../../types/recipeTypes";

declare global {
   interface Window {
      recipe: Recipe | undefined;
   }
}

var recJSON = {};

// interface Recipe {
//    ID: number;
//    name: string;
//    description: string;
//    category: string;   
// }

// interface MetaData {
//    id: number;
//    recipeId: number;
//    servings: number;
//    cooktime: number;
//    isKeto: boolean;
//    isVegetarian: boolean;
//    tags: string[];
//    imagePath: string;
// }

// interface RecipeInstruction {
//    id: number;
//    recipeId: number;
//    lineNum: number;
//    instruction: string;
// }

const route = useRoute();
const recipe = ref<Recipe | null>(null);
const metadata = ref<MetaData | null>(null);
const instructions = ref<RecipeInstruction[]>([]);

/*
      const [recipeResponse, instructionsResponse, metadataResponse] = await Promise.all([
         axios.get(`http://localhost:4444/api/recipe/${recipeId}`),
         axios.get(`http://localhost:4444/api/recipe/${recipeId}/instructions`),
         axios.get(`http://localhost:4444/api/recipe/${recipeId}/metadata`),
      ]);
*/

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
         // instructions.value = response.data.instructions;
         // metadata.value = response.data.Metadata;      
      window.recipe = recipe.value;  

      if (recipe.value)
         console.log(recipe.value);
      
   } catch (error) {
      console.error("Error fetching recipe data", error);
   }

});

/*
onMounted(async () => {
   const recipeId = route.params.id;
   console.log('id: '+ recipeId);
   try{
      const recipeResponse = await axios.get(`http://localhost:4444/api/recipe/${recipeId}`);
      
      recipe.value = recipeResponse.data.recipe;
      metadata.value = recipeResponse.data.metadata;
      instructions.value = recipeResponse.data.instructions;

   } catch (error) {
      console.error("Error fetching recipe data", error);
   }

   console.log(recipe);
});
*/


</script>