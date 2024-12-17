import { createApp } from 'vue'
import App from './App.vue'
import components from './components'
import router from './router/router'
import { createStore } from 'vuex'

import '@/assets/styles.scss'

const app = createApp(App)

components.forEach(component => {
	app.component(component.name, component)
})

app
	.use(createStore())
	.use(router)
	.mount('#app')
