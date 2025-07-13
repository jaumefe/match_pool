<script setup>
import TeamSelect from '../components/TeamSelect.vue'
import { ref } from 'vue'

const selectedA = ref(null)
const selectedB = ref(null)
const selectedC = ref(null)
</script>

<template>
  <h2 class="mb-4 text-lg font-bold">Asignar Equipos</h2>

  <TeamSelect v-model="selectedA" :group-id="1" class="mb-4" />
  <TeamSelect v-model="selectedB" :group-id="2" class="mb-4" />
  <TeamSelect v-model="selectedC" :group-id="3" class="mb-4" />

</template>