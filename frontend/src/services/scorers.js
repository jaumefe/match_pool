import axios from 'axios'
import { getToken, API_URL } from './auth'

export async function getScorersByGroupId(groupId) {
    const token = getToken()
    if (!token) {
        throw new Error("User not authenticated")
    }

    try {
        const response = await axios.get(`${API_URL}/scorers/${groupId}`,
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

export async function submitScorers(scorers) {
    const token = getToken()
    if (!token) {
        throw new Error("User not authenticated")
    }
    
    try {
        const response = await axios.post(`${API_URL}/pool/scorers`, {'scorers_id': scorers}, {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        })
        return response.data
    } catch (error) {
        throw error
    }
}

export async function getScorerByUser() {
    const token = getToken()
    if (!token) {
        throw new Error("User not authenticated")
    }

    try {
        const response = await axios.get(`${API_URL}/pool/scorers`, {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        })
        return response.data
    } catch (error) {
        console.error('Error fetching team by user:', error)
        throw error
    }
    
}

export async function getScorerByTeamId(teamId) {
    const token = getToken()
    if (!token) {
        throw new Error("User not authenticated")
    }

    try {
        const response = await axios.get(`${API_URL}/scorers/teams/${teamId}`,
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