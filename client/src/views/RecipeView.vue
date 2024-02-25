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
const data = ref<Recipe[] | null>(null);
const error = ref< string | null>(null);



const fetchData = async () => {
   loading.value = true;
   try {
      const response = await axios.get('http://localhost:4444/api/recipe/1');
      data.value = response.data;
   } catch (err: any) {
      error.value = err.message || 'Error fetching data';
   } finally {
      loading.value = false;
      console.log(data)
   }

}

onMounted(fetchData);



console.log(data);

</script>


<template>
   <div class="recipe">
      <p v-if="loading">Loading...</p>
      <ul v-if="data">
         <li class="recipe-link" v-for="recipe in data" :key="recipe.id">{{ recipe }}</li>
      </ul>
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
      font:white;
   }
}
</style>
