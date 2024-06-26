package main

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	//"os"
	//"github.com/BurntSushi/toml"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port int
	Host string
}

type DatabaseConfig struct {
	User     string
	Password string
	DBName   string
}

func loadConfig(file string) (Config, error) {
	var config Config
	if _, err := toml.DecodeFile(file, &config); err != nil {
		return config, err
	}
	return config, nil
}

func main() {
	config, err := loadConfig("config.toml")
	if err != nil {
		log.Fatalf("Failed to load config: %s", err)
	}
	fmt.Printf("Server is running on %s:%d\n", config.Server.Host, config.Server.Port)
	fmt.Printf("Database: %s, User: %s\n", config.Database.DBName, config.Database.User)
}
