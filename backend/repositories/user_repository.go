package repositories

import (
    "log"
    "database/sql"
    "github.com/bmarquescost/expense-tracker/models"
)

type UserRepository struct {
    DB *sql.DB
}

func (r *UserRepository) CreateUser(user *models.User) error {
    query := `
		INSERT INTO users (username, password) 
		VALUES (?, ?)
	`
    _, err := r.DB.Exec(query, user.ID, user.Password)
    return err
}

func (r *UserRepository) CheckForUser(user *models.User) (bool, error) {
    query := `
		SELECT COUNT(*) FROM users 
        WHERE username = ? AND password = ? 
	`
    
    rows, err := r.DB.Query(query, user.ID, user.Password)
	if err != nil {
		log.Fatal("%s", err)
		return false, err 
	}
	defer rows.Close() 
	
	var count int
	for rows.Next() {
		rows.Scan(&count)
	}
	log.Printf("Count: %d", count)

	return count != 0, nil
}