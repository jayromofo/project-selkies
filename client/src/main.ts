import 'vuestic-ui/styles/essential.css';
import 'vuestic-ui/styles/typography.css';
import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { plugin, defaultConfig } from '@formkit/vue';

import { createVuestic } from 'vuestic-ui';
import config from '../vuestic.config'

// import "vuestic-ui/css"

import App from './App.vue'
import router from './router'



const pinia = createPinia();
const app = createApp(App);

app.use(router)
app.use(pinia);
app.use(plugin, defaultConfig);
app.use(createVuestic({config}));


app.mount('#app')
