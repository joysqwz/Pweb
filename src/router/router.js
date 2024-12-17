import { createRouter, createWebHistory } from "vue-router"
import Storage from '@/pages/Storage.vue'
import Main from '@/pages/Main.vue'

const routes = [
	{
		path: '/',
		base: '',
		component: Main
	},
	{
		path: '/storage',
		component: Storage
	},
]

const router = createRouter({
	routes,
	history: createWebHistory()
})

export default router
