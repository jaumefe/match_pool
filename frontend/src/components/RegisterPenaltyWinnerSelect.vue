<script setup>
import { ref, computed } from 'vue'
import Select from 'primevue/select'

const props = defineProps({
    teamHome: {
        type: Object,
        required: true
    },
    teamAway: {
        type: Object,
        required: true
    },
    stageId: {
        type: Number,
        required: true
    },
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

const teams = computed(() => {
    const list = []
  if (props.teamHome) list.push(props.teamHome)
  if (props.teamAway) list.push(props.teamAway)
  return list
})


const inner = computed({
  get: () => props.modelValue,
  set: val => emit('update:modelValue', val)
})

</script>

<template>
    <Select
        v-model="inner"  
        :options="teams"
        optionLabel="name"
        :placeholder="teams.length === 0 ? placeHolder : (inner?.name || props.placeHolder)"
        class="w-full md:w-56"
    />
    <p v-if="error" class="text-red-500 text-xs">{{ error }}</p>
</template>