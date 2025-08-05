<script setup>
import { ref } from 'vue'
import RegisterTeamSelect from '../components/RegisterTeamSelect.vue'
import RegisterStageSelect from '../components/RegisterStageSelect.vue'
import RegisterScorerSelect from '../components/RegisterScorerSelect.vue'
import { submitRegisterMatch, registerScorerMatch } from '../services/result_register.js'
import RegisterPenaltyWinnerSelect from '../components/RegisterPenaltyWinnerSelect.vue'

const Stage = ref(null)
const TeamHome = ref(null)
const TeamAway = ref(null)
const PenaltyWinner = ref(null)
const Scorer = ref(null)
const formData = ref({
  scoreHome: 0,
  scoreAway: 0,
  goals: 0
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
          penaltyWinnerId: PenaltyWinner.value?.id || null
      })
  } catch (err) {
      error.value = 'Error registering match'
  }
}

async function submitScorers(){
  error.value = ''
  try {
      await registerScorerMatch({
          stageId: Stage.value?.id,
          teamHomeId: TeamHome.value?.id,
          teamAwayId: TeamAway.value?.id,
          goals: formData.value.goals,
          scorerId: Scorer.value?.id
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
      <div class="column"><RegisterPenaltyWinnerSelect v-model="PenaltyWinner" :team-home="TeamHome" :team-away="TeamAway" class="mb-4"/></div>
    </div>
  <p><button @click="submitResult">Confirmar partit</button></p>
  <h2>Registrar golejadors</h2>
    <div class="container">
      <div class="column"><RegisterScorerSelect v-model="Scorer" :team-home-id="TeamHome?.id" :team-away-id="TeamAway?.id" class="mb-4" /></div>
      <div><input name="goals" :value="formData.goals" @input="onInput" placeholder="Goals"></input></div>
    </div>
<p><button @click="submitScorers">Confirmar golejador</button></p>
</template>

<style scoped>
.container {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.column {
  flex: 1 1 calc(16.67% - 10px); /* 5 columnas por fila */
  box-sizing: border-box;
  padding: 10px;
  text-align: center;
  min-width: 160px; /* evita que se compriman demasiado */
}
</style>