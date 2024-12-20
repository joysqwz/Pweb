import { createApp } from 'vue'
import App from './App.vue'
import components from './components'
import router from './router/router'
import store from './components/store/store' // Импортируйте ваше хранилище

import '@/assets/styles.scss'

const app = createApp(App)

components.forEach(component => {
	app.component(component.name, component)
})

app
	.use(store)
	.use(router)
	.mount('#app')