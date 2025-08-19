<script setup>
import { onMounted, ref } from 'vue';
import { getConfiguration, getGoalsPointsPerStage, getStagesPoints, setConfiguration, submitGoalsPointsPerStage, submitStagePoints } from '../services/configuration';
import Select from 'primevue/select';


const error = ref('')
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
    error.value = ''
    try {
        await submitStagePoints({
            name: formData.value?.stage,
            points: parseInt(formData.value?.stagePoints, 10)
        })
        await loadStages()
        await loadGoalsPointsPerStage()
    } catch (err){
        error.value = 'Error setting stage points'
    }
}

async function loadStages(){
  error.value = ''
  try {
  const data = await getStagesPoints()
      stages.value = data
  } catch (err) {
      error.value = 'Error al cargar los datos'
      console.log(err)
    }
}

async function setGoalsPerStagePoints(){
  error.value = ''
    try {
      const data = await submitGoalsPointsPerStage({
        points_per_goal: parseInt(formData.value?.goalPoints, 10),
        stage_id: formData.value?.stageGoal
      })
      await loadGoalsPointsPerStage()
    } catch (err) {
        error.value = 'Error al cargar los datos'
        console.log(err)
    }
}

async function loadGoalsPointsPerStage(){
  error.value = ''
  try {
    const data = await getGoalsPointsPerStage()
    pointsGoalStage.value = data
  } catch (err) {
    error.value = 'Error al cargar los datos'
    console.log(err)
  }
}

async function loadConfiguration(){
  error.value = ''
  try {
    const data = await getConfiguration()
    configuration.value = data
  } catch (err) {
    error.value = 'Error al cargar la configuración'
    console.log(error)
  }
}

async function updateConfiguration(config){
  error.value = ''
  try {
    await setConfiguration(config)
    await loadConfiguration()
  } catch (err) {
    error.value = 'Error al actualizar la configuración'
    console.log(err)
  }
}

onMounted(() => {
  loadStages(),
  loadGoalsPointsPerStage(),
  loadConfiguration()
})
</script>

<template>
    <h2>Configuració del joc</h2>
    <h3>Punts per fases</h3>
    <p>
        Introdueix el nom de la fase: <input v-model="formData.stage" placeholder="Nom de la fase"></input>
         Punts: <input v-model="formData.stagePoints" placeholder="Punts per victòria"></input>
        <button @click="setStagePoints">Confirmar</button>
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
          </td>
        </tr>
      </tbody>
    </table>
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