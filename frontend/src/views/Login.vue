<script setup>

import { ref } from 'vue'
import { login } from '../services/auth.js'
import { useRouter } from 'vue-router'

const user = ref('')
const password = ref('')
const error = ref('')
const router = useRouter()

async function submit() {
  error.value = ''
  try  {
    const result = await login(user.value, password.value)
    router.push('/pool')
  } catch (err) {
    error.value = err.message
  }
}
</script>

<template>
  <div><h3>Login</h3>
    <p>User: <input v-model="user"></p>
    <p>Password: <input type="password" v-model="password"></p>
    <p><button @click="submit">Enter</button></p>
    <p v-if="error" class="text-red-600">{{ error }}</p>
  </div>
</template>
