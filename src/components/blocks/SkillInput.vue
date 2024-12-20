<template>
	<div class="content__input-group">
		<div v-for="(field, index) in fields" :key="index" class="content__input-group-inner">
			<label :for="field.id">{{ field.label }}:</label>
			<input
						 type="text"
						 :name="field.name"
						 :id="field.id"
						 v-model="firstSection[field.name]"
						 @input="updateField(field.name, firstSection[field.name])"
						 class="input-field" />
		</div>
	</div>
</template>

<script>
import { mapMutations } from 'vuex'

export default {
	name: 'SkillInput',
	props: {
		fields: {
			type: Array,
			required: true
		},
		firstSection: {
			type: Object,
			required: true
		}
	},
	methods: {
		...mapMutations({
			updateField: 'UPDATE_FIELD' // Мутация для обновления поля
		}),
		updFirstSectionFromJSON(data) {
			// Обновляем firstSection на основе загруженного JSON
			Object.keys(data).forEach(key => {
				if (this.firstSection.hasOwnProperty(key)) {
					this.firstSection[key] = data[key] // Обновляем значение поля
					this.updateField(key, data[key]) // Вызываем мутацию для обновления в Store
				}
			})
		}
	}
}
</script>