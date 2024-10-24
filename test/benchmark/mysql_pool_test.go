package benchmark

// import (
// 	"fmt"
// 	"log"
// 	"testing"
// 	"time"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var db *gorm.DB

// // Initialize MySQL connection
// func initDB() {
// 	dsn := "root:123456@tcp(127.0.0.1:3306)/boxfetch?charset=utf8mb4&parseTime=True&loc=Local"
// 	var err error
// 	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("failed to connect database: %v", err)
// 	}
// }

// // Benchmark inserting a single record with 1 open connection
// func BenchmarkInsertOneRecordMaxOpenConns1(b *testing.B) {
// 	initDB()

// 	// Drop the table if it exists, then create the table
// 	resetUserTable()

// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		log.Fatalf("failed to get sql.DB from gorm.DB: %v", err)
// 	}

// 	// Set the connection pool to 1
// 	sqlDB.SetMaxOpenConns(1)
// 	defer sqlDB.Close()

// 	b.ResetTimer() // Reset the timer before the benchmark starts

// 	// Insert one record in the benchmark
// 	for i := 0; i < b.N; i++ {
// 		insertRecord(b, db, i)
// 	}
// }

// // Benchmark inserting a single record with 10 open connections
// func BenchmarkInsertOneRecordMaxOpenConns10(b *testing.B) {
// 	initDB()

// 	// Drop the table if it exists, then create the table
// 	resetUserTable()

// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		log.Fatalf("failed to get sql.DB from gorm.DB: %v", err)
// 	}

// 	// Set the connection pool to 10
// 	sqlDB.SetMaxOpenConns(10)
// 	defer sqlDB.Close()

// 	b.ResetTimer() // Reset the timer before the benchmark starts

// 	// Insert one record in the benchmark
// 	for i := 0; i < b.N; i++ {
// 		insertRecord(b, db, i)
// 	}
// }

// // Helper function to insert one record
// func insertRecord(b *testing.B, db *gorm.DB, i int) {
// 	user := User{
// 		Username: fmt.Sprintf("testuser_%d_%d", i, time.Now().UnixNano()), // Ensure a unique username
// 		Email:    fmt.Sprintf("testuser_%d@example.com", i),               // Ensure a unique email
// 		Active:   true,
// 	}
// 	if err := db.Create(&user).Error; err != nil {
// 		b.Fatalf("failed to insert record: %v", err)
// 	}
// }

// // Helper function to drop and recreate the user table
// func resetUserTable() {
// 	if db.Migrator().HasTable(&User{}) {
// 		if err := db.Migrator().DropTable(&User{}); err != nil {
// 			log.Println("Error dropping table:", err)
// 		}
// 	}
// 	db.AutoMigrate(&User{})
// }

// // User struct
// type User struct {
// 	ID       int    `gorm:"primaryKey;autoIncrement"`
// 	Username string `gorm:"type:varchar(100);unique"`
// 	Email    string `gorm:"type:varchar(100);unique"`
// 	Active   bool
// }
