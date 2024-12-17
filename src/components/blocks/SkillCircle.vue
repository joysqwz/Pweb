<template>
	<div class="circle-container">
		<div class="dot"></div>
		<button class='content_skill-clear-button content_skill-clear-button--center-circle' @click="clearAllSelections">
			<CrossSVG />
		</button>

		<div v-for="(item, index) in checkboxes" :key="index" class="checkbox-wrapper">
			<input type="checkbox"
						 :id="'check' + (index + 1)"
						 v-model="item.checked"
						 @change="trackCheckbox(index)">
			<label :for="'check' + (index + 1)"
						 :class="item.class"
						 @dblclick="toggleClass(index)">
			</label>
		</div>
	</div>
</template>

<script>
export default {
	name: 'SkillCircle',
	data() {
		return {
			checkboxes: [
				{ checked: false, class: 'pos-1 o' },
				{ checked: false, class: 'pos-2 o' },
				{ checked: false, class: 'pos-3 o' },
				{ checked: false, class: 'pos-4 o' },
				{ checked: false, class: 'pos-5 o' },
				{ checked: false, class: 'pos-6 x' },
				{ checked: false, class: 'pos-7 x' },
				{ checked: false, class: 'pos-8 x' },
				{ checked: false, class: 'pos-9 x' },
				{ checked: false, class: 'pos-10 x' },
				{ checked: false, class: 'pos-11 x' },
				{ checked: false, class: 'pos-12 x' },
				{ checked: false, class: 'pos-13 x' },
				{ checked: false, class: 'pos-14 x' },
				{ checked: false, class: 'pos-15 x' },
				{ checked: false, class: 'pos-16 o' },
				{ checked: false, class: 'pos-17 o' },
				{ checked: false, class: 'pos-18 o' },
				{ checked: false, class: 'pos-19 o' },
				{ checked: false, class: 'pos-20 o' },
			],
			activeCheckboxes: []
		}
	},
	methods: {
		clearAllSelections() {
			this.checkboxes.forEach(checkbox => {
				checkbox.checked = false
				checkbox.class = checkbox.class.split(' ')[0] + ' o'
			})
			this.activeCheckboxes = []
		},
		trackCheckbox(index) {
			console.log(`Checkbox ${index + 1} was clicked. Checked: ${this.checkboxes[index].checked}`)
			this.updateActiveCheckboxes()
		},
		toggleClass(index) {
			const currentClass = this.checkboxes[index].class
			const positionClass = currentClass.split(' ')[0]
			let newClass

			if (currentClass.includes('x')) {
				newClass = `${positionClass} o`
			} else {
				newClass = `${positionClass} x`
			}

			this.checkboxes[index].class = newClass
			this.updateActiveCheckboxes()
		},
		updateActiveCheckboxes() {
			this.activeCheckboxes = this.checkboxes
				.map((checkbox, index) => ({
					index: index,
					state: checkbox.checked ? (checkbox.class.includes('x') ? 'X' : 'O') : null
				}))
				.filter(active => active.state !== null)
		}
	}
}
</script>

<style></style>