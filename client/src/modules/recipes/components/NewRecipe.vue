<template>


<!-- Original Create -->
<div :v-if="isOriginal" style="background-color: gray;">
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
// import { FormKit } from '@formkit/vue';
import type { Recipe } from '@/types/recipeTypes';

const name = ref('');
const description = ref('');
const category = ref('');
const isSubmitting = ref<boolean>(false);
const errorMessage = ref<string | null>();
const isOriginal = ref<boolean>(true)

const submitted = ref(false)
const submitHandler = async () => {
      await new Promise((r) => setTimeout(r, 2000))
      submitted.value = true
}

const newRecipe = ref<Recipe>();

/*** Error Handling ***/

interface CustomError {
   name: string;
   message: string;
   status: string;
}


function isCustomError(error: unknown): error is CustomError {
   return (
      typeof error === 'object' &&
      error !== null &&
      'name' in error &&
      'message' in error &&
      'status' in error
   );
}

function createCustomError(message: string, status: string,): CustomError {
   return {
      name: "CustomError",
      message,
      status
   };
}

function someFunction() {
   throw createCustomError('Something went wrong', '400');
}


const customError = ref<CustomError>()
/*** Error Handling ***/



const emit = defineEmits<{
   (e: 'addRecipe', newRecipe: Recipe): void;
}>();

const submitRecipe = async () => {
   isSubmitting.value = true;
   errorMessage.value = null;

   const newRecipe: Recipe = {
      ID: Date.now(),
      name: name.value,
      description: description.value,
      category: category.value,
      type: "" ,
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

   } catch (error: unknown ) {
      if (isCustomError(error)) {
         errorMessage.value = error.message || 'An error has occured';
      }
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