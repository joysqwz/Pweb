<template>
  <input
         type="text"
         :name="`${inputName}i`"
         v-model="inputValue" />
  <div class="content__skill-level">
    <RadioButton
                 v-for="value in [5, 4, 3, 2, 1]"
                 :key="value"
                 :value="value"
                 :name="inputName"
                 v-model="selectedValue" />
  </div>
  <button class='content_skill-clear-button' @click="clearSelection">
    <CrossSVG />
  </button>
</template>

<script>
export default {
  name: 'SkillInputPoint',
  props: {
    inputName: {
      type: String,
      required: true
    }
  },
  data() {
    return {
      inputValue: '',
      selectedValue: null
    }
  },
  methods: {
    clearSelection() {
      this.selectedValue = null
      this.inputValue = ''
    }
  }
}
</script>

<style></style>