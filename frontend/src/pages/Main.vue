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
				<button
								type="button"
								class="content__button"
								id="btn-two"
								@click="setActivePage(2)"
								:class="{
									'content__button--active': activePage === 2,
									'content__button--save': saveStates[2]
								}">
					Лист 2
				</button>
				<button
								type="button"
								class="content__button"
								id="btn-three"
								@click="setActivePage(3)"
								:class="{
									'content__button--active': activePage === 3,
									'content__button--save': saveStates[3]
								}">
					Лист 3
				</button>
				<button
								type="button"
								class="content__button"
								id="btn-four"
								@click="setActivePage(4)"
								:class="{
									'content__button--active': activePage === 4,
									'content__button--save': saveStates[4]
								}">
					Лист 4
				</button>
			</div>
			<div class="content__wrapper">
				<PageOne v-if="activePage === 1"></PageOne>
				<PageTwo v-if="activePage === 2"></PageTwo>
				<PageThree v-if="activePage === 3"></PageThree>
				<PageFour v-if="activePage === 4"></PageFour>
			</div>
			<button type="button" @click="setSavePage">SAVE</button>
		</div>
	</main>
	<Footer></Footer>
</template>

<script>
export default {
	data() {
		return {
			activePage: 1,
			saveStates: {
				1: false,
				2: false,
				3: false,
				4: false,
			}
		}
	},
	methods: {
		setActivePage(page) {
			this.activePage = page
		},
		setSavePage(page) {
			this.saveStates[this.activePage] = true
		}
	}
}
</script>

<style></style>