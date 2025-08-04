import axios from 'axios'
import { getToken, API_URL } from './auth'

export async function getMatchId(data){
    const token = getToken()
    if (!token) {
        throw new Error("User not authenticated")
    }

    try {
        const match_id = await axios.post(`${API_URL}/match/id`,
            {
                'team_home_id': data.teamHomeId,
                'team_away_id': data.teamAwayId,
                'stage_id': data.stageId
            },
            {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            })
            return match_id.data
    } catch (error) {
        console.log(error)
        throw error
    }
}