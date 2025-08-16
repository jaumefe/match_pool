<script setup>
import { ref, computed } from 'vue';
import { updateUserName } from '../services/users';

const error = ref('')

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
    error.value = ''
    try {
        await updateUserName({
            name: formData.value.name
        })
    } catch (err) {
        error.value = 'Error updating name'
    }
}
</script>

<template>
    <h2>Perfil de l'usuari</h2>
    <p>Nom</p>
    <p>Canviar nom: <input v-model="formData.name" placeholder="Nou nom d'usuari"></input> <button v-if="newNameOK" @click="changeName">Confirmar canvi de nom</button></p>
    <p>Canviar contrasenya: <input type="password" v-model="formData.newPass" placeholder="Nova contrasenya"></input></p>
    <p>Confirmar contrasenya: <input type="password" v-model="formData.newPassConf" placeholder="Nova contrasenya"></input></p>
    <p v-if="passConfirm"><button @click="changePass">Confirmar canvi de contrasenya</button></p>
    <p v-if="!newPassOK">La contrasenya no és idèntica</p>
</template>