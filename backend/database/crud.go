package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"

	"example.com/go/backend/models"
)

func GetUserByUsername(username string, db *sql.DB) (models.User, error) {
	var user models.User
	err := db.QueryRow("SELECT ID, username, password, email, phone, role FROM user_view WHERE username = ?", username).Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.Phone, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, nil // No user found
		}
		return models.User{}, err
	}
	return user, nil
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

func AuthenticateUser(username string, password string, db *sql.DB) (models.User, error) {
	user, err := GetUserByUsername(username, db)
	if err != nil {
		log.Warn().Msgf("Error occurred while fetching user: %s", err)
		return models.User{Name: ""}, err
	}
	if user.Name == "" {
		return models.User{Name: ""}, nil // User not found
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{Name: ""}, fmt.Errorf("Error occurred while hashing password: %w", err)
	}
	fmt.Printf("Hashed password for debugging: %s\n", string(hashedPassword))

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		return models.User{Name: ""}, err
	}

	return user, nil
}

func CreateUser(SignupReq models.SignupRequest, db *sql.DB) {

}
