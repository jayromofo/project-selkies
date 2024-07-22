<template>
   <FormKit
      type="form"
      id="create-recipe"
      :form-class="submitted ? 'hide' : 'show'"
      submit-label="Create"
      @submit="submitHandler"
      :actions="false"
      #default="{ value } "
   >

   
</template>

<script setup lang="ts">
import { FormKit } from '@formkit/vue';
import { ref } from 'vue'

const submitted = ref(false)
const submitHandler = async () => {
      await new Promise((r) => setTimeout(r, 2000))
      submitted.value = true
}

</script>