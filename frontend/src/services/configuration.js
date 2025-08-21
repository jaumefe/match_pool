import axios from 'axios'
import { API_URL, getToken } from './auth'

export async function getStagesPoints() {
    try {
        const response = await axios.get(`${API_URL}/stages/points`)
        return response.data
    } catch(error){
        throw new Error("Error fetching stages")
    }
}

export async function submitStagePoints(data) {
    const token = getToken()
    if (!token) {
        throw new Error("User not authenticated")
    }

    try {
        const response = await axios.post(`${API_URL}/configuration/stage`,
            {
                'name': data.name,
                'points_win': data.points
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

export async function getGoalsPointsPerStage() {
    const token = getToken()
    if (!token) {
        throw new Error("User not authenticated")
    }

    try {
        const response = await axios.get(`${API_URL}/goals/stage`,
            {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            }
        )
        return response.data
    } catch(error){
        throw new Error("Error fetching goals per stages")
    }
}

export async function submitGoalsPointsPerStage(data){
    const token = getToken()
    if (!token) {
        throw new Error("User not authenticated")
    }
    try {
        const response = await axios.post(`${API_URL}/configuration/goals`,
            {
                'points_per_goal': data.points_per_goal,
                'stage_id': data.stage_id
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

export async function getConfiguration() {
    const token = getToken()
    if (!token) {
        throw new Error("User not authenticated")
    }
    try {
        const response = await axios.get(`${API_URL}/configuration`,
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

export async function setConfiguration(config) {
    const token = getToken()
    if (!token) {
        throw new Error("User not authenticated")
    }
    try {
        const response = await axios.post(`${API_URL}/configuration`,
            {
                'key': config.key,
                'value': config.value
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

export async function uploadFileTeams(file){
    const token = getToken()
    if (!token) {
        throw new Error("User not authenticated")
    }

    const formData = new FormData()
    formData.append('file', file)

    try {
        const response = await axios.post(`${API_URL}/register/teams`, formData, {
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'multipart/form-data'
            }
        })
        return response.data
    } catch (error) {
        console.log(error)
        throw error
    }
}

export async function uploadFileScorers(file){
    const token = getToken()
    if (!token) {
        throw new Error("User not authenticated")
    }

    const formData = new FormData()
    formData.append('file', file)

    try {
        const response = await axios.post(`${API_URL}/register/scorers`, formData, {
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'multipart/form-data'
            }
        })
        return response.data
    } catch (error) {
        console.log(error)
        throw error
    }
}