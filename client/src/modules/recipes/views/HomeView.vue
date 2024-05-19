<template>
   <div class="">

      <table class="border border-separate border-spacing-2 border-slate-500">
         <thead>
            <tr>
               <th>Recipe Name</th>
               <th>Description</th>
               <th>Category</th>
            </tr>
         </thead>
<!--  
         <tbody>
            <tr v-for="recipe in recipes" :key="recipe.id">
               <td>{{recipe.name}}</td>
               <td>{{recipe.description}}</td>
               <td>{{recipe.category}}</td>
            </tr>
         </tbody>
-->   
      </table>


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
// import  BaseInput  from "@/components/BaseInput.vue"
import { ref, onMounted } from 'vue';
import axios from "axios";

const recipes = ref<RecipeType[]>([]);

interface RecipeType {
   name: string;
   description: string;
   category: string;
}

const createRecipe = () => {
   console.log('create recipe');
}


const fetchRecipes = async () => {
   try {
      const response = await axios.get<RecipeType[]>("http://localhost:4444/api/recipe/");
      const data = response.data;

      data.forEach((recipe: RecipeType) => {
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
   background-color: brown;
}

.tb-5 {
   background-color: blue;
   color: orange;
   /* width:50%; */
}
</style>
