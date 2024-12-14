<script setup lang="ts">
import { RouterView } from 'vue-router'
import HeaderComponent from './components/HeaderComponent.vue'
import FooterComponent from './components/FooterComponent.vue'
</script>

<template>
  
  <!-- <div class="wrapper">
    <MainLayout />
  </div> -->

  <HeaderComponent />
  <div id="main-layout">

      <router-view></router-view>

  </div>
  <FooterComponent />
      
</template>

<style scoped>

</style>
