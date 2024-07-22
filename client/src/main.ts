import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { plugin, defaultConfig } from '@formkit/vue';

import App from './App.vue'
import router from './router'
import PrimeVue from 'primevue/config' 



const pinia = createPinia();
const app = createApp(App);
const app = createApp(App)


app.use(PrimeVue, {
   unstyled: true,
})

app.use(router)
app.use(pinia);
app.use(plugin, defaultConfig);


app.mount('#app')
