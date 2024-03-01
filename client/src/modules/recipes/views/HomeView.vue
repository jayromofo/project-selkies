<template>
   <div class="container col-md-6">
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
   </div>
</template>

<script setup lang="ts">
import  BaseInput  from "@/components/BaseInput.vue"
import {ref} from 'vue';

const recipe = ref({title: '', description: '', cook_time: 0})




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
