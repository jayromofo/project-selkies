<script setup lang="ts">
import axios from 'axios'
import { onMounted, ref } from 'vue';

interface Recipe {
   id: number,
   name: string,
   description: string,
   type: string,
   category: string
}

const loading = ref(false);
const data = ref<Recipe[] | null >([]);
const error = ref< String | null>(null);
const state = ref <State | null>(null);
const newRecipe = ref<Recipe>({
   id: 0,
   name: '',
   description: '',
   type: '',
   category: '',
});

defineProps({
   modelValue: String,
})

const emit = defineEmits<{
   (event: 'update:modelValue', payload: String): void;
}>();

enum State {
   Create = 1,
   ViewAll,
   View,
   Edit,
}

const fetchData = async () => {
   loading.value = true;
   try {
      const response = await axios.get('http://localhost:4444/api/recipe/');
      data.value = response.data;

   } catch (err: any) {
      error.value = err.message || 'Error fetching data';
   } finally {
      loading.value = false;
      console.log(data);
   }
}

const addRecipe = () => {
   console.log('Form Submitted with data: ', newRecipe.value)
   if (newRecipe.value?.name.trim() != undefined && data.value != undefined) {
      const newId  = Math.max(...data.value.map((item) => item.id), 0) + 1;
   }
   axios.post('http://localhost:4444/api/recipe/', newRecipe.value);
}


const deleteRecipe = (itemId: number) => {
   const index = data.value?.findIndex((item) => item.id === itemId)
   if (index !== -1 && index !== undefined) {
      data.value?.splice(index, 1);
   }
};




onMounted(fetchData);
state.value = State.ViewAll;
console.log(state.value);



console.log(data);

</script>


<template>
   <div class="recipe">
      <div v-if="state == State.Create">
         <span>This is State Create</span>
         <form @submit.prevent="addRecipe">
            <label for="recipeName">Recipe Name: </label>
            <input type="text" id="recipeName" v-model="newRecipe.name" required /><br />
            <label for="recipeDescription">Description: </label>
            <input type="text" id="recipeDescription" v-model="newRecipe.description" required /><br />
            <label for="recipeType">Recipe Type: </label>
            <input type="text" id="recipeName" v-model="newRecipe.type" required /><br />
            <label for="recipeCategory">Recipe Category: </label>
            <input type="text" id="recipeCategory" v-model="newRecipe.category" required /><br />

            <button type="submit">Add Recipe</button>          
         </form>
      </div>
      <div v-if="state == State.ViewAll">
         <p v-if="loading">Loading...</p>
         <ul v-if="data">
            <li class="recipe-link" v-for="recipe in data" :key="recipe.id">{{ recipe }}</li>
         </ul>
      </div>
   </div>
</template>
 
<style>
@media (min-width: 1024px) {
   .recipe {
      min-height: 100vh;
      display: flex;
      align-items: center;      
   }
   .recipe-link {
      color:white;
   }
}
</style>
