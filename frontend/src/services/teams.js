import axios from 'axios'
const API_URL = 'http://localhost:5050'

export async function getTeams(groupId){
    try {
        const response = await axios.get(`${API_URL}/teams/${groupId}`)
        return response.data
    } catch (error) {
        console.error('Error fetching teams:', error)
        throw error
    }
}