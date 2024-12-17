<template>
	<form class="modal__form modal__form--authorization" @submit.prevent="submit">
		<p class="modal__title">Авторизация</p>
		<div class="modal__group-input">
			<label for="login">Логин</label>
			<input type="text" name="login" id="login" v-model='data.login' placeholder="qwerty" required>
		</div>
		<div class="modal__group-input">
			<label for="password">Пароль</label>
			<input type="password" name="password" id="password" v-model='data.password' placeholder="********"
						 autocomplete='none' required>
		</div>
		<p class="modal__password-question">Восстановить пароль?</p>
		<button type="submit" class="modal__button modal__button--submit">Войти</button>
		<p class="modal__account-question">Нет аккаунта? <span class='register-question'
						@click='$emit("toggleForm")'>Зарегистрироваться</span></p>
	</form>
</template>

<script lang="ts">
import { reactive } from 'vue'
import { useRouter } from "vue-router"

export default {
	name: "Login",
	setup() {
		const data = reactive({
			login: '',
			password: ''
		})
		var resp = reactive({
			ok: false,
			statusText: ''
		})
		const router = useRouter()

		const submit = async () => {
			try {
				const response = await fetch('http://localhost:8000/api/login', {
					method: 'POST',
					headers: { 'Content-Type': 'X-www-form-urlencoded' },
					credentials: 'include',
					body: JSON.stringify(data)
				})
				resp = await response.json()
				console.log(resp.ok)
				console.log(resp.statusText)
				if (!resp.ok) {
					throw new Error('Ошибка при входе: ' + resp.statusText)
				}
				await location.reload()
			} catch (error) {
				console.error('Ошибка:', error)
			}
		}

		return {
			data,
			submit
		}
	}
}
</script>

<style lang="">
</style>