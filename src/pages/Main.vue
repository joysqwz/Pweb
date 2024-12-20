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
				<PageOne v-if="activePage === 1" ref="pageOne" :firstSection="firstSection" />
			</div>
			<div class="content__menu" v-if='authenticated'>
				<button type="button" class='content__button content__button--alt' @click="setSavePage">Сохранить</button>
				<button type="button" class='content__button content__button--alt' @click="triggerFileInput">Загрузить</button>
				<input type="file" @change="handleFileUpload" accept=".json" style="display: none;" ref="fileInput" />
			</div>
		</div>
	</main>
	<Footer></Footer>
</template>

<script>
import { mapState } from 'vuex'

export default {
	computed: {
		...mapState(['authenticated']),
	},
	data() {
		return {
			activePage: 1,
			saveStates: {
				1: false,
			},
			firstSection: {
				attr_Name: '',
				attr_Player: '',
				attr_Chronicle: '',
				attr_Nature: '',
				attr_Mask: '',
				attr_Energy: '',
				attr_Fraction: '',
				attr_Sect: '',
				attr_Concept: ''
			}
		}
	},
	methods: {
		setActivePage(page) {
			this.activePage = page
		},
		setSavePage() {
			const jsonData = JSON.stringify(this.firstSection, null, 2) // Используйте firstSection для актуальных данных
			const blob = new Blob([jsonData], { type: 'application/json' })
			const url = URL.createObjectURL(blob)
			const a = document.createElement('a')
			a.href = url
			a.download = 'list.json'
			document.body.appendChild(a)
			a.click()
			document.body.removeChild(a)
			URL.revokeObjectURL(url)
			this.saveStates[this.activePage] = true
			console.log(jsonData)
		},
		triggerFileInput() {
			this.$refs.fileInput.click()
		},
		handleFileUpload(event) {
			const file = event.target.files[0]
			if (file) {
				const reader = new FileReader()
				reader.onload = (e) => {
					try {
						const jsonData = JSON.parse(e.target.result)
						console.log('Загруженные данные:', jsonData) // Проверка загруженных данных
						this.$refs.pageOne.updFirstSectionFromJSON(jsonData) // Обновляем данные в PageOne
						this.$store.commit('UPDATE_FIRST_SECTION', jsonData) // Обновление данных в Vuex
						console.log('Состояние firstSection после обновления:', this.$store.state.firstSection) // Проверка состояния
						this.saveStates[this.activePage] = false
					} catch (error) {
						console.error('Ошибка при парсинге JSON:', error)
					}
				}
				reader.readAsText(file)
			}
		}
	}
}
</script>

<style></style>