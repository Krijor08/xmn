package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func init() {
	if os.Getenv("DOCKER_ENV") != "true" {
		_ = godotenv.Load()
	}
}

func getDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local&timeout=%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
		os.Getenv("MYSQL_TIMEOUT"),
	)
}

var DB *sql.DB

func InitMySQL() {
	dsn := getDSN()

	fmt.Printf("Attempting MySQL connection with DSN: %s\n", dsn)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Warn().Msg("Failed to open MySQL connection: " + err.Error())
	}

	// Critical production settings
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)
	DB.SetConnMaxLifetime(5 * time.Minute)
	DB.SetConnMaxIdleTime(10 * time.Minute)

	// Retry connection
	for i := range 10 {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		if err = DB.PingContext(ctx); err == nil {
			cancel()
			log.Info().Msg("MySQL connected successfully!")
			return
		}
		cancel()
		log.Printf("MySQL not ready, retry %d/10...\n", i+1)
		log.Printf("Error: %s\n", err.Error())
		time.Sleep(time.Duration(i+1) * time.Second)
	}

	log.Warn().Msg("Could not connect to MySQL after retries")
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}
