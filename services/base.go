package services

import (
	"fmt"
	"os"
	helper "superapps/helpers"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbDefault *gorm.DB
	// dbPPOB    *gorm.DB
	// dbPayment *gorm.DB
)

// InitDBs initializes all database connections.
func InitDBs() {
	err := godotenv.Load()
	if err != nil {
		helper.Logger("error", "Error loading .env file")
	}

	dbDefault = connectDB(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	if dbDefault == nil {
		panic("❌ dbDefault is nil, failed to connect to default DB")
	}

	// dbPPOB = connectDB(
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_PORT"),
	// 	os.Getenv("DB_PPOB_NAME"),
	// )

	// if dbPPOB == nil {
	// 	panic("❌ dbPPOB is nil, failed to connect to PPOB DB")
	// }

	// dbPayment = connectDB(
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_PORT"),
	// 	os.Getenv("DB_PPOB_PAYMENT"),
	// )

	// if dbPayment == nil {
	// 	panic("❌ dbPayment is nil, failed to connect to Payment DB")
	// }
}

// connectDB connects to a database using GORM v2 and returns *gorm.DB
func connectDB(user, pass, host, port, name string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name,
	)

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		helper.Logger("error", fmt.Sprintf("Failed to connect to DB (%s): %s", name, err.Error()))
		return nil
	}

	// Get the generic DB object
	sqlDB, err := conn.DB()
	if err != nil {
		helper.Logger("error", fmt.Sprintf("Failed to get generic DB instance (%s): %s", name, err.Error()))
		return nil
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	helper.Logger("info", fmt.Sprintf("Connected to database: %s", name))
	return conn
}

// GetDefaultDB returns the default DB instance
func GetDefaultDB() *gorm.DB {
	return dbDefault
}

// GetPPOBDB returns the PPOB DB instance
// func GetPPOBDB() *gorm.DB {
// 	return dbPPOB
// }

// GetPaymentDB returns the Payment DB instance
// func GetPaymentDB() *gorm.DB {
// 	return dbPayment
// }
