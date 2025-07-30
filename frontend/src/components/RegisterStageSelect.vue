<script setup>
import { ref, onMounted, computed } from 'vue'
import Select from 'primevue/select'
import { getStages } from '../services/stage.js'

const props = defineProps({
  modelValue: {
    type: [Object, Number, String, null],
    default: null
  },
  placeHolder: {
    type: String,
    default: 'Fase de la competició'
  }
})
const emit = defineEmits(['update:modelValue'])
const stages = ref([])
const error = ref('')

async function fetchStages() {
    error.value = ''
    try {
        const result = await getStages()
        stages.value = result.map(stages=>(
      { name: stages.name, id: stages.id, value: stages.value}
    ))
    } catch (err) {
        error.value = err.message
        console.log('Error fetching stages: ', err)
    }
}

onMounted(() => fetchStages())
const inner = computed({
  get: () => props.modelValue,
  set: val => emit('update:modelValue', val)
})
</script>

<template>
    <Select
        v-model="inner"  
        :options="stages"
        optionLabel="name"
        :placeholder="stages.length === 0 ? placeHolder : (inner?.name || props.placeHolder)"
        class="w-full md:w-56"
    />
    <p v-if="error" class="text-red-500 text-xs">{{ error }}</p>
</template>