<script setup lang="ts">



</script>

<template>
  <div> 
    <p>
      <router-link to="/"></router-link>
      <router-link to="/about">About</router-link>
      <router-link to="/bio">Bio</router-link>
      <router-link to="/recipes">Recipes</router-link>         
    </p>
  </div>
  
  <router-view></router-view>
</template>

<style scoped></style>
