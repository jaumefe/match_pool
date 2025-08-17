<script setup>
import { onMounted, ref } from 'vue';
import { getStagesPoints, submitStagePoints } from '../services/stage';


const error = ref('')
const stages = ref(null)
const formData = ref({
    stage: '',
    stagePoints: ''
})

async function setStagePoints(){
    error.value = ''
    try {
        await submitStagePoints({
            name: formData.value?.stage,
            points: parseInt(formData.value?.stagePoints, 10)
        })
        await loadStages()
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

onMounted(() => {
  loadStages()
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
    <h3></h3>
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