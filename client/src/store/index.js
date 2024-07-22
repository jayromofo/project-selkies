import { defineStore } from "pinia"
import { ref } from 'vue'

// Need a user store that has a userId so it can access the data from database regarding
// data stores
export const useUserStore = defineStore("user", () => {
   const userId = ref('');
   const userName = ref('');
   const userEmail = ref('');

   // Computed  ex: const doubleCount = computed (() => count.value * 2)


   // Functions


   // Return (Everything important)
   return {
      userId, 
      userName, 
      userEmail

   }
})

// Stores the user settings 
export const useUserSettingsStore = defineStore('userSettings', () => {
   const userId = ref('')
   const userName = ref('')
   const darkMode = ref(false)


   return {
      userId,
      userName,
      darkMode,

   }
})
