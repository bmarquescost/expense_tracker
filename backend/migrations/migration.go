package migrations

import (
    "database/sql"
	"log"
)

func Migrate(db *sql.DB) {

	query := `
	CREATE TABLE IF NOT EXISTS expense_tracker_db.users (
		username VARCHAR(255) PRIMARY KEY,
		password VARCHAR(255) NOT NULL
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	query = `
	CREATE TABLE IF NOT EXISTS expense_tracker_db.expense_types (
		expense_type VARCHAR(255) PRIMARY KEY,
		title        VARCHAR(255),
		color        VARCHAR(255)
	);
    `
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	query = `
	CREATE TABLE IF NOT EXISTS expense_tracker_db.expenses (
		username      VARCHAR(255)  NOT NULL,
		expense_title VARCHAR(255)  NOT NULL,
		title         VARCHAR(255)  NOT NULL,
		amount        DECIMAL(10,2) NOT NULL,
		date          VARCHAR(10)   NOT NULL,
		expense_type  VARCHAR(255),
		cron_schedule VARCHAR(255),
		start_date    VARCHAR(10),
		end_date      VARCHAR(10),

		PRIMARY KEY (username, expense_title),
		FOREIGN KEY (username) REFERENCES users(username) ON DELETE CASCADE,
		FOREIGN KEY (expense_type) REFERENCES expense_types(expense_type) ON DELETE SET NULL
	);`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migration completed")
}