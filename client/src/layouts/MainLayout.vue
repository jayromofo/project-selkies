<template>
   <div id="main-layout">
      <HeaderComponent />
      <main>
         <router-view name="mainContent"></router-view>
      </main>
      <FooterComponent />
   </div>
</template>

<script setup lang="ts">
/* Import Components */
import HeaderComponent from '@/components/HeaderComponent.vue';
import FooterComponent from '@/components/FooterComponent.vue';

</script>

<style scoped> 
#main-layout {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  margin: auto;
}

main {
  flex: 3;
}


</style>