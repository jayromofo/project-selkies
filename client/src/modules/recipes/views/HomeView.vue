<template>
   <div class="container-create">
      <button class="white bg-slate-400" @click="showCreate">Create a recipe</button>
      <NewRecipe v-if="isCreate" @addRecipe="addRecipe"/>
   </div>




   <div v-if="!isCreate" class="container mx-auto px-4 py-8">
      <h1>Recipe List</h1>
      <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
         <RecipeList v-for="recipe in recipes" :key="recipe.id" :recipe="recipe" />
      </div>
   </div>
   <div>
      <button id="sample form" class='btn' :onClick="showSample">Show Sample </button>

   </div> 
     
      <!-- <button id="btnCreate" v-on:click="createRecipe" class="btn border-t-neutral-400"><a href="/recipes/create">Create a recipe</a></button> -->

   <div v-if="isVisible" class="sampleContainer" >     
      <div class="max-w-lg mx-auto p-8 bg-white rounded-lg shadow-md">
         <h2 class="text-2xl font-semibold mb-6">Contact Us</h2>
         <form form @submit.prevent="handleSomething">
            <div class="mb-4">
               <label for="name" class="block text-sm font-medium text-gray-700">Name</label>
               <input type="text" id="name"  class="mt-1 block w-full p-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm" />
            </div>
            <div class="mb-4">
               <label for="email" class="block text-sm font-medium text-gray-700">Email</label>
               <input type="email" id="email"  class="mt-1 block w-full p-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm" />
            </div>
            <div class="mb-4">
               <label for="message" class="block text-sm font-medium text-gray-700">Message</label>
               <textarea id="message" rows="4" class="mt-1 block w-full p-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"></textarea>
            </div>      
         </form>
      </div>
   </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import axios from "axios";

import { NewRecipe, RecipeList } from '../components';

const recipes = ref<Recipe[]>([]);
const isVisible = ref(false);
const isCreate = ref(false);

interface Recipe {
   id: number;
   name: string;
   description: string;
   category: string;
}

const addRecipe = (newRecipe: Recipe) => {
   recipes.value.push(newRecipe);
}

const createRecipe = () => {
   console.log('create recipe');
}

const showSample = () => {
   console.log('show sample');
   isVisible.value = !isVisible.value;
}

const handleSomething = () => {

}

const showCreate = () => {
   isCreate.value = !isCreate.value;
}

/* http://localhost:4444/api/recipe/*/
const fetchRecipes = async () => {
   try {
      const response = await axios.get<Recipe[]>("http://localhost:4444/api/recipe/");
      const data = response.data;

      console.log(response.data);
      
      data.forEach((recipe: Recipe) => {
         console.log(recipe.name);
         
      });
      
      recipes.value = data;

   } catch (error) {
      console.error("Error fetching data: ", error);
   }
}  /* fetchRecipes */ 

onMounted(() =>  {
   fetchRecipes();
});

</script>

<style scoped>
.container {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  text-align: center;
  color: blue;
  margin-top: 60px;
}

form {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
}

div {
  margin-bottom: 10px;
}

label {
  margin-right: 10px;
}

input, textarea {
  padding: 8px;
  font-size: 16px;
  width: 300px;
}

button {
  padding: 8px 16px;
  font-size: 16px;
  cursor: pointer;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  padding: 16px;
  border-bottom: 1px solid #ccc;
}

h2 {
  margin: 0;
}

p {
  margin: 5px 0;
}
</style>
