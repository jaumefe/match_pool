<script setup>
import { onMounted, ref } from 'vue';
import { uploadFileScorers, uploadFileTeams, getConfiguration, getGoalsPointsPerStage, getStagesPoints, setConfiguration, submitGoalsPointsPerStage, submitStagePoints } from '../services/configuration';
import Select from 'primevue/select';


const error = ref({
  setStage: '',
  getStage: '',
  setGoal: '',
  getGoal: '',
  setConfig: '',
  getConfig: '',
  sendTeam: '',
  sendScorer: ''
})

const stages = ref(null)
const pointsGoalStage = ref(null)
const configuration = ref(null)
const formData = ref({
    stage: '',
    stagePoints: 0,
    stageGoal: '',
    goalPoints: 0
})

async function setStagePoints(){
    error.value.setStage = ''
    try {
        await submitStagePoints({
            name: formData.value?.stage,
            points: parseInt(formData.value?.stagePoints, 10)
        })
        await loadStages()
        await loadGoalsPointsPerStage()
    } catch (err){
        error.value.setStage = 'Error setting stage points'
    }
}

async function loadStages(){
  error.value.getStage = ''
  try {
  const data = await getStagesPoints()
      stages.value = data
  } catch (err) {
      error.value.getStage = 'Error al cargar los datos'
      console.log(err)
    }
}

async function setGoalsPerStagePoints(){
  error.value.setGoal = ''
    try {
      const data = await submitGoalsPointsPerStage({
        points_per_goal: parseInt(formData.value?.goalPoints, 10),
        stage_id: formData.value?.stageGoal
      })
      await loadGoalsPointsPerStage()
    } catch (err) {
        error.value.setGoal = 'Error al cargar los datos'
        console.log(err)
    }
}

async function loadGoalsPointsPerStage(){
  error.value.getGoal = ''
  try {
    const data = await getGoalsPointsPerStage()
    pointsGoalStage.value = data
  } catch (err) {
    error.value.getGoal = 'Error al cargar los datos'
    console.log(err)
  }
}

async function loadConfiguration(){
  error.value.getConfig = ''
  try {
    const data = await getConfiguration()
    configuration.value = data
  } catch (err) {
    error.value.getConfig = 'Error al cargar la configuración'
    console.log(error)
  }
}

async function updateConfiguration(config){
  error.value.setConfig = ''
  try {
    await setConfiguration(config)
    await loadConfiguration()
  } catch (err) {
    error.value.setConfig = 'Error al actualizar la configuración'
    console.log(err)
  }
}

const fileTeams = ref(null);
const fileScorers = ref(null);

function onTeamsFileChange(event) {
  fileTeams.value = event.target.files[0];
}

function onScorersFileChange(event) {
  fileScorers.value = event.target.files[0];
}

async function sendTeamsCsv() {
  if (!fileTeams.value) return;
  try {
    await uploadFileTeams(fileTeams.value);
  } catch (err) {
    console.error("Error al enviar:", err);
    error.value.sendTeam = 'Error al enviar el archivo de equipos';
  }
}

async function sendScorersCsv(){
  if (!fileScorers.value) return;
  try {
    await uploadFileScorers(fileScorers.value);
  } catch (err) {
    console.error("Error al enviar:", err);
    error.value.sendScorer = 'Error al enviar el archivo de goleadores';
  }
}

onMounted(() => {
  loadStages(),
  loadGoalsPointsPerStage(),
  loadConfiguration()
})
</script>

<template>
  <div v-if="configuration">
    <h2>Configuració del joc</h2>
    <h3>Punts per fases</h3>
    <p>
        Introdueix el nom de la fase: <input v-model="formData.stage" placeholder="Nom de la fase"></input>
         Punts: <input v-model="formData.stagePoints" placeholder="Punts per victòria"></input>
        <button @click="setStagePoints">Confirmar</button>
        <p v-if="error.setStage" class="error">{{ error.setStage }}</p>
    </p>
    <table>
      <thead>
        <tr>
          <th>Tipus d'eliminatòria</th>
          <th>Punts per victòria</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="stage in stages" :key="stage.name">
          <td>{{ stage.name }}</td>
          <td>{{ stage.points_win }}</td>
        </tr>
      </tbody>
    </table>

    <h3>Punts per gol</h3>
    <p>
      <Select v-model="formData.stageGoal" :options="stages" optionLabel="name" optionValue="id" placeholder="Nom de la fase"></Select>
        Punts: <input v-model="formData.goalPoints" placeholder="Punts per gol"></input>
        <button @click="setGoalsPerStagePoints">Confirmar</button>
        <p v-if="error.setGoal" class="error">{{ error.setGoal }}</p>
      </p>
      <table>
      <thead>
        <tr>
          <th>Tipus d'eliminatòria</th>
          <th>Punts per gol</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="pgs in pointsGoalStage" :key="pgs.stage">
          <td>{{ pgs.stage }}</td>
          <td>{{ pgs.points_per_goal }}</td>
        </tr>
      </tbody>
    </table>

    <h3>Configuració addicional</h3>
    <table>
      <thead>
        <tr>
          <th>Configuració</th>
          <th>Valor</th>
          <th>Nou valor</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="c in configuration" :key="c.key">
          <td>{{ c.key }}</td>
          <td>{{ c.value }}</td>
          <td>
            <input v-if="c.key !== 'limit_date'" v-model="c.value" placeholder="Nou valor"></input>
            <input v-if="c.key === 'limit_date'" v-model="c.value" type="date"></input>
            <button @click="updateConfiguration(c)">Confirmar</button>
            <p v-if="error.setConfig" class="error">{{ error.setConfig }}</p>
          </td>
        </tr>
      </tbody>
    </table>

    <h3>Registrar equips</h3>
    <a href="/Teams.csv" download="template_teams.csv">Descarregar plantilla per registrar equips</a>
    <p><input type="file" accept=".csv" @change="onTeamsFileChange" /></p>
    <p><button @click="sendTeamsCsv" :disabled="!fileTeams">Enviar</button></p>
    <p v-if="error.sendTeam" class="error">{{ error.sendTeam }}</p>

    <h3>Registrar jugadors</h3>
    <a href="/Scorers.csv" download="template_scorers.csv">Descarregar plantilla per registrar equips</a>
    <p><input type="file" accept=".csv" @change="onScorersFileChange" /></p>
    <p><button @click="sendScorersCsv" :disabled="!fileScorers">Enviar</button></p>
    <p v-if="error.sendScorer" class="error">{{ error.sendScorer }}</p>
  </div>
  <div v-if="error.getConfig" >
    <h2>Carregant configuració</h2>
    <p class="error"> {{ error.getConfig }}</p>
  </div>

</template>

<style scoped>
table {
  border-collapse: collapse;
  width: 100%;
  text-align: center;
}
th, td {
  padding: 8px;
}
</style>