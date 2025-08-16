import axios from 'axios'
import { getToken, API_URL } from './auth'

export async function submitUser(data){
    const token = getToken()
    if (!token) {
        throw new Error("User not authenticated")
    }
    try {
        const response = await axios.post(`${API_URL}/users`,
            {
                'name': data.name,
                'password': data.passwd,
                'role': data.role
            },
            {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            }
        )
        return response.data
    } catch (error) {
        console.log(error)
        throw error
    }
}

export async function updateUserName(data){
    const token = getToken()
    if (!token) {
        throw new Error("User not authenticated")
    }

    try {
        const response = await axios.post(`${API_URL}/users/name`,
            {
                'name': data.name,
            },
            {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            }
        )
        return response.data
    } catch (error) {
        console.log(error)
        throw error
    }
    
}

export async function getUserName(){
    const token = getToken()
    if (!token) {
        throw new Error("User not authenticated")
    }

    try {
        const response = await axios.get(`${API_URL}/users/name`,
            {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            }
        )
        return response.data
    } catch (error) {
        console.log(error)
        throw error
    }
}