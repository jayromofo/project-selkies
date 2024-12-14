<!-- https://formkit.com/guides/build-a-multi-step-form -->

<template>
   <div>

   </div>
   <!--
   <h1 class="text-2xl font-bold mb-2">Create a recipe!</h1>
   <p class="text-sm mb-4">
      Add recipe to storage so you can use the same recipe later
   </p>
   
   <FormKit type="form" default="{ value }" :plugins="[stepPlugin]">
      <ul class="steps">
         <li 
            v-for:="(step, stepName) in stepNames"
            class="step"
            @click="step.value = stepName"
            :data-step-active="stepName === stepName"
            >


         </li>
      </ul>
      
   

   <div class="form-body bg-gray-200 flex justify-center h-screen">

      <section v-show="step.value === 'recipeInfo'">
         <FormKit
            type="text"
            name="recipe_name"
            label="Recipe Name"
            placeholder="Pizza"
            help="What is your recipe called?"
            validation="required"
            class="rounded border border-solid border-slade-100 shadow-lg block"
         />
         <FormKit
            type="text"
            name="description"
            label="Recipe Description"
            placeholder="How does the recipe taste?"
            help="Describe the recipe"   
         />
      </section>
      
      <section v-show="step.value === 'recipeMetadata'">
            <FormKit type="group" v-model="metadata">
            <FormKit
            type="text"
            name="cook_time"
            label="Cook Time"
            placeholder="15m"
            help="How Long"   
            />
            </FormKit>      
         <pre wrap>{{ metadata }}</pre>
      </section>
   
      <section v-show="step.value === 'recipeInstructions'">
         <FormKit type="group" name:="recipeInstruction">
            <FormKit type="text" label="Instruction" validation="required|text" />
   
            
         </FormKit>
      </section>
      </div>
      <FormKit type="submit" label="Register" />
      <pre wrap>{{ value }}</pre>
      </FormKit>
      <div v-if="submitted">
         <h2 class="text-xl text-green-500">Submission successful!</h2>
      </div>
      
      
      
   -->
</template>

<script setup lang="ts">
import { ref, toRef } from 'vue'
import type { MetaData, Recipe, RecipeInstruction, RecipeCategory } from '@/types/recipeTypes';
import { FormKitMessages } from '@formkit/vue';
// import useSteps from '@/helpers/'

// const { steps, activeStep, stepPlugin } = useSteps()

const newRecipe = ref<Recipe>();
const submitted = ref(false);
const metadata = ref<MetaData>();
const instructions = ref<RecipeInstruction[]>([])
// const step = ref('recipeInfo');
// const steps = ref()

const stepNames = ['recipeInfo', 'recipeMetadata', 'recipeInstruction']

const submitHandler = () => {

}



const stepPlugin = (node: any) => {
   if (node.props.type == 'group') {
      // steps[node.name] = steps[node.name] || {}

      node.on('created', () => {
         // steps[node.name].valid = toRef(node.context.state, 'valid');
      })
      return false
   }
}
   


</script>


<style scoped>


</style>

