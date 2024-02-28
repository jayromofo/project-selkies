<template>
<div class="input-wrap">
<label>{{ label }}</label>
<input 
   :class="classValue"
   :placeholder="label"
   :value="modelValue"
   
   @input="$emit('update:modelValue', $event.target)"
/>
</div>

</template>

<script setup lang="ts">
import { ref } from 'vue';

// @input="handleInput"

defineProps({
   label: {
      type: String,
      default: '',
   },
   modelValue: {
      type: [ String, Number ],
      default: ''
   },
   classValue: {
      type: String,
      default: ''
   },
   
})

const textInput = ref("");


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