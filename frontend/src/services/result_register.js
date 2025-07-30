import axios from 'axios'
import { getToken, API_URL } from './auth'

export async function submitRegisterMatch(matchData) {
    const token = getToken()
    if (!token) {
        throw new Error("User not authenticated")
    }
    
    try {
        const response = await axios.post(`${API_URL}/matches/id`, 
            {
                'leg': matchData.leg,
                'stage_id': matchData.stageId,
                'team_home_id': matchData.teamHomeId,
                'team_away_id': matchData.teamAwayId,
                'team_home_score': matchData.teamHomeScore,
                'team_away_score': matchData.teamAwayScore,
                'penalty_winner_id': matchData.penaltyWinnerId
            },
            {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            })
        return response.data
    } catch (error) {
        throw error
    }
}
