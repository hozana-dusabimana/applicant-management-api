package database

import (
	"fmt"
	"job-tracker/models"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB //CONNECTION POINTER

func ConnectDB() {
	// Get database configuration from environment variables
	host := getEnv("DB_HOST", "localhost")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "123")
	dbname := getEnv("DB_NAME", "postgres")
	port := getEnv("DB_PORT", "5432")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		host, user, password, dbname, port)

	// Configure GORM with better settings
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Configure connection pool
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatal("Failed to get database instance: ", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Auto-migrate the schema
	err = database.AutoMigrate(&models.Applicant{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	// Create indexes for better performance
	database.Exec("CREATE INDEX IF NOT EXISTS idx_applicants_email ON applicants(email)")
	database.Exec("CREATE INDEX IF NOT EXISTS idx_applicants_status ON applicants(status)")
	database.Exec("CREATE INDEX IF NOT EXISTS idx_applicants_created_at ON applicants(created_at)")

	DB = database
	log.Println("Connected to database successfully")
}

// Helper function to get environment variable with default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
