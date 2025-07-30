<script setup>
import { ref } from 'vue'
import RegisterTeamSelect from '../components/RegisterTeamSelect.vue'
import RegisterStageSelect from '../components/RegisterStageSelect.vue'
import { submitRegisterMatch } from '../services/result_register.js'

const Stage = ref(null)
const TeamHome = ref(null)
const TeamAway = ref(null)
const formData = ref({
  scoreHome: 0,
  scoreAway: 0,
  penaltyWinner: null
})
const error = ref('')

function onInput(e) {
  const { name, value } = e.target
  if (/^\d*$/.test(value)) {
    formData.value[name] = value === '' ? 0 : parseInt(value, 10)
  }
}

async function submitResult(){
  error.value = ''
  try {
      await submitRegisterMatch({
          leg: 1,
          stageId: Stage.value?.id,
          teamHomeId: TeamHome.value?.id,
          teamAwayId: TeamAway.value?.id,
          teamHomeScore: formData.value.scoreHome,
          teamAwayScore: formData.value.scoreAway,
          penaltyWinnerId: formData.value.penaltyWinner?.id || null
      })
  } catch (err) {
      error.value = 'Error registering match'
  }
}

</script>

<template>
<h2>Registrar partits jugats</h2>
  <div class="container">
    <div class="column"><RegisterStageSelect v-model="Stage" class="mb-4" /></div>
    <div class="column"><RegisterTeamSelect v-model="TeamHome" class="mb-4" /></div>
    <div><input name="scoreHome" :value="formData.scoreHome" @input="onInput" placeholder="Home Goals"></input></div>
    <div class="column"><RegisterTeamSelect v-model="TeamAway" class="mb-4" /></div>
    <div><input name="scoreAway" :value="formData.scoreAway" @input="onInput" placeholder="Away Goals"></input></div>
  </div>
  <p><button @click="submitResult">Confirmar partit</button></p>
</template>

<style scoped>
.container {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.column {
  flex: 1 1 calc(20% - 10px); /* 5 columnas por fila */
  box-sizing: border-box;
  padding: 10px;
  text-align: center;
  min-width: 160px; /* evita que se compriman demasiado */
}
</style>