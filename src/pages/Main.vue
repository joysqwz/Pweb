<template>
	<Header></Header>
	<main class="content">
		<div class="content__body container container--width800">
			<div class="content__navigation">
				<button
								type="button"
								class="content__button"
								id="btn-one"
								@click="setActivePage(1)"
								:class="{
									'content__button--active': activePage === 1,
									'content__button--save': saveStates[1]
								}">
					Лист 1
				</button>
			</div>
			<div class="content__wrapper">
				<PageOne v-if="activePage === 1"></PageOne>
			</div>
			<div class="content__menu" v-if='authenticated'>
				<button type="button" class='content__button content__button--alt' @click="setSavePage">Сохранить</button>
				<button type="button" class='content__button content__button--alt' @click="exportJSON">Загрузить</button>
			</div>
		</div>
	</main>
	<Footer></Footer>
</template>

<script>
import { mapState } from 'vuex'
export default {
	computed: {
		...mapState([
			'authenticated',
		]),
	},
	data() {
		return {
			activePage: 1,
			saveStates: {
				1: false,
			}
		}
	},
	methods: {
		setActivePage(page) {
			this.activePage = page
		},
		setSavePage(page) {
			this.saveStates[this.activePage] = true
		},
		exportJSON() {

		}
	}
}
</script>

<style></style>