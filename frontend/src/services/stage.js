import axios from 'axios'
import { API_URL } from './auth'

export async function getStages() {
    try {
        const response = await axios.get(`${API_URL}/stages`)
        return response.data
    } catch(error){
        throw new Error("Error fetching stages")
    }
}