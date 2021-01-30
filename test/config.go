package test

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Config struct for webapp config
type Config struct {
	Server struct {
		Scheme string `yaml:"scheme"`
		// Host is the local machine IP Address to bind the HTTP Server to
		Host string `yaml:"host"`

		// Port is the local machine TCP Port to bind the HTTP Server to
		Port int `yaml:"port"`
	} `yaml:"server"`
}

//LoadConfig loads the local config
func LoadConfig() (*Config, error) {
	return NewConfig("./config.yml")
}

// NewConfig returns a new decoded Config struct
func NewConfig(configPath string) (*Config, error) {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

//GetURL builds the url
func (c *Config) GetURL() string {
	return fmt.Sprintf("%s://%s:%d", c.Server.Scheme, c.Server.Host, c.Server.Port)
}
