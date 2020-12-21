package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/kelseyhightower/envconfig"

	"github.com/playmean/scoper/logger"
)

// User entry
type User struct {
	Password string `json:"password"`
	Role     string `json:"role"`
}

// Database config
type Database struct {
	Host     string `json:"host" envconfig:"DATABASE_HOST"`
	User     string `json:"user" envconfig:"DATABASE_USER"`
	Password string `json:"password" envconfig:"DATABASE_PASS"`
	DBName   string `json:"dbname" envconfig:"DATABASE_NAME"`
	Port     int    `json:"port" envconfig:"DATABASE_PORT"`
}

// Config of server
type Config struct {
	Address  string   `json:"address" envconfig:"SERVER_ADDRESS"`
	Port     int      `json:"port" envconfig:"SERVER_PORT"`
	Password string   `json:"password"`
	Database Database `json:"database"`
}

// Default config
var Default = Config{
	Address:  "",
	Port:     8080,
	Password: "password",
	Database: Database{
		Host:     "localhost",
		User:     "postgres",
		Password: "",
		DBName:   "scoper",
		Port:     5432,
	},
}

// SuperUsers predefined list
var SuperUsers = make(map[string]string)

var tag = "CONFIG"

// Dump config info
func (c Config) Dump() {
	var output []string

	output = append(output, fmt.Sprintf("HTTP listen at %s:%d", c.Address, c.Port))
	output = append(output, fmt.Sprintf("database from %s:%d", c.Database.Host, c.Database.Port))

	for _, line := range output {
		logger.Log(tag, line)
	}
}

func readFileConfig(configPath string, c *Config) error {
	buf, err := ioutil.ReadFile(configPath)

	if err != nil {
		return err
	}

	return json.Unmarshal(buf, &c)
}

// Load config from file
func Load(configPath string) (*Config, error) {
	var err error
	var c Config = Default

	err = readFileConfig(configPath, &c)

	if err == nil {
		logger.Log(tag, "loaded configuration from \"%v\"", configPath)
	}

	err = envconfig.Process("", &c)

	SuperUsers["super"] = c.Password

	return &c, err
}
