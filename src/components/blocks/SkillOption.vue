<template>
	<div class="content__item-group">
		<div v-for="(item, index) in items" :key="index" class="content__item-group-inner">
			<span>{{ item.label }}</span>
			<label :for="item.name">{{ item.value }}</label>
			<select :name="item.name" :id="item.name" v-model="item.selected">
				<option value="">&nbsp;</option>
				<option value="\">\</option>
				<option value="x">x</option>
				<option value="Ж">ж</option>
			</select>
		</div>
	</div>
</template>

<script>
export default {
	name: 'SkillOption',
	data() {
		return {
			items: [
				{ label: 'Пустяк', value: '0', name: 'bruised', selected: '' },
				{ label: 'Боль', value: '-1', name: 'hurt', selected: '' },
				{ label: 'Потрёпан', value: '-1', name: 'injured', selected: '' },
				{ label: 'Травма', value: '-2', name: 'wounded', selected: '' },
				{ label: 'Истерзан', value: '-2', name: 'mauled', selected: '' },
				{ label: 'Увечья', value: '-5', name: 'crippled', selected: '' },
				{ label: 'Обморок', value: '', name: 'incapacitated', selected: '' },
			]
		}
	},
	methods: {
		handleSelectionChange(index) {
			// console.log(`Selected value for ${this.items[index].label}: ${this.items[index].selected}`)
		}
	}
}
</script>

<style></style>