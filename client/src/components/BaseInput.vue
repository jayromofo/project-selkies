<template>
<div class="input-wrap">
<label>{{ label }}</label>
<input 
type="text" 
:value="modelValue"
:class="classValue"
@input="handleInput"
/>
</div>

</template>

<script setup lang="ts">
defineProps({
   label: {
      type: String,
      default: '',
   },
   modelValue: {
      type: String,
      default: ''
   },
   classValue: {
      type: String,
      default: ''
   },
   
})

const handleInput = (event: Event) => {
   const target = event.target as HTMLInputElement;
   const inputValue: string = target.value;
   console.log(inputValue);
}

</script>


<style scoped>
.input-wrap {
   display: flex;
   flex-direction: column;


   input {
      padding: 8px 12px;
      font-size: 16px;
   }
}

</style>