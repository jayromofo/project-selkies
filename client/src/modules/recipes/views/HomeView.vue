<template>
   <div class="container">
      <h1>Recipe List</h1>
      <NewRecipe @addRecipe="addRecipe"/>
      <RecipeList :recipes="recipes" />

      <!-- 
         <table class="border border-separate border-spacing-2 border-slate-500">
            <thead>
               <tr>
                  <th>Recipe Name</th>
                  <th>Description</th>
                  <th>Category</th>
               </tr>
            </thead>

       -->
<!--  
         <tbody>
            <tr v-for="recipe in recipes" :key="recipe.id">
               <td>{{recipe.name}}</td>
               <td>{{recipe.description}}</td>
               <td>{{recipe.category}}</td>
            </tr>
         </tbody>
-->   


      <!--
      <form>
         <BaseInput 
            v-model="recipe.title"
            label="Recipe Name"
            type="text"
            class="tb-5"
            width="10"
         />
         <button>Button</button>

      </form>    
   -->
      <button id="btnCreate" v-on:click="createRecipe" class="btn border-t-neutral-400"><a href="/recipes/create">Create a recipe</a></button>
   </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import axios from "axios";

import { NewRecipe, RecipeList } from '../components';

const recipes = ref<Recipe[]>([]);

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


const fetchRecipes = async () => {
   try {
      const response = await axios.get<Recipe[]>("http://localhost:4444/api/recipe/");
      const data = response.data;

      data.forEach((recipe: Recipe) => {
         console.log(recipe.name);
         
      });

      recipes.value = data;

   } catch (error) {
      console.error("Error fetching data: ", error);
   }
} /* fetchRecipes */

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
