package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// User entry
type User struct {
	Password string `json:"password"`
	Role     string `json:"role"`
}

// Config of manager
type Config struct {
	filePath string

	Port  int `json:"port,omitempty"`
	Users map[string]User
}

// Default of manager
var Default = Config{
	Port: 3000,
}

var tag = "CONFIG"

// Dump config info
func (c Config) Dump() {
	var output []string

	output = append(output, fmt.Sprintf("loaded configuration from \"%v\"", c.filePath))
	output = append(output, fmt.Sprintf("web port: %v", c.Port))

	for _, line := range output {
		log.Printf("[%v] %v", tag, line)
	}
}

// Load config from file
func Load(configPath string) (*Config, error) {
	buf, err := ioutil.ReadFile(configPath)

	if err != nil {
		return nil, err
	}

	var c Config = Default

	c.filePath = configPath

	err = json.Unmarshal(buf, &c)

	return &c, nil
}
