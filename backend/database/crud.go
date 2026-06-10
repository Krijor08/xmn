package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"

	"example.com/go/backend/models"
)

func GetUserByUsername(username string, db *sql.DB) (*models.LoginRequest, error) {
	var user models.LoginRequest
	err := db.QueryRow("SELECT username, password FROM user_view WHERE username = ?", username).Scan(&user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found
		}
		return nil, err
	}
	return &user, nil
}

func GetRoleId(roleName string, db *sql.DB) (int, error) {
	var role models.Role
	err := db.QueryRow("SELECT ID, role FROM roles WHERE role = ?", roleName).Scan(&role.ID, &role.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return role.ID, nil
}

func AuthenticateUser(username string, password string, db *sql.DB) (bool, error) {
	user, err := GetUserByUsername(username, db)
	if err != nil {
		log.Warn().Msgf("Error occurred while fetching user: %s", err)
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

	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)) == nil, nil
}

func CreateUser(SignupReq models.SignupRequest, db *sql.DB) {

}
