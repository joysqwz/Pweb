<template>
	<input
				 type="radio"
				 :value="value"
				 :name="name"
				 :id="id"
				 :checked="modelValue === value"
				 @change="updateValue" />
	<label :for="id"></label>
</template>

<script>
export default {
	name: 'RadioButton',
	props: {
		value: {
			type: Number,
			required: true
		},
		name: {
			type: String,
			required: true
		},
		modelValue: {
			type: [String, Number],
			default: null,
		}
	},
	computed: {
		id() {
			return `${this.name}-${this.value}`
		}
	},
	methods: {
		updateValue() {
			this.$emit('update:modelValue', this.value)
		}
	}
}
</script>