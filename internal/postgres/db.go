package postgres

import (
	"fmt"
	"os"
	"time"
)

type (

	Spent struct {
		Id 			int 		`json:"id"`
		Name 		string 		`json:"name"`
		Amount 		float32 	`json:"amount"`
		CreatedAt 	time.Time 	`json:"crated_at"`
		UpdatedAt 	time.Time 	`json:"updated_at"`
	}
)

var (
	PostgresHost           = getEnv("POSTGRES_HOST", "localhost")
	PostgresPort           = getEnv("POSTGRES_PORT", "5432")
	PostgresDB             = getEnv("POSTGRES_DB", "list_expense_development")
	PostgresDBTest         = getEnv("POSTGRES_DB_TEST", "list_expense_test")
	PostgresUser           = getEnv("POSTGRES_USER", "ivan")
	PostgresPassword       = getEnv("POSTGRES_PASSWORD", "Kup0lA")
	PostgresConnectTimeout = getEnv("POSTGRES_CONNECT_TIMEOUT", "3")

	PostgresSys = fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s connect_timeout=%s sslmode=disable",
		PostgresUser, PostgresPassword, PostgresHost, PostgresPort, PostgresDB, PostgresConnectTimeout)

	PostgresTest = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s connect_timeout=%s sslmode=disable",
		PostgresUser, PostgresPassword, PostgresHost, PostgresPort, PostgresDBTest, PostgresConnectTimeout)

)

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		value = fallback
	}

	return value
}

