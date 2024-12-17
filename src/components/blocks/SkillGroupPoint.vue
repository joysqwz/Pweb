<template>
	<RadioButton
							 v-for="value in values"
							 :key="value"
							 :value="value"
							 :name="groupName"
							 v-model="selectedValue" />
</template>

<script>
export default {
	name: "SkillGroupPoint",
	props: {
		groupName: {
			type: String,
			required: true,
		},
		values: {
			type: Array,
			required: true,
		},
	},
	data() {
		return {
			selectedValue: null,
		}
	},
	methods: {
		clearSelection() {
			this.selectedValue = null
		}
	}
}
</script>

<style></style>
