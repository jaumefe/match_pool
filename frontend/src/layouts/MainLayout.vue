<script setup>
import { ref } from 'vue'
import { removeToken } from '../services/auth'
import { useRouter } from 'vue-router'

const error = ref('')

const userRole = ref(localStorage.getItem('role'))
const isAdmin = userRole.value === 'admin'
const router = useRouter()

async function logout() {
    error.value = ''
    try {
        const result = await removeToken()
        router.push('/')
    } catch (err) {
        error.value = err.message
        console.log(error)
    }
}
</script>
<template>
  <div>
    <nav>
      <router-link to="/profile">Perfil</router-link>
      <router-link to="/pool">Porra</router-link>
      <router-link to="/points">Classificació</router-link>
      <router-link v-if="isAdmin" to="/result_register">Registrar partis/gols</router-link>
      <router-link v-if="isAdmin" to="/register_user">Registrar usuaris</router-link>

      <button @click="logout">Logout</button>
    </nav>
    <main>
      <router-view />
    </main>
  </div>
</template>

<style scoped>
nav {
  background: #333;
  color: white;
  padding: 1rem;
}
nav a {
  margin-right: 1rem;
  color: white;
}
</style>
