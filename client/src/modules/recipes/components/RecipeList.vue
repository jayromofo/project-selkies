<template>
<div style="background-color: white;">
   <div class="recipe-container">
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
</template>

<script setup lang="ts">

interface Recipe {
   id: number;
   name: string;
   description: string;
   category: string;
}

const props = defineProps<{
   recipes: Recipe;
}>();

console.log(props);




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