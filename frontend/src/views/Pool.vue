<script setup>
import TeamSelect from '../components/TeamSelect.vue'
import ScorerSelect from '../components/ScorerSelect.vue'
import { ref, onMounted } from 'vue'
import { submitTeams, getTeamByUser } from '../services/teams.js'
import { submitScorers, getScorerByUser } from '../services/scorers.js'

const selectedA = ref(null)
const selectedB1 = ref(null)
const selectedB2 = ref(null)
const selectedC1 = ref(null)
const selectedC2 = ref(null)
const selectedD1 = ref(null)
const selectedD2 = ref(null)
const selectedScorerA = ref(null)
const selectedScorerB1 = ref(null)
const selectedScorerB2 = ref(null)
const selectedScorerC1 = ref(null)
const selectedScorerC2 = ref(null)
const selectedScorerC3 = ref(null)
const selectedScorerD1 = ref(null)
const selectedScorerD2 = ref(null)
const error = ref('')

// Initialize selected teams with values from getTeamByUser
// getTeamByUser is assumed to return a promise that resolves to an object with team IDs
onMounted(async() => {
  const userTeams = await getTeamByUser()
  if (userTeams) {
    selectedA.value = userTeams[0] || { value: 0 }
    selectedB1.value = userTeams[1] || { value: 0 }
    selectedB2.value = userTeams[2] || { value: 0 }
    selectedC1.value = userTeams[3] || { value: 0 }
    selectedC2.value = userTeams[4] || { value: 0 }
    selectedD1.value = userTeams[5] || { value: 0 }
    selectedD2.value = userTeams[6] || { value: 0 }
  }

  const userScorers = await getScorerByUser()
  if (userScorers) {
    selectedScorerA.value = userScorers[0]
    selectedScorerB1.value = userScorers[1]
    selectedScorerB2.value = userScorers[2]
    selectedScorerC1.value = userScorers[3]
    selectedScorerC2.value = userScorers[4]
    selectedScorerC3.value = userScorers[5]
    selectedScorerD1.value = userScorers[6]
    selectedScorerD2.value = userScorers[7]
  }
})

selectedA.value = { value: 0 }
selectedB1.value = { value: 0 }
selectedB2.value = { value: 0 }
selectedC1.value = { value: 0 }
selectedC2.value = { value: 0 }
selectedD1.value = { value: 0 }
selectedD2.value = { value: 0 }

async function submitTeamsUser(){  
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

async function submitScorerUser(){
  error.value = ''
  try {
    const scorers = [
      selectedScorerA.value?.id,
      selectedScorerB1.value?.id,
      selectedScorerB2.value?.id,
      selectedScorerC1.value?.id,
      selectedScorerC2.value?.id,
      selectedScorerC3.value?.id,
      selectedScorerD1.value?.id,
      selectedScorerD2.value?.id
    ]
    const result = await submitScorers(scorers)
  } catch (err) {
    error.value = "Error assigning scorers"
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
  <p><button @click="submitTeamsUser">Confirmar equips</button></p>
  <div class="container">
    <div class="column"><ScorerSelect v-model="selectedScorerA" :group-id="1" class="mb-4" /></div>
    <div class="column"><ScorerSelect v-model="selectedScorerB1" :group-id="2" class="mb-4" /></div>
    <div class="column"><ScorerSelect v-model="selectedScorerB2" :group-id="2" class="mb-4" /></div>
    <div class="column"><ScorerSelect v-model="selectedScorerC1" :group-id="3" class="mb-4" /></div>
    <div class="column"><ScorerSelect v-model="selectedScorerC2" :group-id="3" class="mb-4" /></div>
    <div class="column"><ScorerSelect v-model="selectedScorerC3" :group-id="3" class="mb-4" /></div>
    <div class="column"><ScorerSelect v-model="selectedScorerD1" :group-id="4" class="mb-4" /></div>
    <div class="column"><ScorerSelect v-model="selectedScorerD2" :group-id="4" class="mb-4" /></div>
  </div>
  <p><button @click="submitScorerUser">Confirmar jugadors</button></p>
</template>


<style scoped>
.container {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.column {
  flex: 1 1 calc(25% - 10px); /* 7 columnas por fila */
  box-sizing: border-box;
  padding: 10px;
  text-align: center;
  min-width: 160px; /* evita que se compriman demasiado */
}
</style>
