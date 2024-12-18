<template>
	<dialog id="modal" class="user-profile__modal">
		<div v-if='showLogin' class="modal__wrapper modal__wrapper--authorization">
			<Login @toggleForm="toggleForm" />
		</div>
		<div v-if='!showLogin' class="modal__wrapper modal__wrapper--registration">
			<Register @toggleForm="toggleForm" />
		</div>
	</dialog>
</template>

<script>
import { onMounted, ref } from 'vue'
import store from './store/store'
import { reactive } from 'vue'

export default {
	name: 'Modal',
	setup() {
		const message = ref('Вы не вошли в систему')
		var resp = reactive({
			login: '',
			found: false,
		})
		onMounted(async () => {
			try {
				const response = await fetch('http://localhost:8000/api/user', {
					method: 'POST',
					headers: { 'Content-Type': 'X-www-form-urlencoded' },
					credentials: 'include'
				},)
				resp = await response.json()
				message.value = resp.login
				console.log(resp.found)
				await store.dispatch('setAuth', resp.found)
				await store.dispatch('setLogin', resp.login)

			}
			catch (e) {
				await store.dispatch('setAuth', false)
			}
		})
		return {
			message
		}
	},
	data() {
		return {
			showLogin: true,
		}
	},
	methods: {
		toggleForm() {
			this.showLogin = !this.showLogin
		}
	},
}
</script>

<style lang="">
</style>