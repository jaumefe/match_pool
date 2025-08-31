<script setup>
import { ref, onMounted, computed } from 'vue'
import { getPoints } from '../services/points'

const error = ref('')
const points = ref(null)

onMounted(async () => {
  try {
    const data = await getPoints()
    points.value = data
  } catch (err) {
    error.value = 'Error loading points'
    console.error(err)
  }
})

// Convertir a array y ordenar por total
const sortedPoints = computed(() => {
  if (!points.value) return []
  return Object.entries(points.value)
    .map(([nombre, valores]) => ({ nombre, ...valores }))
    .sort((a, b) => b.total - a.total)
})
</script>

<template>
  <div>
    <p v-if="error" class="error">{{ error }}</p>

    <table v-if="sortedPoints.length">
      <thead>
        <tr>
          <th>Usuari</th>
          <th>Punts gols</th>
          <th>Punts equips</th>
          <th>Total</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="jugador in sortedPoints" :key="jugador.nombre">
          <td>{{ jugador.nombre }}</td>
          <td>{{ jugador.scorers }}</td>
          <td>{{ jugador.teams }}</td>
          <td><strong>{{ jugador.total }}</strong></td>
        </tr>
      </tbody>
    </table>
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