import axios from 'axios'
import { getToken, API_URL } from './auth'

export async function getTeams(groupId){
    try {
        const response = await axios.get(`${API_URL}/teams/${groupId}`)
        return response.data
    } catch (error) {
        console.error('Error fetching teams:', error)
        throw error
    }
}

export async function submitTeams(teams) {
    const token = getToken()
    if (!token) {
        throw new Error("User not authenticated")
    }
    
    try {
        const response = await axios.post(`${API_URL}/pool/teams`, {'teams_id': teams}, {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        })
        return response.data
    } catch (error) {
        throw error
    }
}