package initialize

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // MySQL driver import
	"github.com/tienphuckx/ecom-backbone-api.git/global"
	"github.com/tienphuckx/ecom-backbone-api.git/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *sql.DB

func InitMySql() error {
	fmt.Println("MySQL User:", global.SysConfig.MySQLConfig.User)
	fmt.Println("MySQL Password:", global.SysConfig.MySQLConfig.Password)
	fmt.Println("MySQL Host:", global.SysConfig.MySQLConfig.Host)
	fmt.Println("MySQL Port:", global.SysConfig.MySQLConfig.Port)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.SysConfig.MySQLConfig.User,
		global.SysConfig.MySQLConfig.Password,
		global.SysConfig.MySQLConfig.Host,
		global.SysConfig.MySQLConfig.Port,
		global.SysConfig.MySQLConfig.DbName,
	)

	// Open the connection to MySQL using GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to open MySQL connection: %w", err)
	}

	// Set the global GORM DB connection
	global.Mdb = db // Assign the *gorm.DB object to global.Mdb

	SetPool()
	MigrateTables()

	global.Logger.Info("MySQL successfully connected and migrated", zap.String("host", global.SysConfig.MySQLConfig.Host))
	return nil
}

// SetPool sets the connection pool configurations for the MySQL connection
func SetPool() {
	// Retrieve the MySQL config from the global configuration
	m := global.SysConfig.MySQLConfig

	// Ensure that the database connection is already established
	if global.Mdb == nil {
		fmt.Println("MySQL connection is not initialized")
		return
	}

	// Get the underlying *sql.DB object from *gorm.DB
	sqlDB, err := global.Mdb.DB()
	if err != nil {
		global.Logger.Error("Failed to get underlying sql.DB", zap.Error(err))
		return
	}

	// Set the connection pool configurations on *sql.DB
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime) * time.Second)

	// Optionally, log success
	fmt.Println("MySQL connection pool configured successfully")
}

// MigrateTables runs the database migrations for User and Role tables
func MigrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)

	if err != nil {
		global.Logger.Error("Migration failed", zap.Error(err))
		return
	}

	global.Logger.Info("Migration completed successfully!")
}

func CloseMySql() {
	if db != nil {
		if err := db.Close(); err != nil {
			global.Logger.Error("Failed to close MySQL connection", zap.Error(err))
		} else {
			global.Logger.Info("MySQL connection closed successfully")
		}
	}
}
