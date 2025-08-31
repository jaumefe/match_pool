<script setup>
import { ref, computed, onMounted } from 'vue';
import { getUserName, updateUserName, updatePasswd } from '../services/users';

const error = ref({
    name: '',
    pass: ''
})
const name = ref('')

const formData = ref({
    name: '',
    newPass: '',
    newPassConf:''
})

const newNameOK = computed(() => {
    return formData.value?.name !== ''
})

const newPassOK = computed(() => {
    return formData.value?.newPass === formData.value?.newPassConf
})

const passConfirm = computed(() => {
    return newPassOK.value && (formData.value?.newPass !== '')
})

async function changeName(){
    error.value.name = ''
    try {
        await updateUserName({
            name: formData.value.name
        })
        loadName()
    } catch (err) {
        error.value.name = 'Error updating name'
    }
}

async function changePass(){
    error.value.pass = ''
    try {
        await updatePasswd({
            token: formData.value.newPass
        })
    } catch (err) {
        error.value.pass = 'Error updating password'
    }
}

async function loadName(){
    const userName = await getUserName()
        if (userName) {
            name.value = userName
        }
}

onMounted(() => {
    loadName()
})
</script>

<template>
    <h2>Perfil de l'usuari</h2>
    <p><b>Nom: </b>{{ name }}</p>
    <p><b>Canviar nom: </b><input v-model="formData.name" placeholder="Nou nom d'usuari"></input> <button v-if="newNameOK" @click="changeName">Confirmar canvi de nom</button></p>
    <p v-if="error.name" class="error">{{ error.name }}</p>
    <p><b>Canviar contrasenya: </b><input type="password" v-model="formData.newPass" placeholder="Nova contrasenya"></input></p>
    <p><b>Confirmar contrasenya: </b><input type="password" v-model="formData.newPassConf" placeholder="Nova contrasenya"></input></p>
    <p v-if="passConfirm"><button @click="changePass">Confirmar canvi de contrasenya</button></p>
    <p v-if="!newPassOK">La contrasenya no és idèntica</p>
    <p v-if="error.pass" class="error">{{ error.pass }}</p>
</template>