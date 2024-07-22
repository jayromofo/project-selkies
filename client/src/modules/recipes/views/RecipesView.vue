<template>
   <div>
      This is a recipe
      {{ recipe?.name }}

   </div>

</template>

<script setup lang="ts">
import {ref, onMounted} from 'vue';
import {useRoute} from 'vue-router';
import axios from 'axios';


interface Recipe {
   id: number;
   name: string;
   description: string;
   category: string;   
}

interface MetaData {
   id: number;
   recipeId: number;
   servings: number;
   cooktime: number;
   isKeto: boolean;
   isVegetarian: boolean;
   tags: string[];
   imagePath: string;
}

interface RecipeInstruction {
   id: number;
   recipeId: number;
   lineNum: number;
   instruction: string;
}

const route = useRoute();
const recipe = ref<Recipe | null>(null);
const recipeMetadata = ref<MetaData | null>(null);
const recipeInstructions = ref<RecipeInstruction | null>(null);


onMounted(async () => {
   const recipeId = route.params.id;
   console.log('id: '+ recipeId);
   try{
      const [recipeResponse, instructionsResponse, metadataResponse] = await Promise.all([
         axios.get(`http://localhost:4444/api/recipe/${recipeId}`),
         axios.get(`http://localhost:4444/api/recipe/${recipeId}/instructions`),
         axios.get(`http://localhost:4444/api/recipe/${recipeId}/metadata`),
      ]);
      recipe.value = recipeResponse.data;
      recipeInstructions.value = instructionsResponse.data;
      recipeMetadata.value = metadataResponse.data;

   } catch (error) {
      console.error("Error fetching recipe data", error);
   }
});


</script>