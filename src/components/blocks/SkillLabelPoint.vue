<template>
	<div class="content__item-group-inner">
		<h4>{{ skillName }}</h4>
		<div class="content__skill-level">
			<RadioButton
									 v-for="value in skillValues"
									 :key="value"
									 :value="value"
									 :name="skillName"
									 v-model="selectedValue" />
		</div>
		<button class='content_skill-clear-button' :data-group="skillName" @click="clearSelection">
			<CrossSVG />
		</button>
	</div>
</template>
<script>
export default {
	name: 'SkillLabelPoint',
	props: {
		skillName: {
			type: String,
			required: true
		},
		skillValues: {
			type: Array,
			required: true
		}
	},
	data() {
		return {
			selectedValue: null
		}
	},
	methods: {
		clearSelection() {
			this.selectedValue = null
		}
	}
}
</script>