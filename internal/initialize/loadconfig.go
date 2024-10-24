package initialize

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"github.com/tienphuckx/ecom-backbone-api.git/global"
)

// LoadConfig loads the application configuration from a file
func LoadConfig() {
	// Create a new viper instance
	viper := viper.New()

	// Set the config file location and type
	viper.AddConfigPath("./config/") // Path to config directory
	viper.SetConfigName("dev")       // Config file name (without extension)
	viper.SetConfigType("yaml")      // Config file type

	// Attempt to read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Unmarshal SysConfig into the global SysConfig variable
	if err := viper.Unmarshal(&global.SysConfig); err != nil {
		log.Fatalf("Error unmarshalling into SysConfig: %v", err)
	}

	// Unmarshal ApplicationConfig into the global ApplicationConfig variable
	if err := viper.Unmarshal(&global.ApplicationConfig); err != nil {
		log.Fatalf("Error unmarshalling into ApplicationConfig: %v", err)
	}

	// Debugging output: print some loaded values to ensure they are loaded correctly
	fmt.Println("MySQL Host:", global.SysConfig.MySQLConfig.Host)
	fmt.Println("Server Port:", global.ApplicationConfig.Server.Port)
	fmt.Println("JWT Key:", global.ApplicationConfig.Security.JWT.Key)
}
