package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(dataSource string) error {
	var err error
	DB, err = sql.Open("sqlite3", dataSource)
	if err != nil {
		return err
	}

	// Create tables if they do not exist
	if err := createTables(); err != nil {
		return err
	}

	return nil
}

func createTables() error {
	queries := []string{
		`PRAGMA foreign_keys = ON;`,

		`CREATE TABLE IF NOT EXISTS match (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			tie_id INTEGER, -- ID of the tie this match belongs to
			leg INTEGER DEFAULT 1, -- 1 for first leg, 2 for second leg
			stage_id INTEGER NOT NULL,
			team_home_id INTEGER NOT NULL,
			team_away_id INTEGER NOT NULL,
			team_home_score INTEGER DEFAULT 0,
			team_away_score INTEGER DEFAULT 0,
			penalty_winner_id INTEGER,
			FOREIGN KEY (team_home_id) REFERENCES teams(id),
			FOREIGN KEY (team_away_id) REFERENCES teams(id),
			FOREIGN KEY (penalty_winner_id) REFERENCES teams(id),
			FOREIGN KEY (stage_id) REFERENCES stage(id)
		);`,

		`CREATE TABLE IF NOT EXISTS teams (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			group_name TEXT NOT NULL,
			value INTEGER NOT NULL,
			pool_position INTEGER,
			pool_group TEXT
		);`,

		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			token TEXT NOT NULL,
			role TEXT NOT NULL DEFAULT 'user',
			score INTEGER DEFAULT 0
		);`,

		`CREATE TABLE IF NOT EXISTS user_teams (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			team_id INTEGER NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (team_id) REFERENCES teams(id)
		);`,

		`CREATE TABLE IF NOT EXISTS stage (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			points_win INTEGER NOT NULL DEFAULT 3
		);`,

		`CREATE TABLE IF NOT EXISTS scorers (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			team_id INTEGER NOT NULL,
			position TEXT NOT NULL,
			FOREIGN KEY (team_id) REFERENCES teams(id)
		);`,

		`CREATE TABLE IF NOT EXISTS player_goals (
			player_id INTEGER NOT NULL,
			match_id INTEGER NOT NULL,
			goals INTEGER NOT NULL,
			FOREIGN KEY (player_id) REFERENCES scorers(id),
			FOREIGN KEY (match_id) REFERENCES match(id),
			PRIMARY KEY (player_id, match_id)
		);`,

		`CREATE TABLE IF NOT EXISTS user_scorers (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			player_id INTEGER NOT NULL,
			user_id INTEGER NOT NULL,
			FOREIGN KEY (player_id) REFERENCES scorers(id),
			FOREIGN KEY (user_id) REFERENCES users(id)
		);`,

		`CREATE TABLE IF NOT EXISTS points_goals (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			pointsPerGoal INTEGER NOT NULL DEFAULT 1,
			stage_id INTEGER NOT NULL,
			FOREIGN KEY (stage_id) REFERENCES stage(id)
		);`,
	}

	for _, q := range queries {
		if _, err := DB.Exec(q); err != nil {
			return err
		}
	}
	return nil
}
