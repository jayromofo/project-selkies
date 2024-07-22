<template>
<div style="background-color: gray;">
   <form @submit.prevent="submitRecipe">
      <div>
         <label for="name">Name:</label>
         <input id='name' v-model="name" required />
      </div>
      <div>
         <label for="description">Description:</label>
         <textarea id='description' v-model="description" required></textarea>
      </div>
      <div>
         <label for="category">Category:</label>
         <input id='category' v-model="category" required />
      </div>      
      <button type="submit" :disabled="isSubmitting" @click="submitRecipe">Add Recipe</button>
   </form>
   <p v-if="errorMessage" class="error"> {{ errorMessage }}</p>
</div>
</template>

<script setup lang="ts">

import { ref } from 'vue';
import axios from 'axios';

interface Recipe {
   id: number;
   name: string;
   description: string;
   category: string;
}

const name = ref('');
const description = ref('');
const category = ref('');
const isSubmitting = ref<boolean>(false);
const errorMessage = ref<string | null>();

// const newRecipe = ref<Recipe>();

const emit = defineEmits<{
   (e: 'addRecipe', newRecipe: Recipe): void;
}>();

const submitRecipe = async () => {
   isSubmitting.value = true;
   errorMessage.value = null;

   const newRecipe: Recipe = {
      id: Date.now(),
      name: name.value,
      description: description.value,
      category: category.value
   };

   try {
      const response = await axios.post('http://localhost:4444/api/recipe/create/', newRecipe)
      console.log(response);
      if (response.status == 201) {
         emit('addRecipe', response.data);

         clearValues();
      } else {
         throw new Error('Failed to add recipe');
      }
   } catch (error) {
      errorMessage.value = error.message || 'An error has occured';
   }finally {
      isSubmitting.value = false;
   }
};

const clearValues = () => {
   name.value = '';
   description.value = '';
   category.value = '';
}


</script>


<style scoped>
.error {
   color: red;
   margin-top: 10px;
}

</style>