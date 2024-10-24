package initialize

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // MySQL driver import
	"github.com/tienphuckx/ecom-backbone-api.git/global"
	"go.uber.org/zap"
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

	// Print the DSN for debugging
	fmt.Println("DSN:", dsn)

	// Open the connection to MySQL
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open MySQL connection: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping MySQL: %w", err)
	}

	global.Logger.Info("MySQL successfully connected", zap.String("host", global.SysConfig.MySQLConfig.Host))
	return nil
}
