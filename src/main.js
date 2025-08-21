import { createApp } from 'vue'
import { createPinia } from 'pinia'
import router from './router'
import App from './App.vue'
import './style.css'

// Create Vue app instance
const app = createApp(App)

// Use Pinia for state management (Vue's equivalent to NgRx/Redux)
app.use(createPinia())

// Use Vue Router (similar to Angular Router)
app.use(router)

// Mount the app
app.mount('#app')
