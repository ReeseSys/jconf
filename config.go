package jconf

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	file   string
	config map[string]interface{}
}

// NewConfig is used to parse configuration values from a JSON file.
func NewConfig(filename string) (*Config, error) {
	config := &Config{
		file: filename,
	}
	err := config.readConfig()

	return config, err
}

// Internal function to read config JSON.
func (c *Config) readConfig() error {
	file, err := os.Open(c.file)
	if nil != err {
		return fmt.Errorf("Could not read config %s: %s", c.file, err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c.config)

	if nil != err {
		return fmt.Errorf("Could not parse config %s: %s", c.file, err)
	}

	return nil
}

// The HasConfig function can be used to determine if a config value
// exists before reading.
func (c *Config) HasConfig(key string) bool {
	_, ok := c.config[key]
	return ok
}

// The GetConfig function is used to get a config value as an interface.
// nil is returned if the config value does not exist.
func (c *Config) GetConfig(key string) interface{} {
	val, _ := c.config[key]

	return val
}

// The GetConfigStr function is used to get a config value as a string.
// An empty string is returned if the config value does not exist.
func (c *Config) GetConfigStr(key string) string {
	val, ok := c.config[key]

	if ok {
		return val.(string)
	}

	return ""
}

// The GetConfigStr function is used to get a config as an integter.
// -1 is returned if the config value does not exist.
func (c *Config) GetConfigInt(key string) int {
	val, ok := c.config[key]

	if ok {
		return val.(int)
	}

	return -1
}
