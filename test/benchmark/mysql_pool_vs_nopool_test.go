package benchmark

import (
	"fmt"
	"log"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbWithPool, dbWithoutPool *gorm.DB

// Initialize MySQL connection with connection pooling enabled
func initDBWithPool() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/boxfetch?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	dbWithPool, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database with pool: %v", err)
	}
}

// Initialize MySQL connection without connection pooling
func initDBWithoutPool() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/boxfetch?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	dbWithoutPool, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database without pool: %v", err)
	}
}

// Benchmark inserting a single record with connection pool enabled
func BenchmarkInsertOneRecordWithPool(b *testing.B) {
	initDBWithPool()

	// Drop the table if it exists, then recreate the table
	resetUserTable(dbWithPool)

	sqlDB, err := dbWithPool.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB from gorm.DB (with pool): %v", err)
	}

	// Set the connection pool to 10
	sqlDB.SetMaxOpenConns(10)
	defer sqlDB.Close()

	b.ResetTimer() // Reset the timer before the benchmark starts

	for i := 0; i < b.N; i++ {
		insertRecord(b, dbWithPool, i)
	}
}

// Benchmark inserting a single record with connection pooling disabled (1 connection)
func BenchmarkInsertOneRecordWithoutPool(b *testing.B) {
	initDBWithoutPool()

	// Drop the table if it exists, then recreate the table
	resetUserTable(dbWithoutPool)

	sqlDB, err := dbWithoutPool.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB from gorm.DB (without pool): %v", err)
	}

	// Set the connection pool to 1 (disable pooling effectively)
	sqlDB.SetMaxOpenConns(1)
	defer sqlDB.Close()

	b.ResetTimer() // Reset the timer before the benchmark starts

	for i := 0; i < b.N; i++ {
		insertRecord(b, dbWithoutPool, i)
	}
}

// Insert record helper function (shared between both tests)
func insertRecord(b *testing.B, db *gorm.DB, i int) {
	user := User{
		Username: fmt.Sprintf("testuser_%d_%d", i, time.Now().UnixNano()), // Ensure a unique username
		Email:    fmt.Sprintf("testuser_%d@example.com", i),               // Ensure a unique email
		Active:   true,
	}
	if err := db.Create(&user).Error; err != nil {
		b.Fatalf("failed to insert record: %v", err)
	}
}

// Reset the user table for each test
func resetUserTable(db *gorm.DB) {
	if db.Migrator().HasTable(&User{}) {
		if err := db.Migrator().DropTable(&User{}); err != nil {
			log.Println("Error dropping table:", err)
		}
	}
	db.AutoMigrate(&User{})
}

// User struct
type User struct {
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"type:varchar(100);unique"`
	Email    string `gorm:"type:varchar(100);unique"`
	Active   bool
}
