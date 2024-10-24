package setting

type SysConfig struct {
	LogConfig   LogConfig   `mapstructure:"log"`
	MySQLConfig MySQLConfig `mapstructure:"mysql"`
	RedisConfig RedisConfig `mapstructure:"redis"`
	KafkaConfig KafkaConfig `mapstructure:"kafka"`
	// other configurations...
}

// RedisConfig represents the configuration settings for Redis
type RedisConfig struct {
	Addr         string `mapstructure:"addr"`           // Redis server address
	Password     string `mapstructure:"password"`       // Password for Redis
	DB           int    `mapstructure:"db"`             // Redis DB index
	PoolSize     int    `mapstructure:"pool_size"`      // Connection pool size
	MinIdleConns int    `mapstructure:"min_idle_conns"` // Minimum idle connections in the pool
	DialTimeout  int    `mapstructure:"dial_timeout"`   // Dial timeout in seconds
	ReadTimeout  int    `mapstructure:"read_timeout"`   // Read timeout in seconds
	WriteTimeout int    `mapstructure:"write_timeout"`  // Write timeout in seconds
}

type MySQLConfig struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	DbName          string `mapstructure:"db_name"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"`
}

type LogConfig struct {
	LogLevel    string `mapstructure:"log_level"`
	FileLogName string `mapstructure:"file_log_name"`
	MaxSize     int    `mapstructure:"max_size"`
	MaxBackups  int    `mapstructure:"max_backups"`
	MaxAge      int    `mapstructure:"max_age"`
	Compress    bool   `mapstructure:"compress"`
}

type KafkaConfig struct {
	Brokers  []string `mapstructure:"brokers"`
	ClientID string   `mapstructure:"client_id"`
	Topic    string   `mapstructure:"topic"`
	Acks     string   `mapstructure:"acks"`
}

type ApplicationConfig struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	Security struct {
		JWT struct {
			Key string `mapstructure:"key"`
		} `mapstructure:"jwt"`
	} `mapstructure:"security"`

	Databases []struct {
		User   string `mapstructure:"user"`
		Host   string `mapstructure:"host"`
		DbName string `mapstructure:"db_name"`
	} `mapstructure:"databases"`

	MySQLConfig struct {
		Host            string `mapstructure:"host"`
		Port            int    `mapstructure:"port"`
		User            string `mapstructure:"user"`
		Password        string `mapstructure:"password"`
		DbName          string `mapstructure:"db_name"`
		MaxOpenConns    int    `mapstructure:"max_open_conns"`
		MaxIdleConns    int    `mapstructure:"max_idle_conns"`
		ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"`
	} `mapstructure:"mysql"`
}
