import axios from 'axios'
import { getToken, API_URL } from './auth'

export async function getPoints(){
    const token = getToken()
    if (!token) {
        throw new Error("User not authenticated")
    }

    try {
        const response = await axios.get(`${API_URL}/pool/points`,
            {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            })
        return response.data
    } catch (error){
        console.log(error)
        throw error
    }
}