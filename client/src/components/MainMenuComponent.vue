<script setup lang="ts">
</script>

<template>
  <nav id="header-links"> 
      <router-link class="header-link" to="/">Home</router-link>
      <router-link class="header-link" to="about">About</router-link>
      <router-link class="header-link" to="budget">Budget</router-link>
      <router-link class="header-link" to="recipes">Recipes</router-link>         
  </nav>
</template>

<style scoped>
  .header-link {
    padding: 10px;

  }

  .router-link-exact-active {
    color: red;
  }
</style>
