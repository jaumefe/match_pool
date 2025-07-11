import axios from 'axios'
import Cookies from 'js-cookie'

const API_URL = 'http://localhost:5050'
const TOKEN_COOKIE = 'jwt'
const ONE_DAY = 1

/*
Login function to authenticate a user.
*/
export async function login(username, password) {
    try {
        const response = await axios.post(`${API_URL}/login`, { 'name': username, 'password': password })
        const token = response.data.bearer
        Cookies.set(TOKEN_COOKIE, token, { expires: ONE_DAY })

        return {token, message: response.data.message}
    } catch (error) {
        if (error.response?.status === 401){
            throw new Error('Invalid username or password')
        }
        throw new Error('An error occurred while trying to log in')
    }
}


export function getToken () {
  return Cookies.get(TOKEN_COOKIE) || null
}

export function logout () {
  Cookies.remove(TOKEN_COOKIE)
}