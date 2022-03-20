package postgresdb

import (
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

type DBFilter struct {
	FilterName  string
	FilterValue string
}

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	sqlDB, err := sql.Open("postgres", os.Getenv("DB_URI"))
	if err != nil {
		panic("Error opening database connection.")
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = gormDB
}

func IsUniqeConstraintError(err error, constraint string) bool {
	if pqErr, ok := err.(*pq.Error); ok {
		return pqErr.Code == "23505" && pqErr.Constraint == constraint
	}

	return false
}
