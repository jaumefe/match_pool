<script setup>
import { ref, onMounted, watch, computed } from 'vue'
import { getTeams } from '../services/teams.js'
import Select from 'primevue/select'

const props = defineProps({
  groupId: {
    type: Number,
    required: true
  },
  modelValue: {
    type: [Object, Number, String, null],
    default: null
  }
})
const emit = defineEmits(['update:modelValue'])

const teams = ref([])
const error = ref('')

async function fetchTeams(groupId) {
  error.value = ''
  try {
    const result = await getTeams(groupId)
    teams.value = result.map(team=>(
      { name: team.name, id: team.id, value: team.value}
    ))
  } catch (err) {
    error.value = err.message
    console.log('Error fetching teams:', err)
  }
}

onMounted(() => fetchTeams(props.groupId))
watch(() => props.groupId, fetchTeams)

const inner = computed({
  get: () => props.modelValue,
  set: val => emit('update:modelValue', val)
})
</script>

<template>
    <Select
        v-model="inner"  
        :options="teams"
        optionsLabel="name"
        placeholder="Selecciona un equipo"
        class="w-full md:w-14rem"
    />
    <p v-if="error" class="text-red-500 text-xs">{{ error }}</p>
</template>

