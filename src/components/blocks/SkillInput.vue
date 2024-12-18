<template>
	<div class="content__input-group">
		<div v-for="(field, index) in fields" :key="index" class="content__input-group-inner">
			<label :for="field.id">{{ field.label }}:</label>
			<input type="text" :name="field.name" :id="field.id" v-model="formData[field.name]">
		</div>
	</div>
</template>

<script>
export default {
	name: 'SkillInput',
	props: {
		fields: {
			type: Array,
			required: true
		},
		formData: {
			type: Object,
			required: true
		}
	}
}
</script>

<style></style>