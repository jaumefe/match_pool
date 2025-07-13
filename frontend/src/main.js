import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router/route.js'
import PrimeVue from 'primevue/config'

createApp(App)
.use(router)
.use(PrimeVue)
.mount('#app')
