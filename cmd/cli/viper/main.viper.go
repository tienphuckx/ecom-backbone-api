package main

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config struct to map the YAML file structure
type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	// Adjusted to reflect a list of databases
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		DbName   string `mapstructure:"dbName"`
	} `mapstructure:"databases"`

	Security struct {
		JWT struct {
			Key string `mapstructure:"key"`
		} `mapstructure:"jwt"`
	} `mapstructure:"security"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config/") // Path to config file directory
	viper.SetConfigName("dev")       // Config file name without extension
	viper.SetConfigType("yaml")      // Specify YAML as the config type

	// Read configuration file
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read configuration: %w", err))
	}

	// Output some specific values (for debugging)
	fmt.Println("Server Port:", viper.GetInt("server.port"))
	fmt.Println("JWT Key:", viper.GetString("security.jwt.key"))

	// Map configuration into the Config struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("Failed to unmarshal config: %w", err))
	}

	// Accessing server port
	fmt.Printf("Server Port from struct: %d\n", config.Server.Port)

	// Accessing the list of databases
	for i, db := range config.Databases {
		fmt.Printf("Database %d - User: %s, Host: %s, DB Name: %s\n", i+1, db.User, db.Host, db.DbName)
	}

	// Accessing JWT key
	fmt.Printf("JWT Key: %s\n", config.Security.JWT.Key)
}
