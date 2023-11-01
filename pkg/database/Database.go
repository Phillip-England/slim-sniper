package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func GetDatabase() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	return db, nil
}

func DeleteTable(db *sql.DB, tableName string) error {
    query := fmt.Sprintf("DROP TABLE IF EXISTS %s", tableName)
    _, err := db.Exec(query)
    if err != nil {
        return err
    }
    return nil
}

func InitCemTable(db *sql.DB) error {
	tableExistsQuery := `
		SELECT EXISTS (
			SELECT 1
			FROM information_schema.tables
			WHERE table_name = 'cem_score'
		);
	`
	var tableExists bool
	err := db.QueryRow(tableExistsQuery).Scan(&tableExists)
	if err != nil {
		return err
	}
	if !tableExists {
		createTableQuery := `
			CREATE TABLE cem_score (
				id SERIAL PRIMARY KEY,
				timescale TEXT NOT NULL,
				store TEXT NOT NULL,
				osat DOUBLE PRECISION NOT NULL,
				taste DOUBLE PRECISION NOT NULL,
				ace DOUBLE PRECISION NOT NULL,
				speed DOUBLE PRECISION NOT NULL,
				cleanliness DOUBLE PRECISION NOT NULL,
				accuracy DOUBLE PRECISION NOT NULL
			);
		`
		_, err := db.Exec(createTableQuery)
		if err != nil {
			return err
		}
	}
	return nil
}
