package database

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"

	"example.com/go/backend/models"
)

func GetUserByUsername(username string, db *sql.DB) (*models.User, error) {
	var user models.User
	err := db.QueryRow("SELECT id, name, role, email, phone FROM user_view WHERE name = ?", username).Scan(&user.ID, &user.Name, &user.Role, &user.Email, &user.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found
		}
		return nil, err
	}
	return &user, nil
}

func verifyPassword(providedPassword, storedHash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(providedPassword))
	return err
}

func AuthenticateUser(username string, password string, db *sql.DB) (bool, error) {
	user, err := GetUserByUsername(username, db)
	if err != nil {
		log.Warn().Msgf("Error occurred while fetching user: %s", err)

		http.Redirect(nil, nil, "/error", http.StatusNotImplemented)
		return false, err
	}
	if user == nil {
		return false, nil // User not found
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return false, fmt.Errorf("Error occurred while hashing password: %w", err)
	}
	fmt.Printf("Hashed password for debugging: %s\n", string(hashedPassword))

	return verifyPassword(password, string(hashedPassword)) == nil, nil
}
