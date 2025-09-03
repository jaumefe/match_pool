<script setup>
import { ref, computed, watch } from 'vue'
import { getScorerByTeamId } from '../services/scorers.js'
import Select from 'primevue/select'

const props = defineProps({
    teamHomeId: {
        type: Number,
        required: true
    },
    teamAwayId: {
        type: Number,
        required: true
    },
    modelValue: {
        type: [Object, Number, String, null],
        default: null
    },
    placeHolder: {
        type: String,
        default: 'Selecciona un jugador'
    }
})
const emit = defineEmits(['update:modelValue'])

const scorers = ref([])
const error = ref('')

async function fetchScorers() {
    error.value = ''
    scorers.value = []
    const teamIds = [props.teamHomeId, props.teamAwayId].filter(id=>id)

    try {
        const result = await Promise.all(
            teamIds.map(id => getScorerByTeamId(id))
        )

        const allScorers = result.flat()
        const unique = Object.values(
            allScorers.reduce((acc, scorer) => {
                acc[scorer.id] = scorer
                return acc
            }, {})
        )

        scorers.value = unique.map(scorer => ({
            name: scorer.name,
            id: scorer.id
        }))
    } catch (err) {
        error.value = err.message
        console.log("Error fetching scorers: ", err)
    }
}

watch(
    () => [props.teamHomeId, props.teamAwayId], 
    fetchScorers,
    {immediate: true})

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
        :options="scorers"
        optionLabel="name"
        :placeholder="scorers.length === 0 ? placeHolder : (inner?.name || props.placeHolder)"
        class="w-full md:w-56"
    />
</template>