<script setup>
import { ref, onMounted, computed } from 'vue'
import { getTeams } from '../services/teams.js'
import Select from 'primevue/select'

const props = defineProps({
  modelValue: {
    type: [Object, Number, String, null],
    default: null
  },
  placeHolder: {
    type: String,
    default: 'Selecciona un equip'
  }
})
const emit = defineEmits(['update:modelValue'])

const teams = ref([])
const error = ref('')

async function fetchTeams() {
  error.value = ''
  try {
    const result = await getTeams()
    teams.value = result.map(team=>(
      { name: team.name, id: team.id, value: team.value}
    )).sort((a, b) => a.name.localeCompare(b.name));
  } catch (err) {
    error.value = err.message
    console.log('Error fetching teams:', err)
  }
}

onMounted(() => fetchTeams())

const inner = computed({
  get: () => props.modelValue,
  set: val => emit('update:modelValue', val)
})
</script>

<template>
  <p v-if="error" class="error">{{ error }}</p>
    <Select
        v-else
        v-model="inner"  
        :options="teams"
        optionLabel="name"
        :placeholder="teams.length === 0 ? placeHolder : (inner?.name || props.placeHolder)"
        class="w-full md:w-56"
    />
</template>
