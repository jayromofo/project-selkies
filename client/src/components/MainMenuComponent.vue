<script setup lang="ts">

import axios from 'axios'

const body = async function testdata(){
  try {
    const response = await axios('http://localhost:4444/sample');
    console.log(response);
  } catch (error){
    console.error(error)
  }

}

</script>

<template>
  <div>
    <h2>This is the Main Menu</h2>
    <p>
      <router-link to="/"></router-link>
      <router-link to="/about">About</router-link>
      <router-link to="/bio">Bio</router-link>
      <router-link to="/recipes">Recipes</router-link>
      <span>{{ body.name }}</span>
      
      
    </p>
  </div>
  
  <router-view></router-view>
  <button onclick="testdata()">Testing</button>
</template>

<style scoped></style>
