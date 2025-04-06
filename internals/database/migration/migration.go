package migration

import "database/sql"

func RunMigrations(db *sql.DB) error {
	_, err := db.Exec(`
 CREATE TABLE IF NOT EXISTS passwords (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            service TEXT NOT NULL,
            password TEXT NOT NULL,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP    )
    `)

	return err
}
