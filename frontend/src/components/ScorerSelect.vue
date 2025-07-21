<script setup>
import { ref, onMounted, computed, watch } from 'vue'
import { getScorers } from '../services/scorers.js'
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

async function fetchScorers(groupId) {
    error.value = ''
    try {
        const result = await getScorers(groupId)
        scorers.value = result.map(scorer => ({
            name: scorer.name,
            id: scorer.id,
            position: scorer.position,
            team_id: scorer.team_id
        }))
    } catch (err) {
        error.value = err.message
        console.log('Error fetching scorers:', err)
    }
}

onMounted(() => fetchScorers(props.groupId))
watch(() => props.groupId, fetchScorers)

const inner = computed({
  get: () => props.modelValue,
  set: val => emit('update:modelValue', val)
})
</script>
<template>
    <Select
        v-model="inner"
        :options="scorers"
        optionLabel="name"
        :placeholder="scorers.length === 0 ? placeHolder : (inner?.name || props.placeHolder)"
        class="w-full md:w-56"
    />
    <p v-if="error" class="text-red-500 text-xs">{{ error }}</p>
</template>