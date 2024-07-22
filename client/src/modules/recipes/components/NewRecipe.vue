<template>
 <FormKit
    type="form"
    id="registration-example"
    :form-class="submitted ? 'hide' : 'show'"
    submit-label="Register"
    @submit="submitHandler"
    :actions="false"
    #default="{ value }"
  >
    <h1 class="text-2xl font-bold mb-2">Register!</h1>
    <p class="text-sm mb-4">
      You can put any type of element inside a form, not just FormKit inputs
      (although only FormKit inputs are included with the submission).
    </p>
    <FormKit
      type="text"
      name="name"
      label="Your name"
      placeholder="Jane Doe"
      help="What do people call you?"
      validation="required"
    />
    <FormKit
      type="text"
      name="email"
      label="Your email"
      placeholder="jane@example.com"
      help="What email should we use?"
      validation="required|email"
    />
    <div class="double">
      <FormKit
        type="password"
        name="password"
        label="Password"
        validation="required|length:6|matches:/[^a-zA-Z]/"
        :validation-messages="{
          matches: 'Please include at least one symbol',
        }"
        placeholder="Your password"
        help="Choose a password"
      />
      <FormKit
        type="password"
        name="password_confirm"
        label="Confirm password"
        placeholder="Confirm password"
        validation="required|confirm"
        help="Confirm your password"
      />
    </div>

    <FormKit type="submit" label="Register" />
    <pre wrap>{{ value }}</pre>
  </FormKit>
  <div v-if="submitted">
    <h2 class="text-xl text-green-500">Submission successful!</h2>
  </div>






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
import { FormKit } from '@formkit/vue';
import { Recipe } from '@/types/recipeTypes';

const name = ref('');
const description = ref('');
const category = ref('');
const isSubmitting = ref<boolean>(false);
const errorMessage = ref<string | null>();
const isOriginal = ref<boolean>(false)

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