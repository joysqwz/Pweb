<template>
	<img class="content__logo" src="/images/logoMtA.svg" alt="" width="220" height="110" loading="lazy">
	<div class="content__page" id="page1">
		<div class="content__section">
			<div class="content__input-group">
				<div v-for="(field, index) in firstGroupFields" :key="index" class="content__input-group-inner">
					<label :for="field.id">{{ field.label }}:</label>
					<input type="text" :name="field.name" :id="field.id" v-model="formData[field.name]">
				</div>
			</div>
			<div class="content__input-group">
				<div v-for="(field, index) in secondGroupFields" :key="index" class="content__input-group-inner">
					<label :for="field.id">{{ field.label }}:</label>
					<input type="text" :name="field.name" :id="field.id" v-model="formData[field.name]">
				</div>
			</div>
			<div class="content__input-group">
				<div v-for="(field, index) in thirdGroupFields" :key="index" class="content__input-group-inner">
					<label :for="field.id">{{ field.label }}:</label>
					<input type="text" :name="field.name" :id="field.id" v-model="formData[field.name]">
				</div>
			</div>
		</div>
		<h2>Атрибуты</h2>
		<div class="content__section">
			<div class="content__item-group">
				<h3>Физические</h3>
				<div class="content__item-group-inner">
					<h4>Сила</h4>
					<div class="content__skill-level">
						<input value="5" name="strength" id="strength-5" type="radio">
						<label for="strength-5"></label>
						<input value="4" name="strength" id="strength-4" type="radio">
						<label for="strength-4"></label>
						<input value="3" name="strength" id="strength-3" type="radio">
						<label for="strength-3"></label>
						<input value="2" name="strength" id="strength-2" type="radio">
						<label for="strength-2"></label>
						<input value="1" name="strength" id="strength-1" type="radio">
						<label for="strength-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="strength">
						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Ловкость</h4>
					<div class="content__skill-level">
						<input value="5" name="dexterity" id="dexterity-5" type="radio">
						<label for="dexterity-5"></label>
						<input value="4" name="dexterity" id="dexterity-4" type="radio">
						<label for="dexterity-4"></label>
						<input value="3" name="dexterity" id="dexterity-3" type="radio">
						<label for="dexterity-3"></label>
						<input value="2" name="dexterity" id="dexterity-2" type="radio">
						<label for="dexterity-2"></label>
						<input value="1" name="dexterity" id="dexterity-1" type="radio">
						<label for="dexterity-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="dexterity">
						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Выносливость</h4>
					<div class="content__skill-level">
						<input value="5" name="stamina" id="stamina-5" type="radio">
						<label for="stamina-5"></label>
						<input value="4" name="stamina" id="stamina-4" type="radio">
						<label for="stamina-4"></label>
						<input value="3" name="stamina" id="stamina-3" type="radio">
						<label for="stamina-3"></label>
						<input value="2" name="stamina" id="stamina-2" type="radio">
						<label for="stamina-2"></label>
						<input value="1" name="stamina" id="stamina-1" type="radio">
						<label for="stamina-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="stamina">
						<CrossSVG></CrossSVG>
					</button>
				</div>
			</div>
			<div class="content__item-group">
				<h3>Социальные</h3>
				<div class="content__item-group-inner">
					<h4>Обаяние</h4>
					<div class="content__skill-level">
						<input value="5" name="charisma" id="charisma-5" type="radio">
						<label for="charisma-5"></label>
						<input value="4" name="charisma" id="charisma-4" type="radio">
						<label for="charisma-4"></label>
						<input value="3" name="charisma" id="charisma-3" type="radio">
						<label for="charisma-3"></label>
						<input value="2" name="charisma" id="charisma-2" type="radio">
						<label for="charisma-2"></label>
						<input value="1" name="charisma" id="charisma-1" type="radio">
						<label for="charisma-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="charisma">
						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Манипулирование</h4>
					<div class="content__skill-level">
						<input value="5" name="manipulation" id="manipulation-5" type="radio">
						<label for="manipulation-5"></label>
						<input value="4" name="manipulation" id="manipulation-4" type="radio">
						<label for="manipulation-4"></label>
						<input value="3" name="manipulation" id="manipulation-3" type="radio">
						<label for="manipulation-3"></label>
						<input value="2" name="manipulation" id="manipulation-2" type="radio">
						<label for="manipulation-2"></label>
						<input value="1" name="manipulation" id="manipulation-1" type="radio">
						<label for="manipulation-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="manipulation">
						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Внешность</h4>
					<div class="content__skill-level">
						<input value="5" name="appearance" id="appearance-5" type="radio">
						<label for="appearance-5"></label>
						<input value="4" name="appearance" id="appearance-4" type="radio">
						<label for="appearance-4"></label>
						<input value="3" name="appearance" id="appearance-3" type="radio">
						<label for="appearance-3"></label>
						<input value="2" name="appearance" id="appearance-2" type="radio">
						<label for="appearance-2"></label>
						<input value="1" name="appearance" id="appearance-1" type="radio">
						<label for="appearance-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="appearance">
						<CrossSVG></CrossSVG>
					</button>
				</div>
			</div>
			<div class="content__item-group">
				<h3>Ментальные</h3>
				<div class="content__item-group-inner">
					<h4>Восприятие</h4>
					<div class="content__skill-level">
						<input value="5" name="perception" id="perception-5" type="radio">
						<label for="perception-5"></label>
						<input value="4" name="perception" id="perception-4" type="radio">
						<label for="perception-4"></label>
						<input value="3" name="perception" id="perception-3" type="radio">
						<label for="perception-3"></label>
						<input value="2" name="perception" id="perception-2" type="radio">
						<label for="perception-2"></label>
						<input value="1" name="perception" id="perception-1" type="radio">
						<label for="perception-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="perception">
						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Интеллект</h4>
					<div class="content__skill-level">
						<input value="5" name="intelligence" id="intelligence-5" type="radio">
						<label for="intelligence-5"></label>
						<input value="4" name="intelligence" id="intelligence-4" type="radio">
						<label for="intelligence-4"></label>
						<input value="3" name="intelligence" id="intelligence-3" type="radio">
						<label for="intelligence-3"></label>
						<input value="2" name="intelligence" id="intelligence-2" type="radio">
						<label for="intelligence-2"></label>
						<input value="1" name="intelligence" id="intelligence-1" type="radio">
						<label for="intelligence-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="intelligence">
						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Смекалка</h4>
					<div class="content__skill-level">
						<input value="5" name="wits" id="wits-5" type="radio">
						<label for="wits-5"></label>
						<input value="4" name="wits" id="wits-4" type="radio">
						<label for="wits-4"></label>
						<input value="3" name="wits" id="wits-3" type="radio">
						<label for="wits-3"></label>
						<input value="2" name="wits" id="wits-2" type="radio">
						<label for="wits-2"></label>
						<input value="1" name="wits" id="wits-1" type="radio">
						<label for="wits-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="wits">
						<CrossSVG></CrossSVG>
					</button>
				</div>
			</div>
		</div>
		<h2>Способности</h2>
		<div class="content__section">
			<div class="content__item-group">
				<h3>Таланты</h3>
				<div class="content__item-group-inner">
					<h4>Атлетика</h4>
					<div class="content__skill-level">
						<input value="5" name="athletics" id="athletics-5" type="radio">
						<label for="athletics-5"></label>
						<input value="4" name="athletics" id="athletics-4" type="radio">
						<label for="athletics-4"></label>
						<input value="3" name="athletics" id="athletics-3" type="radio">
						<label for="athletics-3"></label>
						<input value="2" name="athletics" id="athletics-2" type="radio">
						<label for="athletics-2"></label>
						<input value="1" name="athletics" id="athletics-1" type="radio">
						<label for="athletics-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="athletics">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Бдительность</h4>
					<div class="content__skill-level">
						<input value="5" name="alertness" id="alertness-5" type="radio">
						<label for="alertness-5"></label>
						<input value="4" name="alertness" id="alertness-4" type="radio">
						<label for="alertness-4"></label>
						<input value="3" name="alertness" id="alertness-3" type="radio">
						<label for="alertness-3"></label>
						<input value="2" name="alertness" id="alertness-2" type="radio">
						<label for="alertness-2"></label>
						<input value="1" name="alertness" id="alertness-1" type="radio">
						<label for="alertness-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="alertness">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Драка</h4>
					<div class="content__skill-level">
						<input value="5" name="brawl" id="brawl-5" type="radio">
						<label for="brawl-5"></label>
						<input value="4" name="brawl" id="brawl-4" type="radio">
						<label for="brawl-4"></label>
						<input value="3" name="brawl" id="brawl-3" type="radio">
						<label for="brawl-3"></label>
						<input value="2" name="brawl" id="brawl-2" type="radio">
						<label for="brawl-2"></label>
						<input value="1" name="brawl" id="brawl-1" type="radio">
						<label for="brawl-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="brawl">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Запугивание</h4>
					<div class="content__skill-level">
						<input value="5" name="intimidation" id="intimidation-5" type="radio">
						<label for="intimidation-5"></label>
						<input value="4" name="intimidation" id="intimidation-4" type="radio">
						<label for="intimidation-4"></label>
						<input value="3" name="intimidation" id="intimidation-3" type="radio">
						<label for="intimidation-3"></label>
						<input value="2" name="intimidation" id="intimidation-2" type="radio">
						<label for="intimidation-2"></label>
						<input value="1" name="intimidation" id="intimidation-1" type="radio">
						<label for="intimidation-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="intimidation">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Искусство</h4>
					<div class="content__skill-level">
						<input value="5" name="art" id="art-5" type="radio">
						<label for="art-5"></label>
						<input value="4" name="art" id="art-4" type="radio">
						<label for="art-4"></label>
						<input value="3" name="art" id="art-3" type="radio">
						<label for="art-3"></label>
						<input value="2" name="art" id="art-2" type="radio">
						<label for="art-2"></label>
						<input value="1" name="art" id="art-1" type="radio">
						<label for="art-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="art">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Красноречие</h4>
					<div class="content__skill-level">
						<input value="5" name="expression" id="expression-5" type="radio">
						<label for="expression-5"></label>
						<input value="4" name="expression" id="expression-4" type="radio">
						<label for="expression-4"></label>
						<input value="3" name="expression" id="expression-3" type="radio">
						<label for="expression-3"></label>
						<input value="2" name="expression" id="expression-2" type="radio">
						<label for="expression-2"></label>
						<input value="1" name="expression" id="expression-1" type="radio">
						<label for="expression-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="expression">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Лидерство</h4>
					<div class="content__skill-level">
						<input value="5" name="leadership" id="leadership-5" type="radio">
						<label for="leadership-5"></label>
						<input value="4" name="leadership" id="leadership-4" type="radio">
						<label for="leadership-4"></label>
						<input value="3" name="leadership" id="leadership-3" type="radio">
						<label for="leadership-3"></label>
						<input value="2" name="leadership" id="leadership-2" type="radio">
						<label for="leadership-2"></label>
						<input value="1" name="leadership" id="leadership-1" type="radio">
						<label for="leadership-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="leadership">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Уличные порядки</h4>
					<div class="content__skill-level">
						<input value="5" name="streetwise" id="streetwise-5" type="radio">
						<label for="streetwise-5"></label>
						<input value="4" name="streetwise" id="streetwise-4" type="radio">
						<label for="streetwise-4"></label>
						<input value="3" name="streetwise" id="streetwise-3" type="radio">
						<label for="streetwise-3"></label>
						<input value="2" name="streetwise" id="streetwise-2" type="radio">
						<label for="streetwise-2"></label>
						<input value="1" name="streetwise" id="streetwise-1" type="radio">
						<label for="streetwise-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="streetwise">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Хитрость</h4>
					<div class="content__skill-level">
						<input value="5" name="subterfuge" id="subterfuge-5" type="radio">
						<label for="subterfuge-5"></label>
						<input value="4" name="subterfuge" id="subterfuge-4" type="radio">
						<label for="subterfuge-4"></label>
						<input value="3" name="subterfuge" id="subterfuge-3" type="radio">
						<label for="subterfuge-3"></label>
						<input value="2" name="subterfuge" id="subterfuge-2" type="radio">
						<label for="subterfuge-2"></label>
						<input value="1" name="subterfuge" id="subterfuge-1" type="radio">
						<label for="subterfuge-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="subterfuge">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Экстрасенсорика</h4>
					<div class="content__skill-level">
						<input value="5" name="awareness" id="awareness-5" type="radio">
						<label for="awareness-5"></label>
						<input value="4" name="awareness" id="awareness-4" type="radio">
						<label for="awareness-4"></label>
						<input value="3" name="awareness" id="awareness-3" type="radio">
						<label for="awareness-3"></label>
						<input value="2" name="awareness" id="awareness-2" type="radio">
						<label for="awareness-2"></label>
						<input value="1" name="awareness" id="awareness-1" type="radio">
						<label for="awareness-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="awareness">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Эмпатия</h4>
					<div class="content__skill-level">
						<input value="5" name="empathy" id="empathy-5" type="radio">
						<label for="empathy-5"></label>
						<input value="4" name="empathy" id="empathy-4" type="radio">
						<label for="empathy-4"></label>
						<input value="3" name="empathy" id="empathy-3" type="radio">
						<label for="empathy-3"></label>
						<input value="2" name="empathy" id="empathy-2" type="radio">
						<label for="empathy-2"></label>
						<input value="1" name="empathy" id="empathy-1" type="radio">
						<label for="empathy-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="empathy">

						<CrossSVG></CrossSVG>
					</button>
				</div>
			</div>
			<div class="content__item-group">
				<h3>Навыки</h3>
				<div class="content__item-group-inner">
					<h4>Боевые искусства</h4>
					<div class="content__skill-level">
						<input value="5" name="martialarts" id="martialarts-5" type="radio">
						<label for="martialarts-5"></label>
						<input value="4" name="martialarts" id="martialarts-4" type="radio">
						<label for="martialarts-4"></label>
						<input value="3" name="martialarts" id="martialarts-3" type="radio">
						<label for="martialarts-3"></label>
						<input value="2" name="martialarts" id="martialarts-2" type="radio">
						<label for="martialarts-2"></label>
						<input value="1" name="martialarts" id="martialarts-1" type="radio">
						<label for="martialarts-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="martialarts">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Вождение</h4>
					<div class="content__skill-level">
						<input value="5" name="drive" id="drive-5" type="radio">
						<label for="drive-5"></label>
						<input value="4" name="drive" id="drive-4" type="radio">
						<label for="drive-4"></label>
						<input value="3" name="drive" id="drive-3" type="radio">
						<label for="drive-3"></label>
						<input value="2" name="drive" id="drive-2" type="radio">
						<label for="drive-2"></label>
						<input value="1" name="drive" id="drive-1" type="radio">
						<label for="drive-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="drive">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Выживание</h4>
					<div class="content__skill-level">
						<input value="5" name="survival" id="survival-5" type="radio">
						<label for="survival-5"></label>
						<input value="4" name="survival" id="survival-4" type="radio">
						<label for="survival-4"></label>
						<input value="3" name="survival" id="survival-3" type="radio">
						<label for="survival-3"></label>
						<input value="2" name="survival" id="survival-2" type="radio">
						<label for="survival-2"></label>
						<input value="1" name="survival" id="survival-1" type="radio">
						<label for="survival-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="survival">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Исследование</h4>
					<div class="content__skill-level">
						<input value="5" name="research" id="research-5" type="radio">
						<label for="research-5"></label>
						<input value="4" name="research" id="research-4" type="radio">
						<label for="research-4"></label>
						<input value="3" name="research" id="research-3" type="radio">
						<label for="research-3"></label>
						<input value="2" name="research" id="research-2" type="radio">
						<label for="research-2"></label>
						<input value="1" name="research" id="research-1" type="radio">
						<label for="research-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="research">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Медитация</h4>
					<div class="content__skill-level">
						<input value="5" name="meditation" id="meditation-5" type="radio">
						<label for="meditation-5"></label>
						<input value="4" name="meditation" id="meditation-4" type="radio">
						<label for="meditation-4"></label>
						<input value="3" name="meditation" id="meditation-3" type="radio">
						<label for="meditation-3"></label>
						<input value="2" name="meditation" id="meditation-2" type="radio">
						<label for="meditation-2"></label>
						<input value="1" name="meditation" id="meditation-1" type="radio">
						<label for="meditation-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="meditation">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Огнестрельное оружие</h4>
					<div class="content__skill-level">
						<input value="5" name="firearms" id="firearms-5" type="radio">
						<label for="firearms-5"></label>
						<input value="4" name="firearms" id="firearms-4" type="radio">
						<label for="firearms-4"></label>
						<input value="3" name="firearms" id="firearms-3" type="radio">
						<label for="firearms-3"></label>
						<input value="2" name="firearms" id="firearms-2" type="radio">
						<label for="firearms-2"></label>
						<input value="1" name="firearms" id="firearms-1" type="radio">
						<label for="firearms-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="firearms">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Ремёсла</h4>
					<div class="content__skill-level">
						<input value="5" name="crafts" id="crafts-5" type="radio">
						<label for="crafts-5"></label>
						<input value="4" name="crafts" id="crafts-4" type="radio">
						<label for="crafts-4"></label>
						<input value="3" name="crafts" id="crafts-3" type="radio">
						<label for="crafts-3"></label>
						<input value="2" name="crafts" id="crafts-2" type="radio">
						<label for="crafts-2"></label>
						<input value="1" name="crafts" id="crafts-1" type="radio">
						<label for="crafts-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="crafts-2">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Скрытность</h4>
					<div class="content__skill-level">
						<input value="5" name="stealth" id="stealth-5" type="radio">
						<label for="stealth-5"></label>
						<input value="4" name="stealth" id="stealth-4" type="radio">
						<label for="stealth-4"></label>
						<input value="3" name="stealth" id="stealth-3" type="radio">
						<label for="stealth-3"></label>
						<input value="2" name="stealth" id="stealth-2" type="radio">
						<label for="stealth-2"></label>
						<input value="1" name="stealth" id="stealth-1" type="radio">
						<label for="stealth-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="stealth">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Технологии</h4>
					<div class="content__skill-level">
						<input value="5" name="technology" id="technology-5" type="radio">
						<label for="technology-5"></label>
						<input value="4" name="technology" id="technology-4" type="radio">
						<label for="technology-4"></label>
						<input value="3" name="technology" id="technology-3" type="radio">
						<label for="technology-3"></label>
						<input value="2" name="technology" id="technology-2" type="radio">
						<label for="technology-2"></label>
						<input value="1" name="technology" id="technology-1" type="radio">
						<label for="technology-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="technology">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Холодное оружие</h4>
					<div class="content__skill-level">
						<input value="5" name="melee" id="melee-5" type="radio">
						<label for="melee-5"></label>
						<input value="4" name="melee" id="melee-4" type="radio">
						<label for="melee-4"></label>
						<input value="3" name="melee" id="melee-3" type="radio">
						<label for="melee-3"></label>
						<input value="2" name="melee" id="melee-2" type="radio">
						<label for="melee-2"></label>
						<input value="1" name="melee" id="melee-1" type="radio">
						<label for="melee-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="melee">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Этикет</h4>
					<div class="content__skill-level">
						<input value="5" name="etiquette" id="etiquette-5" type="radio">
						<label for="etiquette-5"></label>
						<input value="4" name="etiquette" id="etiquette-4" type="radio">
						<label for="etiquette-4"></label>
						<input value="3" name="etiquette" id="etiquette-3" type="radio">
						<label for="etiquette-3"></label>
						<input value="2" name="etiquette" id="etiquette-2" type="radio">
						<label for="etiquette-2"></label>
						<input value="1" name="etiquette" id="etiquette-1" type="radio">
						<label for="etiquette-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="etiquette">

						<CrossSVG></CrossSVG>
					</button>
				</div>
			</div>
			<div class="content__item-group">
				<h3>Познания</h3>
				<div class="content__item-group-inner">
					<h4>Гуманитарные науки</h4>
					<div class="content__skill-level">
						<input value="5" name="academics" id="academics-5" type="radio">
						<label for="academics-5"></label>
						<input value="4" name="academics" id="academics-4" type="radio">
						<label for="academics-4"></label>
						<input value="3" name="academics" id="academics-3" type="radio">
						<label for="academics-3"></label>
						<input value="2" name="academics" id="academics-2" type="radio">
						<label for="academics-2"></label>
						<input value="1" name="academics" id="academics-1" type="radio">
						<label for="academics-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton">
						<!-- class='academics' -->
						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Естественные науки</h4>
					<div class="content__skill-level">
						<input value="5" name="science" id="science-5" type="radio">
						<label for="science-5"></label>
						<input value="4" name="science" id="science-4" type="radio">
						<label for="science-4"></label>
						<input value="3" name="science" id="science-3" type="radio">
						<label for="science-3"></label>
						<input value="2" name="science" id="science-2" type="radio">
						<label for="science-2"></label>
						<input value="1" name="science" id="science-1" type="radio">
						<label for="science-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="science">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Загадки</h4>
					<div class="content__skill-level">
						<input value="5" name="enigmas" id="enigmas-5" type="radio">
						<label for="enigmas-5"></label>
						<input value="4" name="enigmas" id="enigmas-4" type="radio">
						<label for="enigmas-4"></label>
						<input value="3" name="enigmas" id="enigmas-3" type="radio">
						<label for="enigmas-3"></label>
						<input value="2" name="enigmas" id="enigmas-2" type="radio">
						<label for="enigmas-2"></label>
						<input value="1" name="enigmas" id="enigmas-1" type="radio">
						<label for="enigmas-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="enigmas">
						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Законы</h4>
					<div class="content__skill-level">
						<input value="5" name="law" id="law-5" type="radio">
						<label for="law-5"></label>
						<input value="4" name="law" id="law-4" type="radio">
						<label for="law-4"></label>
						<input value="3" name="law" id="law-3" type="radio">
						<label for="law-3"></label>
						<input value="2" name="law" id="law-2" type="radio">
						<label for="law-2"></label>
						<input value="1" name="law" id="law-1" type="radio">
						<label for="law-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="law">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Компьютеры</h4>
					<div class="content__skill-level">
						<input value="5" name="computer" id="computer-5" type="radio">
						<label for="computer-5"></label>
						<input value="4" name="computer" id="computer-4" type="radio">
						<label for="computer-4"></label>
						<input value="3" name="computer" id="computer-3" type="radio">
						<label for="computer-3"></label>
						<input value="2" name="computer" id="computer-2" type="radio">
						<label for="computer-2"></label>
						<input value="1" name="computer" id="computer-1" type="radio">
						<label for="computer-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="computer">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Космология</h4>
					<div class="content__skill-level">
						<input value="5" name="cosmology" id="cosmology-5" type="radio">
						<label for="cosmology-5"></label>
						<input value="4" name="cosmology" id="cosmology-4" type="radio">
						<label for="cosmology-4"></label>
						<input value="3" name="cosmology" id="cosmology-3" type="radio">
						<label for="cosmology-3"></label>
						<input value="2" name="cosmology" id="cosmology-2" type="radio">
						<label for="cosmology-2"></label>
						<input value="1" name="cosmology" id="cosmology-1" type="radio">
						<label for="cosmology-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="cosmology">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Медицина</h4>
					<div class="content__skill-level">
						<input value="5" name="medicine" id="medicine-5" type="radio">
						<label for="medicine-5"></label>
						<input value="4" name="medicine" id="medicine-4" type="radio">
						<label for="medicine-4"></label>
						<input value="3" name="medicine" id="medicine-3" type="radio">
						<label for="medicine-3"></label>
						<input value="2" name="medicine" id="medicine-2" type="radio">
						<label for="medicine-2"></label>
						<input value="1" name="medicine" id="medicine-1" type="radio">
						<label for="medicine-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="medicine">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Оккультизм</h4>
					<div class="content__skill-level">
						<input value="5" name="occult" id="occult-5" type="radio">
						<label for="occult-5"></label>
						<input value="4" name="occult" id="occult-4" type="radio">
						<label for="occult-4"></label>
						<input value="3" name="occult" id="occult-3" type="radio">
						<label for="occult-3"></label>
						<input value="2" name="occult" id="occult-2" type="radio">
						<label for="occult-2"></label>
						<input value="1" name="occult" id="occult-1" type="radio">
						<label for="occult-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="occult">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Политика</h4>
					<div class="content__skill-level">
						<input value="5" name="politics" id="politics-5" type="radio">
						<label for="politics-5"></label>
						<input value="4" name="politics" id="politics-4" type="radio">
						<label for="politics-4"></label>
						<input value="3" name="politics" id="politics-3" type="radio">
						<label for="politics-3"></label>
						<input value="2" name="politics" id="politics-2" type="radio">
						<label for="politics-2"></label>
						<input value="1" name="politics" id="politics-1" type="radio">
						<label for="politics-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="politics">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Расследование</h4>
					<div class="content__skill-level">
						<input value="5" name="investigation" id="investigation-5" type="radio">
						<label for="investigation-5"></label>
						<input value="4" name="investigation" id="investigation-4" type="radio">
						<label for="investigation-4"></label>
						<input value="3" name="investigation" id="investigation-3" type="radio">
						<label for="investigation-3"></label>
						<input value="2" name="investigation" id="investigation-2" type="radio">
						<label for="investigation-2"></label>
						<input value="1" name="investigation" id="investigation-1" type="radio">
						<label for="investigation-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="investigation">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Эзотерика</h4>
					<div class="content__skill-level">
						<input value="5" name="esoterica" id="esoterica-5" type="radio">
						<label for="esoterica-5"></label>
						<input value="4" name="esoterica" id="esoterica-4" type="radio">
						<label for="esoterica-4"></label>
						<input value="3" name="esoterica" id="esoterica-3" type="radio">
						<label for="esoterica-3"></label>
						<input value="2" name="esoterica" id="esoterica-2" type="radio">
						<label for="esoterica-2"></label>
						<input value="1" name="esoterica" id="esoterica-1" type="radio">
						<label for="esoterica-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="esoterica">

						<CrossSVG></CrossSVG>
					</button>
				</div>
			</div>
		</div>
		<h2>Сферы</h2>
		<div class="content__section">
			<div class="content__item-group">
				<div class="content__item-group-inner">
					<h4>Жизнь</h4>
					<div class="content__skill-level">
						<input value="5" name="life" id="life-5" type="radio">
						<label for="life-5"></label>
						<input value="4" name="life" id="life-4" type="radio">
						<label for="life-4"></label>
						<input value="3" name="life" id="life-3" type="radio">
						<label for="life-3"></label>
						<input value="2" name="life" id="life-2" type="radio">
						<label for="life-2"></label>
						<input value="1" name="life" id="life-1" type="radio">
						<label for="life-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="life">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Материя</h4>
					<div class="content__skill-level">
						<input value="5" name="matter" id="matter-5" type="radio">
						<label for="matter-5"></label>
						<input value="4" name="matter" id="matter-4" type="radio">
						<label for="matter-4"></label>
						<input value="3" name="matter" id="matter-3" type="radio">
						<label for="matter-3"></label>
						<input value="2" name="matter" id="matter-2" type="radio">
						<label for="matter-2"></label>
						<input value="1" name="matter" id="matter-1" type="radio">
						<label for="matter-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="matter">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Силы</h4>
					<div class="content__skill-level">
						<input value="5" name="forces" id="forces-5" type="radio">
						<label for="forces-5"></label>
						<input value="4" name="forces" id="forces-4" type="radio">
						<label for="forces-4"></label>
						<input value="3" name="forces" id="forces-3" type="radio">
						<label for="forces-3"></label>
						<input value="2" name="forces" id="forces-2" type="radio">
						<label for="forces-2"></label>
						<input value="1" name="forces" id="forces-1" type="radio">
						<label for="forces-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="forces">

						<CrossSVG></CrossSVG>
					</button>
				</div>
			</div>
			<div class="content__item-group">
				<div class="content__item-group-inner">
					<h4>Время</h4>
					<div class="content__skill-level">
						<input value="5" name="time" id="time-5" type="radio">
						<label for="time-5"></label>
						<input value="4" name="time" id="time-4" type="radio">
						<label for="time-4"></label>
						<input value="3" name="time" id="time-3" type="radio">
						<label for="time-3"></label>
						<input value="2" name="time" id="time-2" type="radio">
						<label for="time-2"></label>
						<input value="1" name="time" id="time-1" type="radio">
						<label for="time-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="time">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Связи</h4>
					<div class="content__skill-level">
						<input value="5" name="correspondence" id="correspondence-5" type="radio">
						<label for="correspondence-5"></label>
						<input value="4" name="correspondence" id="correspondence-4" type="radio">
						<label for="correspondence-4"></label>
						<input value="3" name="correspondence" id="correspondence-3" type="radio">
						<label for="correspondence-3"></label>
						<input value="2" name="correspondence" id="correspondence-2" type="radio">
						<label for="correspondence-2"></label>
						<input value="1" name="correspondence" id="correspondence-1" type="radio">
						<label for="correspondence-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="correspondence">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Энтропия</h4>
					<div class="content__skill-level">
						<input value="5" name="entropy" id="entropy-5" type="radio">
						<label for="entropy-5"></label>
						<input value="4" name="entropy" id="entropy-4" type="radio">
						<label for="entropy-4"></label>
						<input value="3" name="entropy" id="entropy-3" type="radio">
						<label for="entropy-3"></label>
						<input value="2" name="entropy" id="entropy-2" type="radio">
						<label for="entropy-2"></label>
						<input value="1" name="entropy" id="entropy-1" type="radio">
						<label for="entropy-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="entropy">

						<CrossSVG></CrossSVG>
					</button>
				</div>
			</div>
			<div class="content__item-group">
				<div class="content__item-group-inner">
					<h4>Дух</h4>
					<div class="content__skill-level">
						<input value="5" name="spirit" id="spirit-5" type="radio">
						<label for="spirit-5"></label>
						<input value="4" name="spirit" id="spirit-4" type="radio">
						<label for="spirit-4"></label>
						<input value="3" name="spirit" id="spirit-3" type="radio">
						<label for="spirit-3"></label>
						<input value="2" name="spirit" id="spirit-2" type="radio">
						<label for="spirit-2"></label>
						<input value="1" name="spirit" id="spirit-1" type="radio">
						<label for="spirit-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="spirit">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Разум</h4>
					<div class="content__skill-level">
						<input value="5" name="mind" id="mind-5" type="radio">
						<label for="mind-5"></label>
						<input value="4" name="mind" id="mind-4" type="radio">
						<label for="mind-4"></label>
						<input value="3" name="mind" id="mind-3" type="radio">
						<label for="mind-3"></label>
						<input value="2" name="mind" id="mind-2" type="radio">
						<label for="mind-2"></label>
						<input value="1" name="mind" id="mind-1" type="radio">
						<label for="mind-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="mind">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner">
					<h4>Основы</h4>
					<div class="content__skill-level">
						<input value="5" name="prime" id="prime-5" type="radio">
						<label for="prime-5"></label>
						<input value="4" name="prime" id="prime-4" type="radio">
						<label for="prime-4"></label>
						<input value="3" name="prime" id="prime-3" type="radio">
						<label for="prime-3"></label>
						<input value="2" name="prime" id="prime-2" type="radio">
						<label for="prime-2"></label>
						<input value="1" name="prime" id="prime-1" type="radio">
						<label for="prime-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="prime">

						<CrossSVG></CrossSVG>
					</button>
				</div>
			</div>
		</div>
		<h2>Преимущества</h2>
		<div class="content__section content__section--ai-start">
			<div class="content__item-group content__item-group--gap4">
				<h3>Дополнения</h3>
				<div
						 class="content__item-group-inner content__item-group-inner--input-ht20 content__item-group-inner--gap30">
					<input type="text" name="backgrounds-onei">
					<div class="content__skill-level">
						<input value="5" name="backgrounds-one" id="backgrounds-one-5" type="radio">
						<label for="backgrounds-one-5"></label>
						<input value="4" name="backgrounds-one" id="backgrounds-one-4" type="radio">
						<label for="backgrounds-one-4"></label>
						<input value="3" name="backgrounds-one" id="backgrounds-one-3" type="radio">
						<label for="backgrounds-one-3"></label>
						<input value="2" name="backgrounds-one" id="backgrounds-one-2" type="radio">
						<label for="backgrounds-one-2"></label>
						<input value="1" name="backgrounds-one" id="backgrounds-one-1" type="radio">
						<label for="backgrounds-one-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="backgrounds-one">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div
						 class="content__item-group-inner content__item-group-inner--input-ht20 content__item-group-inner--gap30">
					<input type="text" name="backgrounds-twoi">
					<div class="content__skill-level">
						<input value="5" name="backgrounds-two" id="backgrounds-two-5" type="radio">
						<label for="backgrounds-two-5"></label>
						<input value="4" name="backgrounds-two" id="backgrounds-two-4" type="radio">
						<label for="backgrounds-two-4"></label>
						<input value="3" name="backgrounds-two" id="backgrounds-two-3" type="radio">
						<label for="backgrounds-two-3"></label>
						<input value="2" name="backgrounds-two" id="backgrounds-two-2" type="radio">
						<label for="backgrounds-two-2"></label>
						<input value="1" name="backgrounds-two" id="backgrounds-two-1" type="radio">
						<label for="backgrounds-two-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="backgrounds-two">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div
						 class="content__item-group-inner content__item-group-inner--input-ht20 content__item-group-inner--gap30">
					<input type="text" name="backgrounds-threei">
					<div class="content__skill-level">
						<input value="5" name="backgrounds-three" id="backgrounds-three-5" type="radio">
						<label for="backgrounds-three-5"></label>
						<input value="4" name="backgrounds-three" id="backgrounds-three-4" type="radio">
						<label for="backgrounds-three-4"></label>
						<input value="3" name="backgrounds-three" id="backgrounds-three-3" type="radio">
						<label for="backgrounds-three-3"></label>
						<input value="2" name="backgrounds-three" id="backgrounds-three-2" type="radio">
						<label for="backgrounds-three-2"></label>
						<input value="1" name="backgrounds-three" id="backgrounds-three-1" type="radio">
						<label for="backgrounds-three-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="backgrounds-three">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div
						 class="content__item-group-inner content__item-group-inner--input-ht20 content__item-group-inner--gap30">
					<input type="text" name="backgrounds-fouri">
					<div class="content__skill-level">
						<input value="5" name="backgrounds-four" id="backgrounds-four-5" type="radio">
						<label for="backgrounds-four-5"></label>
						<input value="4" name="backgrounds-four" id="backgrounds-four-4" type="radio">
						<label for="backgrounds-four-4"></label>
						<input value="3" name="backgrounds-four" id="backgrounds-four-3" type="radio">
						<label for="backgrounds-four-3"></label>
						<input value="2" name="backgrounds-four" id="backgrounds-four-2" type="radio">
						<label for="backgrounds-four-2"></label>
						<input value="1" name="backgrounds-four" id="backgrounds-four-1" type="radio">
						<label for="backgrounds-four-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="backgrounds-four">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div
						 class="content__item-group-inner content__item-group-inner--input-ht20 content__item-group-inner--gap30">
					<input type="text" name="backgrounds-fivei">
					<div class="content__skill-level">
						<input value="5" name="backgrounds-five" id="backgrounds-five-5" type="radio">
						<label for="backgrounds-five-5"></label>
						<input value="4" name="backgrounds-five" id="backgrounds-five-4" type="radio">
						<label for="backgrounds-five-4"></label>
						<input value="3" name="backgrounds-five" id="backgrounds-five-3" type="radio">
						<label for="backgrounds-five-3"></label>
						<input value="2" name="backgrounds-five" id="backgrounds-five-2" type="radio">
						<label for="backgrounds-five-2"></label>
						<input value="1" name="backgrounds-five" id="backgrounds-five-1" type="radio">
						<label for="backgrounds-five-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="backgrounds-five">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div
						 class="content__item-group-inner content__item-group-inner--input-ht20 content__item-group-inner--gap30 content__item-group-inner--mb20">
					<input type="text" name="backgrounds-sixi">
					<div class="content__skill-level">
						<input value="5" name="backgrounds-six" id="backgrounds-six-5" type="radio">
						<label for="backgrounds-six-5"></label>
						<input value="4" name="backgrounds-six" id="backgrounds-six-4" type="radio">
						<label for="backgrounds-six-4"></label>
						<input value="3" name="backgrounds-six" id="backgrounds-six-3" type="radio">
						<label for="backgrounds-six-3"></label>
						<input value="2" name="backgrounds-six" id="backgrounds-six-2" type="radio">
						<label for="backgrounds-six-2"></label>
						<input value="1" name="backgrounds-six" id="backgrounds-six-1" type="radio">
						<label for="backgrounds-six-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="backgrounds-six">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<h3>Прочие особенности</h3>
				<div
						 class="content__item-group-inner content__item-group-inner--input-ht20 content__item-group-inner--gap30">
					<input type="text" name="other-traits-onei">
					<div class="content__skill-level">
						<input value="5" name="other-traits-one" id="other-traits-one-5" type="radio">
						<label for="other-traits-one-5"></label>
						<input value="4" name="other-traits-one" id="other-traits-one-4" type="radio">
						<label for="other-traits-one-4"></label>
						<input value="3" name="other-traits-one" id="other-traits-one-3" type="radio">
						<label for="other-traits-one-3"></label>
						<input value="2" name="other-traits-one" id="other-traits-one-2" type="radio">
						<label for="other-traits-one-2"></label>
						<input value="1" name="other-traits-one" id="other-traits-one-1" type="radio">
						<label for="other-traits-one-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="other-traits-one">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div
						 class="content__item-group-inner content__item-group-inner--input-ht20 content__item-group-inner--gap30">
					<input type="text" name="other-traits-twoi">
					<div class="content__skill-level">
						<input value="5" name="other-traits-two" id="other-traits-two-5" type="radio">
						<label for="other-traits-two-5"></label>
						<input value="4" name="other-traits-two" id="other-traits-two-4" type="radio">
						<label for="other-traits-two-4"></label>
						<input value="3" name="other-traits-two" id="other-traits-two-3" type="radio">
						<label for="other-traits-two-3"></label>
						<input value="2" name="other-traits-two" id="other-traits-two-2" type="radio">
						<label for="other-traits-two-2"></label>
						<input value="1" name="other-traits-two" id="other-traits-two-1" type="radio">
						<label for="other-traits-two-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="other-traits-two">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div
						 class="content__item-group-inner content__item-group-inner--input-ht20 content__item-group-inner--gap30">
					<input type="text" name="other-traits-threei">
					<div class="content__skill-level">
						<input value="5" name="other-traits-three" id="other-traits-three-5" type="radio">
						<label for="other-traits-three-5"></label>
						<input value="4" name="other-traits-three" id="other-traits-three-4" type="radio">
						<label for="other-traits-three-4"></label>
						<input value="3" name="other-traits-three" id="other-traits-three-3" type="radio">
						<label for="other-traits-three-3"></label>
						<input value="2" name="other-traits-three" id="other-traits-three-2" type="radio">
						<label for="other-traits-three-2"></label>
						<input value="1" name="other-traits-three" id="other-traits-three-1" type="radio">
						<label for="other-traits-three-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="other-traits-three">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div
						 class="content__item-group-inner content__item-group-inner--input-ht20 content__item-group-inner--gap30">
					<input type="text" name="other-traits-fouri">
					<div class="content__skill-level">
						<input value="5" name="other-traits-four" id="other-traits-four-5" type="radio">
						<label for="other-traits-four-5"></label>
						<input value="4" name="other-traits-four" id="other-traits-four-4" type="radio">
						<label for="other-traits-four-4"></label>
						<input value="3" name="other-traits-four" id="other-traits-four-3" type="radio">
						<label for="other-traits-four-3"></label>
						<input value="2" name="other-traits-four" id="other-traits-four-2" type="radio">
						<label for="other-traits-four-2"></label>
						<input value="1" name="other-traits-four" id="other-traits-four-1" type="radio">
						<label for="other-traits-four-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="other-traits-four">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div
						 class="content__item-group-inner content__item-group-inner--input-ht20 content__item-group-inner--gap30">
					<input type="text" name="other-traits-fivei">
					<div class="content__skill-level">
						<input value="5" name="other-traits-five" id="other-traits-five-5" type="radio">
						<label for="other-traits-five-5"></label>
						<input value="4" name="other-traits-five" id="other-traits-five-4" type="radio">
						<label for="other-traits-five-4"></label>
						<input value="3" name="other-traits-five" id="other-traits-five-3" type="radio">
						<label for="other-traits-five-3"></label>
						<input value="2" name="other-traits-five" id="other-traits-five-2" type="radio">
						<label for="other-traits-five-2"></label>
						<input value="1" name="other-traits-five" id="other-traits-five-1" type="radio">
						<label for="other-traits-five-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="other-traits-five">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div
						 class="content__item-group-inner content__item-group-inner--input-ht20 content__item-group-inner--gap30">
					<input type="text" name="other-traits-sixi">
					<div class="content__skill-level">
						<input value="5" name="other-traits-six" id="other-traits-six-5" type="radio">
						<label for="other-traits-six-5"></label>
						<input value="4" name="other-traits-six" id="other-traits-six-4" type="radio">
						<label for="other-traits-six-4"></label>
						<input value="3" name="other-traits-six" id="other-traits-six-3" type="radio">
						<label for="other-traits-six-3"></label>
						<input value="2" name="other-traits-six" id="other-traits-six-2" type="radio">
						<label for="other-traits-six-2"></label>
						<input value="1" name="other-traits-six" id="other-traits-six-1" type="radio">
						<label for="other-traits-six-1"></label>
					</div>
					<button class='content_skill-clear-button' id="clearButton" data-group="other-traits-six">

						<CrossSVG></CrossSVG>
					</button>
				</div>
			</div>
			<div class="content__item-group content__item-group--gap0">
				<h3>Арете</h3>
				<div class="content__item-group-inner content__item-group-inner--fd-col content__item-group-inner--mb20">
					<div class="content__skill-level">
						<input value="10" name="arete" id="arete-10" type="radio">
						<label for="arete-10"></label>
						<input value="9" name="arete" id="arete-9" type="radio">
						<label for="arete-9"></label>
						<input value="8" name="arete" id="arete-8" type="radio">
						<label for="arete-8"></label>
						<input value="7" name="arete" id="arete-7" type="radio">
						<label for="arete-7"></label>
						<input value="6" name="arete" id="arete-6" type="radio">
						<label for="arete-6"></label>
						<input value="5" name="arete" id="arete-5" type="radio">
						<label for="arete-5"></label>
						<input value="4" name="arete" id="arete-4" type="radio">
						<label for="arete-4"></label>
						<input value="3" name="arete" id="arete-3" type="radio">
						<label for="arete-3"></label>
						<input value="2" name="arete" id="arete-2" type="radio">
						<label for="arete-2"></label>
						<input value="1" name="arete" id="arete-1" type="radio">
						<label for="arete-1"></label>
					</div>
					<button class='content_skill-clear-button content_skill-clear-button--left20' id="clearButton"
									data-group="arete">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<h3>Сила воли</h3>
				<div class="content__item-group-inner content__item-group-inner--fd-col content__item-group-inner--mb20">
					<div class="content__skill-level">
						<input value="10" name="willpower" id="willpower-10" type="radio">
						<label for="willpower-10"></label>
						<input value="9" name="willpower" id="willpower-9" type="radio">
						<label for="willpower-9"></label>
						<input value="8" name="willpower" id="willpower-8" type="radio">
						<label for="willpower-8"></label>
						<input value="7" name="willpower" id="willpower-7" type="radio">
						<label for="willpower-7"></label>
						<input value="6" name="willpower" id="willpower-6" type="radio">
						<label for="willpower-6"></label>
						<input value="5" name="willpower" id="willpower-5" type="radio">
						<label for="willpower-5"></label>
						<input value="4" name="willpower" id="willpower-4" type="radio">
						<label for="willpower-4"></label>
						<input value="3" name="willpower" id="willpower-3" type="radio">
						<label for="willpower-3"></label>
						<input value="2" name="willpower" id="willpower-2" type="radio">
						<label for="willpower-2"></label>
						<input value="1" name="willpower" id="willpower-1" type="radio">
						<label for="willpower-1"></label>
					</div>
					<div class="content__skill-level content__skill-level--content-X">
						<input value="10" name="willpower-2" id="willpower-2-10" type="radio">
						<label for="willpower-2-10"></label>
						<input value="9" name="willpower-2" id="willpower-2-9" type="radio">
						<label for="willpower-2-9"></label>
						<input value="8" name="willpower-2" id="willpower-2-8" type="radio">
						<label for="willpower-2-8"></label>
						<input value="7" name="willpower-2" id="willpower-2-7" type="radio">
						<label for="willpower-2-7"></label>
						<input value="6" name="willpower-2" id="willpower-2-6" type="radio">
						<label for="willpower-2-6"></label>
						<input value="5" name="willpower-2" id="willpower-2-5" type="radio">
						<label for="willpower-2-5"></label>
						<input value="4" name="willpower-2" id="willpower-2-4" type="radio">
						<label for="willpower-2-4"></label>
						<input value="3" name="willpower-2" id="willpower-2-3" type="radio">
						<label for="willpower-2-3"></label>
						<input value="2" name="willpower-2" id="willpower-2-2" type="radio">
						<label for="willpower-2-2"></label>
						<input value="1" name="willpower-2" id="willpower-2-1" type="radio">
						<label for="willpower-2-1"></label>
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
					<button class='content_skill-clear-button content_skill-clear-button--left20' id="clearButton"
									data-group="willpower,willpower-2,willpower-2-alt">

						<CrossSVG></CrossSVG>
					</button>
				</div>
				<div class="content__item-group-inner content__item-group-inner--fd-col">
					<h3>Квинтэссенция</h3>
					<div class="circle-container">
						<div class="dot"></div>
						<input type="checkbox" name='circle-xo' id="check1"><label for="check1" class="pos-1 o"></label>
						<input type="checkbox" name='circle-xo' id="check2"><label for="check2" class="pos-2 o"></label>
						<input type="checkbox" name='circle-xo' id="check3"><label for="check3" class="pos-3 o"></label>
						<input type="checkbox" name='circle-xo' id="check4"><label for="check4" class="pos-4 o"></label>
						<input type="checkbox" name='circle-xo' id="check5"><label for="check5" class="pos-5 o"></label>
						<input type="checkbox" name='circle-xo' id="check6"><label for="check6" class="pos-6 x"></label>
						<input type="checkbox" name='circle-xo' id="check7"><label for="check7" class="pos-7 x"></label>
						<input type="checkbox" name='circle-xo' id="check8"><label for="check8" class="pos-8 x"></label>
						<input type="checkbox" name='circle-xo' id="check9"><label for="check9" class="pos-9 x"></label>
						<input type="checkbox" name='circle-xo' id="check10"><label for="check10" class="pos-10 x"></label>
						<input type="checkbox" name='circle-xo' id="check11"><label for="check11" class="pos-11 x"></label>
						<input type="checkbox" name='circle-xo' id="check12"><label for="check12" class="pos-12 x"></label>
						<input type="checkbox" name='circle-xo' id="check13"><label for="check13" class="pos-13 x"></label>
						<input type="checkbox" name='circle-xo' id="check14"><label for="check14" class="pos-14 x"></label>
						<input type="checkbox" name='circle-xo' id="check15"><label for="check15" class="pos-15 x"></label>
						<input type="checkbox" name='circle-xo' id="check16"><label for="check16" class="pos-16 o"></label>
						<input type="checkbox" name='circle-xo' id="check17"><label for="check17" class="pos-17 o"></label>
						<input type="checkbox" name='circle-xo' id="check18"><label for="check18" class="pos-18 o"></label>
						<input type="checkbox" name='circle-xo' id="check19"><label for="check19" class="pos-19 o"></label>
						<input type="checkbox" name='circle-xo' id="check20"><label for="check20" class="pos-20 o"></label>
					</div>
					<button class='content_skill-clear-button content_skill-clear-button--center-circle' id="clearButton"
									data-group="circle-xo">

						<CrossSVG></CrossSVG>
					</button>
					<h3>Парадокс</h3>
				</div>
			</div>
			<div class="content__item-group">
				<h3>Здоровье</h3>
				<div class="content__item-group-inner">
					<span>Пустяк</span>
					<label for="bruised" style='text-align: center;'>0</label>
					<select name="bruised" id="bruised">
						<option value="">&nbsp;</option>
						<option value="\">\</option>
						<option value="x">x</option>
						<option value="Ж">ж</option>
					</select>
				</div>
				<div class="content__item-group-inner">
					<span>Боль</span>
					<label for="hurt">-1</label>
					<select name="hurt" id="hurt">
						<option value="">&nbsp;</option>
						<option value="\">\</option>
						<option value="x">x</option>
						<option value="Ж">ж</option>
					</select>
				</div>
				<div class="content__item-group-inner">
					<span>Потрёпан</span>
					<label for="injured">-1</label>
					<select name="injured" id="injured">
						<option value="">&nbsp;</option>
						<option value="\">\</option>
						<option value="x">x</option>
						<option value="Ж">ж</option>
					</select>
				</div>
				<div class="content__item-group-inner">
					<span>Травма</span>
					<label for="wounded">-2</label>
					<select name="wounded" id="wounded">
						<option value="">&nbsp;</option>
						<option value="\">\</option>
						<option value="x">x</option>
						<option value="Ж">ж</option>
					</select>
				</div>
				<div class="content__item-group-inner">
					<span>Истерзан</span>
					<label for="mauled">-2</label>
					<select name="mauled" id="mauled">
						<option value="">&nbsp;</option>
						<option value="\">\</option>
						<option value="x">x</option>
						<option value="Ж">ж</option>
					</select>
				</div>
				<div class="content__item-group-inner">
					<span>Увечья</span>
					<label for="crippled">-5</label>
					<select name="crippled" id="crippled">
						<option value="">&nbsp;</option>
						<option value="\">\</option>
						<option value="x">x</option>
						<option value="Ж">ж</option>
					</select>
				</div>
				<div class="content__item-group-inner">
					<span>Обморок</span>
					<label for="incapactitated"></label>
					<select name="incapactitated" id="incapactitated">
						<option value="">&nbsp;</option>
						<option value="\">\</option>
						<option value="x">x</option>
						<option value="Ж">ж</option>
					</select>
				</div>
				<div class="content__item-group-inner content__item-group-inner--fd-col">
					<h3>Опыт</h3>
					<textarea name="experience" id="experience" rows="3" cols="22"></textarea>
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
	}
}
</script>
<style></style>