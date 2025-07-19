<script setup>
import TeamSelect from '../components/TeamSelect.vue'
import { ref } from 'vue'
import { submitTeams } from '../services/teams.js'

const selectedA = ref(null)
const selectedB1 = ref(null)
const selectedB2 = ref(null)
const selectedC1 = ref(null)
const selectedC2 = ref(null)
const selectedD1 = ref(null)
const selectedD2 = ref(null)
const error = ref('')

selectedA.value = { value: 0 }
selectedB1.value = { value: 0 }
selectedB2.value = { value: 0 }
selectedC1.value = { value: 0 }
selectedC2.value = { value: 0 }
selectedD1.value = { value: 0 }
selectedD2.value = { value: 0 }

async function submit(){  
  error.value = ''
  try {
    const teams = [
     selectedA.value?.id,
    selectedB1.value?.id,
    selectedB2.value?.id,
    selectedC1.value?.id,
    selectedC2.value?.id,
    selectedD1.value?.id,
    selectedD2.value?.id
    ]
    const result = await submitTeams(teams)
  } catch (err) {
    error.value = 'Error assigning teams'
  }
}
</script>

<template>
  <h2>Asignar Equipos</h2>

  <div class="container">
    <div class="column"><TeamSelect v-model="selectedA"  :group-id="1" class="mb-4" /><p> {{ selectedA.value }}</p></div>
    <div class="column"><TeamSelect v-model="selectedB1" :group-id="2" class="mb-4" /><p> {{ selectedB1.value }}</p></div>
    <div class="column"><TeamSelect v-model="selectedB2" :group-id="2" class="mb-4" /><p> {{ selectedB2.value }}</p></div>
    <div class="column"><TeamSelect v-model="selectedC1" :group-id="3" class="mb-4" /><p> {{ selectedC1.value }}</p></div>
    <div class="column"><TeamSelect v-model="selectedC2" :group-id="3" class="mb-4" /><p> {{ selectedC2.value }}</p></div>
    <div class="column"><TeamSelect v-model="selectedD1" :group-id="4" class="mb-4" /><p> {{ selectedD1.value }}</p></div>
    <div class="column"><TeamSelect v-model="selectedD2" :group-id="4" class="mb-4" /><p> {{ selectedD2.value }}</p></div>
  </div>
  <p>TOTAL: {{ selectedA.value + selectedB1.value + selectedB2.value + selectedC1.value + selectedC2.value + selectedD1.value + selectedD2.value }}</p>
  <p><button @click="submit">Confirmar</button></p>
</template>


<style scoped>
.container {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.column {
  flex: 1 1 calc(14.2857% - 10px); /* 7 columnas por fila */
  box-sizing: border-box;
  padding: 10px;
  text-align: center;
  min-width: 160px; /* evita que se compriman demasiado */
}
</style>
