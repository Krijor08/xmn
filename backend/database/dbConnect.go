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
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)
}

func InitMySQL() *sql.DB {
	dsn := getDSN()

	fmt.Printf("Attempting MySQL connection with DSN: %s\n", dsn)

	dbHandle, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Warn().Msg("Failed to open MySQL connection: " + err.Error())
		return nil
	}
	defer dbHandle.Close()

	if err := dbHandle.Ping(); err != nil {
		log.Warn().Msg("Can't reach MySQL: " + err.Error())
	}

	// Critical production settings
	dbHandle.SetMaxOpenConns(25)
	dbHandle.SetMaxIdleConns(25)
	dbHandle.SetConnMaxLifetime(5 * time.Minute)
	dbHandle.SetConnMaxIdleTime(10 * time.Minute)

	// Retry connection
	for i := range 10 {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		if err = dbHandle.PingContext(ctx); err == nil {
			cancel()
			log.Info().Msg("MySQL connected successfully!")
			return dbHandle
		}
		cancel()
		log.Printf("MySQL not ready, retry %d/10...\n", i+1)
		log.Printf("Error: %s\n", err.Error())
		time.Sleep(time.Duration(i+1) * time.Second)
	}

	log.Warn().Msg("Could not connect to MySQL after retries")
	return nil
}

func Close(dbHandle *sql.DB) {
	if dbHandle != nil {
		dbHandle.Close()
	}
}
