import axios from 'axios'
import { getToken, API_URL } from './auth'

export async function getScorers(groupId) {
    const token = getToken()
    if (!token) {
        throw new Error("User not authenticated")
    }

    try {
        const response = await axios.get(`${API_URL}/scorers/${groupId}/`,
            {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        })
        return response.data
    } catch (error) {
        console.error('Error fetching scorers:', error)
        throw error
    }
}