<template>
	<form class="modal__form modal__form--registration" @submit.prevent="submit">
		<p class="modal__title">Регистрация</p>
		<div class="modal__group-input">
			<label for="reg-login">Логин</label>
			<input type="text" name="regLogin" id="reg-login" v-model="data.login" placeholder="qwerty" required>
		</div>
		<div class="modal__group-input">
			<label for="email">Адрес электронной почты</label>
			<input type="email" name="email" id="email" v-model="data.email" placeholder="qwerty@mail.com"
						 autocomplete='email' required>
		</div>
		<div class="modal__group-input">
			<label for="reg-password">Пароль</label>
			<input type="password" name="regPassword" id="reg-password" v-model="data.password" placeholder="********"
						 autocomplete='new-password' required>
		</div>
		<button type="submit" class="modal__button modal__button--submit">Зарегистрироваться</button>
		<p class="modal__account-question">Есть аккаунт? <span class='login-question'
						@click='$emit("toggleForm")'>Войти</span></p>
	</form>
</template>

<script lang="ts">
import { reactive } from 'vue'

export default {
	name: "Register",
	setup() {
		const data = reactive({
			login: '',
			email: '',
			password: ''
		})

		const submit = async () => {
			try {
				const response = await fetch('http://localhost:8000/api/register', {
					method: 'POST',
					headers: { 'Content-Type': 'application/json' },
					body: JSON.stringify(data)
				})

				if (!response.ok) {
					throw new Error('Ошибка при регистрации: ' + response.statusText)
				}
				console.log('Регистрация успешна!')
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