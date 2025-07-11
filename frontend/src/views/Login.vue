<script setup>

import { ref } from 'vue'
import { login } from '../services/auth.js'

const user = ref('')
const password = ref('')
const error = ref('')

async function submit() {
  error.value = ''
  const result = await login(user.value, password.value)
  if (result) {
    // Handle successful login
    console.log('Login successful')
  } else {
    error.value = 'Login failed. Please try again.'
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
