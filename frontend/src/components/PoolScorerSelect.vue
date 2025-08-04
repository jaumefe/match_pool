<script setup>
import { ref, onMounted, computed, watch } from 'vue'
import { getScorersByGroupId } from '../services/scorers.js'
import Select from 'primevue/select'

const props = defineProps({
    groupId: {
        type: Number,
        required: true
    },
    modelValue: {
        type: [Object, Number, String, null],
        default: null
    },
    placeHolder: {
        type: String,
        default: 'Selecciona un golejador'
    }
})

const emit = defineEmits(['update:modelValue'])

const scorers = ref([])
const error = ref('')

const selectedPosition = ref(null)
const selectedCountry = ref(null)

const positionLabels = {
  'forward': 'Davanter',
  'midfielder': 'Migcampista',
  'defender': 'Defensa',
  'goalkeeper': 'Porter'
}

async function fetchScorers(groupId) {
    error.value = ''
    try {
        const result = await getScorersByGroupId(groupId)
        scorers.value = result.map(scorer => ({
            name: scorer.name,
            id: scorer.id,
            position: scorer.position,
            team: scorer.team
        }))
    } catch (err) {
        error.value = err.message
        console.log('Error fetching scorers:', err)
    }
}

onMounted(() => fetchScorers(props.groupId))
watch(() => props.groupId, fetchScorers)

const uniquePositions = computed(() => {
  const positions = [...new Set(scorers.value.map(s => s.position).filter(Boolean))]
  return positions.map(pos => ({
    label: positionLabels[pos] || pos, // Usa el mapa, o el valor tal cual si no está definido
    value: pos
  }))
})

const uniqueCountries = computed(() => {
  return [...new Set(scorers.value.map(s => s.team).filter(Boolean))]
})

const filteredScorers = computed(() => {
  return scorers.value.filter(scorer => {
    const matchesPosition = selectedPosition.value?.value ? scorer.position === selectedPosition.value?.value : true
    const matchesCountry = selectedCountry.value ? scorer.team === selectedCountry.value : true
    return matchesPosition && matchesCountry
  })
})




const inner = computed({
  get: () => props.modelValue,
  set: val => emit('update:modelValue', val)
})

</script>
<template>
    <div class="flex flex-col gap-2">
    <!-- Filtro por posición -->
    <Select
      v-model="selectedPosition"
      :options="uniquePositions"
      optionLabel="label"
      placeholder="Posició"
      class="w-full md:w-56"
    />
    <!-- Filtro por país -->
    <Select
      v-model="selectedCountry"
      :options="uniqueCountries"
      placeholder="País"
      class="w-full md:w-56"
    />
    </div>

    <!-- Selector de goleador -->
     <div class="flex flex-col gap-2">
    <Select
        v-model="inner"
        :options="filteredScorers"
        optionLabel="name"
        :placeholder="filteredScorers.length === 0 ? placeHolder : (inner?.name || props.placeHolder)"
        class="w-full md:w-56"
    />
    <p v-if="error" class="text-red-500 text-xs">{{ error }}</p>
    </div>
</template>