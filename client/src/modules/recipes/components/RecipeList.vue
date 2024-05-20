<template>
<div style="background-color: white;">
   <ul>
      <li v-for="(recipe, index) in recipes" :key="recipe.id">
      <h2> {{ recipe.name }}</h2>
      <p>{{ recipe.description }}</p>
      <p><strong>Category: </strong> {{ recipe.category }}</p>
      </li>
   </ul>
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

</style>