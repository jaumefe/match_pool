import axios from 'axios'
import { getToken, API_URL } from './auth'
import { getMatchId } from './match'

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
        console.log(error)
        throw error
    }
}

export async function registerScorerMatch(scorerData){
    const token = getToken()
    if (!token) {
        throw new Error("User not authenticated")
    }

    const match_id = await getMatchId({
        teamHomeId: scorerData.teamHomeId,
        teamAwayId: scorerData.teamAwayId,
        stageId: scorerData.stageId
    })
    try {
        const response = await axios.post(`${API_URL}/goals/id`, 
                {
                    'player_id': scorerData.scorerId,
                    'match_id': match_id.match_id,
                    'goals': scorerData.goals
                },
                {
                    headers: {
                        'Authorization': `Bearer ${token}`
                    }
                })
        return response.data
    } catch (error) {
        console.log(error)
        throw error
    }
}

export async function submitTeamPosition(positionData){
    const token = getToken()
    if (!token) {
        throw new Error("User not authenticated")
    }

    try {
        const response = await axios.post(`${API_URL}/teams/pool_position`, 
            {
                'id': positionData.id,
                'pool_position': positionData.poolPosition
            },
            {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            })
        return response.data
    } catch (error) {
        console.log(error)
        throw error
    }
}