package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"

	. "example.com/go/backend/middleware"
)

func GetUserByUsername(username string) (*User, error) {
	var user User
	err := DB.QueryRow("SELECT id, name, role, email FROM users WHERE name = ?", username).Scan(&user.ID, &user.Name, &user.Role, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found
		}
		return nil, fmt.Errorf("error occurred while fetching user: %w", err)
	}
	return &user, nil
}

func verifyPassword(providedPassword, storedHash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(providedPassword))
	return err
}

func AuthenticateUser(username, password string) (bool, error) {
	user, err := GetUserByUsername(username)
	if err != nil {
		log.Warn().Msgf("Error occurred while fetching user: %s", err)
		// log.Warn("Error occurred while fetching user: %w", err) // This is the correct way to log the error with zerolog
		return false, fmt.Errorf("error occurred while authenticating user: %w", err)
	}
	if user == nil {
		return false, nil // User not found
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return false, fmt.Errorf("error occurred while hashing password: %w", err)
	}
	fmt.Printf("Hashed password for debugging: %s\n", string(hashedPassword))

	return verifyPassword(password, string(hashedPassword)) == nil, nil
}
