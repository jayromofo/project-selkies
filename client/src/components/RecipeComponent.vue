<script setup lang="ts">
import axios from 'axios'

const data = axios.get('http://localhost:4444/api/recipes/1');
console.log(data);

</script>



<template>
   <div>
      Recipe
      {{ data }}
   </div>
</template>

<style scoped>

</style>