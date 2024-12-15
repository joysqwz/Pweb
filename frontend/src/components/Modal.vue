<template>
	<dialog id="modal" class="user-profile__modal">
		<div v-if='showModal'
				 class="modal__wrapper modal__wrapper--authorization">
			<p class="modal__title">Авторизация</p>
			<form class="modal__form modal__form--authorization">
				<div class="modal__group-input">
					<label for="login">Логин</label>
					<input type="text" name="login" id="login" placeholder="qwerty">
				</div>
				<div class="modal__group-input">
					<label for="password">Пароль</label>
					<input type="password" name="password" id="password" placeholder="********" autocomplete='none'>
				</div>
				<p class="modal__password-question">Восстановить пароль?</p>
				<button class="modal__button modal__button--submit">Войти</button>
				<p class="modal__account-question">Нет аккаунта? <span
								class='register-question' @click='changeShow'>Зарегистрироваться</span>
				</p>
			</form>
		</div>
		<div v-if='!showModal'
				 class="modal__wrapper modal__wrapper--registration">
			<p class="modal__title">Регистрация</p>
			<form class="modal__form modal__form--registration">
				<div class="modal__group-input">
					<label for="reg-login">Логин</label>
					<input type="text" name="regLogin" id="reg-login" placeholder="qwerty">
				</div>
				<div class="modal__group-input">
					<label for="email">Адрес электронной почты</label>
					<input type="email" name="email" id="email" placeholder="qwerty@mail.com" autocomplete='email'>
				</div>
				<div class="modal__group-input">
					<label for="reg-password">Пароль</label>
					<input type="password" name="regPassword" id="reg-password" placeholder="********"
								 autocomplete='new-password'>
				</div>
				<div class="modal__group-input">
					<label for="reg-conf-password">Подтверждение пароля</label>
					<input type="password" name="regCPassword" id="reg-conf-password" placeholder="********"
								 autocomplete='new-password'>
				</div>
				<button class="modal__button modal__button--submit">Зарегистрироваться</button>
				<p class="modal__account-question">Есть аккаунт? <span class='login-question' @click='changeShow'>Войти</span>
				</p>
			</form>
		</div>
	</dialog>
</template>
<script>
export default {
	name: 'Modal',
	data() {
		return {
			showModal: true,
		}
	},
	methods: {
		changeShow() {
			this.showModal = !this.showModal
		}
	}
}
</script>
<style lang="">

</style>