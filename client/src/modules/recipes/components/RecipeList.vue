<template>
<div class="max-w-sm rounded overflow-hidden shadow-lg">
    <img class="w-full" src="/src/assets/imgs/Throne.jpg" :alt="recipe.name">
    <div class="px-6 py-4">
      <div class="font-bold text-xl mb-2">
         <button @click="navigateToDetail" > {{ recipe.name }}</button>
      </div>
      <p class="text-gray-700 text-base">{{ recipe.description }}</p>
    </div>
    <div class="px-6 pt-4 pb-2">
      <span class="inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2 mb-2">{{ recipe.category }}</span>
    </div>
</div> 
   <!--
   <div v-if="!test" class="recipe-container">
      <ul>
         <li v-for="(recipe, index) in recipes" :key="recipe.id">
         <h2> {{ recipe.name }}</h2>
         <p>{{ recipe.description }}</p>
         <p><strong>Category: </strong> {{ recipe.category }}</p>
         <div class="row">
            <button class='btn' @click="editRecipe">Edit Recipe</button>
            <button class='btn' @click="deleteRecipe">Delete Recipe</button>
         </div>
         </li>
      </ul>
   </div>
</div> 
-->
<!--^^  default ^^-->
</template>

<script setup lang="ts">
import axios from 'axios';
import {ref, onMounted } from "vue";
import { useRoute, useRouter } from 'vue-router';

// const test = true;

interface Recipe {
   id: number;
   name: string;
   description: string;
   category: string;
}

const props = defineProps<{
   recipe: Recipe;
}>();

const route = useRoute();
const router = useRouter();
const recipeId = ref();

recipeId.value = route.params.ID;
const navigateToDetail = () => {
   router.push({name: 'RecipeDetail', params: {id: recipeId.value}})
}

</script>


<style scoped>
/* .recipe-container {
   color: orange;
}

.btn {
   box-sizing: border-box;
   border: none;
   font-family: inherit;
   padding: 0;
   cursor: pointer;


@media screen and (-ms-high-contrast: active) {
   border: 2px solid currentcolor;
   }  
}

a.btn, btn {
   display: inline-flex;
   align-items: center;
   justify-content: center;

   
} */

</style>