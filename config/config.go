package config

import (
	"github.com/spf13/viper"
	"log"
)

var (
	LogLevel       = "info"
	LogEncoding    = "console"
	LogFile        = false
	LogFileMaxSize = 500
	LogFilePath    = "/tmp/"
	HttpPort       = "8080"
	DatabaseURL    = "postgres://postgres:postgres@localhost:5432/openapi?sslmode=disable"
	RedisHost      = "localhost"
	RedisPort      = "6379"
	RedisPoolSize  = 10
	JwtSecretKey   = "secret"
	SessionKey     = "secret-key"
)

func LoadConfig() error {
	viper.SetConfigType("json")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("open-api-client")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.open-api-client")
	viper.AddConfigPath("/etc/open-api-client")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("/app")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		return err
	}

	setDefaults()

	LogLevel = viper.GetString("LOG_LEVEL")
	LogEncoding = viper.GetString("LOG_ENCODING")
	LogFile = viper.GetBool("LOG_FILE")
	LogFileMaxSize = viper.GetInt("LOG_FILE_MAX_SIZE")
	LogFilePath = viper.GetString("LOG_FILE_PATH")
	HttpPort = viper.GetString("HTTP_PORT")
	DatabaseURL = viper.GetString("DATABASE_URL")
	RedisHost = viper.GetString("REDIS_HOST")
	RedisPort = viper.GetString("REDIS_PORT")
	RedisPoolSize = viper.GetInt("REDIS_POOL_SIZE")
	JwtSecretKey = viper.GetString("JWT_SECRET_KEY")
	SessionKey = viper.GetString("SESSION_KEY")

	return nil
}

func setDefaults() {
	viper.SetDefault("LOG_LEVEL", LogLevel)
	viper.SetDefault("LOG_ENCODING", LogEncoding)
	viper.SetDefault("LOG_FILE", LogFile)
	viper.SetDefault("LOG_FILE_MAX_SIZE", LogFileMaxSize)
	viper.SetDefault("LOG_FILE_PATH", LogFilePath)
	viper.SetDefault("HTTP_PORT", HttpPort)
	viper.SetDefault("DATABASE_URL", DatabaseURL)
	viper.SetDefault("REDIS_HOST", RedisHost)
	viper.SetDefault("REDIS_PORT", RedisPort)
	viper.SetDefault("JWT_SECRET_KEY", RedisPoolSize)
	viper.SetDefault("SESSION_KEY", SessionKey)

}
