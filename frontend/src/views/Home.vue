<script setup>
defineProps({
  msg: String
})

import { useRouter } from 'vue-router'

const router = useRouter()

function Login() {
  router.push('/login')
}

</script>

<template>
    <h1>{{ msg }}</h1>
    <div class="card">
        <p>
            This is the home page of the Match Pool application.
        </p>
        <p>
            <button @click="Login">Login</button>
        </p>
    </div>
</template>