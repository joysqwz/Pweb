<template>
	<img class="content__logo" src="/images/logoMtA.svg" alt="" width="220" height="110" loading="lazy">
	<div class="content__page" id="page1">
		<div class="content__section">
			<SkillInput :fields="firstGroupFields" :formData="formData" />
			<SkillInput :fields="secondGroupFields" :formData="formData" />
			<SkillInput :fields="thirdGroupFields" :formData="formData" />
		</div>
		<h2>Атрибуты</h2>
		<div class="content__section">
			<div class="content__item-group">
				<h3>Физические</h3>
				<SkillLabelPoint skillName="Сила" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Ловкость" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Выносливость" :skillValues="[5, 4, 3, 2, 1]" />
			</div>
			<div class="content__item-group">
				<h3>Социальные</h3>
				<SkillLabelPoint skillName="Обаяние" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Манипулирование" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Внешность" :skillValues="[5, 4, 3, 2, 1]" />
			</div>
			<div class="content__item-group">
				<h3>Ментальные</h3>
				<SkillLabelPoint skillName="Восприятие" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Интеллект" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Смекалка" :skillValues="[5, 4, 3, 2, 1]" />
			</div>
		</div>
		<h2>Способности</h2>
		<div class="content__section">
			<div class="content__item-group">
				<h3>Таланты</h3>
				<SkillLabelPoint skillName="Атлетика" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Бдительность" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Драка" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Запугивание" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Искусство" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Красноречие" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Лидерство" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Уличные порядки" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Хитрость" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Экстрасенсорика" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Эмпатия" :skillValues="[5, 4, 3, 2, 1]" />
			</div>
			<div class="content__item-group">
				<h3>Навыки</h3>
				<SkillLabelPoint skillName="Боевые искусства" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Вождение" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Выживание" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Исследование" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Медитация" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Огнестрельное оружие" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Ремёсла" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Скрытность" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Технологии" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Холодное оружие" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Этикет" :skillValues="[5, 4, 3, 2, 1]" />
			</div>
			<div class="content__item-group">
				<h3>Познания</h3>
				<SkillLabelPoint skillName="Гуманитарные науки" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Естественные науки" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Загадки" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Законы" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Компьютеры" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Космология" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Медицина" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Оккультизм" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Политика" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Расследование" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Эзотерика" :skillValues="[5, 4, 3, 2, 1]" />
			</div>
		</div>
		<h2>Сферы</h2>
		<div class="content__section">
			<div class="content__item-group">
				<SkillLabelPoint skillName="Жизнь" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Материя" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Силы" :skillValues="[5, 4, 3, 2, 1]" />
			</div>
			<div class="content__item-group">
				<SkillLabelPoint skillName="Время" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Связи" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Энтропия" :skillValues="[5, 4, 3, 2, 1]" />
			</div>
			<div class="content__item-group">
				<SkillLabelPoint skillName="Дух" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Разум" :skillValues="[5, 4, 3, 2, 1]" />
				<SkillLabelPoint skillName="Основы" :skillValues="[5, 4, 3, 2, 1]" />
			</div>
		</div>
		<h2>Преимущества</h2>
		<div class="content__section content__section--ai-start">
			<div class="content__item-group content__item-group--gap4">
				<h3>Дополнения</h3>
				<div
						 class="content__item-group-inner content__item-group-inner--input-ht20 content__item-group-inner--gap30">
					<SkillInputPoint inputName="backgrounds-one" />
				</div>
				<div
						 class="content__item-group-inner content__item-group-inner--input-ht20 content__item-group-inner--gap30">
					<SkillInputPoint inputName="backgrounds-two" />
				</div>
				<div
						 class="content__item-group-inner content__item-group-inner--input-ht20 content__item-group-inner--gap30">
					<SkillInputPoint inputName="backgrounds-three" />
				</div>
				<div
						 class="content__item-group-inner content__item-group-inner--input-ht20 content__item-group-inner--gap30">
					<SkillInputPoint inputName="backgrounds-four" />
				</div>
				<div
						 class="content__item-group-inner content__item-group-inner--input-ht20 content__item-group-inner--gap30">
					<SkillInputPoint inputName="backgrounds-five" />
				</div>
				<div
						 class="content__item-group-inner content__item-group-inner--input-ht20 content__item-group-inner--gap30 content__item-group-inner--mb20">
					<SkillInputPoint inputName="backgrounds-six" />
				</div>
				<h3>Прочие особенности</h3>
				<div
						 class="content__item-group-inner content__item-group-inner--input-ht20 content__item-group-inner--gap30">
					<SkillInputPoint inputName="other-traits-one" />
				</div>
				<div
						 class="content__item-group-inner content__item-group-inner--input-ht20 content__item-group-inner--gap30">
					<SkillInputPoint inputName="other-traits-two" />
				</div>
				<div
						 class="content__item-group-inner content__item-group-inner--input-ht20 content__item-group-inner--gap30">
					<SkillInputPoint inputName="other-traits-three" />
				</div>
				<div
						 class="content__item-group-inner content__item-group-inner--input-ht20 content__item-group-inner--gap30">
					<SkillInputPoint inputName="other-traits-four" />
				</div>
				<div
						 class="content__item-group-inner content__item-group-inner--input-ht20 content__item-group-inner--gap30">
					<SkillInputPoint inputName="other-traits-five" />
				</div>
				<div
						 class="content__item-group-inner content__item-group-inner--input-ht20 content__item-group-inner--gap30">
					<SkillInputPoint inputName="other-traits-six" />
				</div>
			</div>
			<div class="content__item-group content__item-group--gap0">
				<h3>Арете</h3>
				<div class="content__item-group-inner content__item-group-inner--fd-col content__item-group-inner--mb20">
					<div class="content__skill-level">
						<SkillGroupPoint
														 ref="skillGroup1"
														 groupName="arete"
														 :values="[10, 9, 8, 7, 6, 5, 4, 3, 2, 1]" />
					</div>
					<button class='content_skill-clear-button content_skill-clear-button--left20'
									@click="clearSelection('skillGroup1')">
						<CrossSVG></CrossSVG>
					</button>
				</div>
				<h3>Сила воли</h3>
				<div class="content__item-group-inner content__item-group-inner--fd-col content__item-group-inner--mb20">
					<div class="content__skill-level">
						<SkillGroupPoint
														 ref="skillGroup2"
														 groupName="willpower"
														 :values="[10, 9, 8, 7, 6, 5, 4, 3, 2, 1]" />
					</div>
					<div class="content__skill-level content__skill-level--content-X">
						<SkillGroupPoint
														 ref="skillGroup2_2"
														 groupName="willpower-2"
														 :values="[10, 9, 8, 7, 6, 5, 4, 3, 2, 1]" />

					</div>
					<div class="content__skill-level content__skill-level--content-block">
						<input value="10" name="willpower-2-alt" id="willpower-2-alt-10" type="radio">
						<label for="willpower-2-alt-10"></label>
						<input value="9" name="willpower-2-alt" id="willpower-2-alt-9" type="radio">
						<label for="willpower-2-alt-9"></label>
						<input value="8" name="willpower-2-alt" id="willpower-2-alt-8" type="radio">
						<label for="willpower-2-alt-8"></label>
						<input value="7" name="willpower-2-alt" id="willpower-2-alt-7" type="radio">
						<label for="willpower-2-alt-7"></label>
						<input value="6" name="willpower-2-alt" id="willpower-2-alt-6" type="radio">
						<label for="willpower-2-alt-6"></label>
						<input value="5" name="willpower-2-alt" id="willpower-2-alt-5" type="radio">
						<label for="willpower-2-alt-5"></label>
						<input value="4" name="willpower-2-alt" id="willpower-2-alt-4" type="radio">
						<label for="willpower-2-alt-4"></label>
						<input value="3" name="willpower-2-alt" id="willpower-2-alt-3" type="radio">
						<label for="willpower-2-alt-3"></label>
						<input value="2" name="willpower-2-alt" id="willpower-2-alt-2" type="radio">
						<label for="willpower-2-alt-2"></label>
						<input value="1" name="willpower-2-alt" id="willpower-2-alt-1" type="radio">
						<label for="willpower-2-alt-1"></label>
					</div>
					<button class='content_skill-clear-button content_skill-clear-button--left20'
									@click="clearGroup2Selections">
						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner content__item-group-inner--fd-col">
					<h3>Квинтэссенция</h3>
					<SkillCircle />
					<h3>Парадокс</h3>
				</div>
			</div>
			<div class="content__item-group">
				<h3>Здоровье</h3>
				<SkillOption />
				<div class="content__item-group-inner content__item-group-inner--fd-col">
					<h3>Опыт</h3>
					<textarea
										name="experience"
										id="experience"
										rows="3"
										cols="22"
										v-model="experienceText">
    			</textarea>
				</div>
			</div>
		</div>
	</div>
