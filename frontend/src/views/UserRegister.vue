<script setup>
import Select from 'primevue/select';
import { ref, computed } from 'vue';
import { submitUser } from '../services/users';


const formData = ref({
    name: '',
    role: '',
    passwd: '',
    passwdConf: '',
})

const error = ref('')

const role = ref([
    {name: "Administrador", short: "admin"},
    {name: "Usuari", short: "user"} 
])

const passOK = computed(() => {
  return formData.value?.passwd === formData.value?.passwdConf
})

const registerReady = computed(() => {
    return passOK.value && (formData.value?.passwd !== '')
})

async function registerUser(){
    error.value = ''
    try {
        await submitUser({
            name: formData.value?.name,
            passwd: formData.value?.passwd,
            role: formData.value?.role
        })
    } catch (err){
        error.value = 'Error registering user'
    }
}
</script>

<template>
<h2>Nou usuari</h2>
    <div class="container">
        <div><input v-model="formData.name" placeholder="Nom de l'usuari"></input></div>
        <div><input type="password" v-model="formData.passwd" placeholder="Contrasenya"></input></div>
        <div>
            <input type="password" v-model="formData.passwdConf" placeholder="Confirma contrasenya"></input>
            <p v-if="!passOK">La contrasenya no és idèntica</p>
        </div>
        <div><Select v-model="formData.role" :options="role" optionLabel="name" optionValue="short" placeholder="Selecciona un rol" class="w-full md:w-56" /></div>
        <div v-if="registerReady"><button @click="registerUser">Registrar usuari</button></div>
    </div>

</template>

<style scoped>
.container {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}
</style>