</template>
<script>
export default {
	name: 'PageOne',
	data() {
		return {
			experienceText: '',
			formData: {
				attr_Name: '',
				attr_Player: '',
				attr_Chronicle: '',
				attr_Nature: '',
				attr_Mask: '',
				attr_Energy: '',
				attr_Fraction: '',
				attr_Sect: '',
				attr_Concept: ''
			},
			firstGroupFields: [
				{ label: 'Имя', name: 'attr_Name', id: 'attr_Name' },
				{ label: 'Игрок', name: 'attr_Player', id: 'attr_Player' },
				{ label: 'Хроника', name: 'attr_Chronicle', id: 'attr_Chronicle' }
			],
			secondGroupFields: [
				{ label: 'Натура', name: 'attr_Nature', id: 'attr_Nature' },
				{ label: 'Маска', name: 'attr_Mask', id: 'attr_Mask' },
				{ label: 'Энергия', name: 'attr_Energy', id: 'attr_Energy' }
			],
			thirdGroupFields: [
				{ label: 'Фракция', name: 'attr_Fraction', id: 'attr_Fraction' },
				{ label: 'Секта', name: 'attr_Sect', id: 'attr_Sect' },
				{ label: 'Концепция', name: 'attr_Concept', id: 'attr_Concept' }
			]
		}
	},
	methods: {
		clearSelection(skillGroupRef) {
			this.$refs[skillGroupRef].clearSelection()
		},
		clearGroup2Selections() {
			this.$refs.skillGroup2.clearSelection()
			this.$refs.skillGroup2_2.clearSelection()
		}
	}
}
</script>
<style></style>