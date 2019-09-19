package server

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// Config load and hold config
type Config struct{}

// NewConfig create Config entity
func NewConfig() *Config {
	c := &Config{}
	file, err := os.Open(".env")
	if err != nil {
		log.Printf("[Warn] load .env error (if you have setting the env, ignore this) %v\n", err)
	}
	defer file.Close()
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.Trim(line, "\n")
		line = strings.Trim(line, " ")
		if len(line) > 0 {
			if line[0] == '#' {
				continue
			}
			if strings.Contains(line, "=") {
				kvSet := strings.SplitN(line, "=", 2)
				key := strings.Trim(kvSet[0], " ")
				value := strings.Trim(kvSet[1], " ")
				log.Printf("[Info] set env %s=%s\n", key, value)
				os.Setenv(key, value)
			}
		}
		if err != nil {
			break
		}
	}
	return c
}

// Set set env kv
func (c *Config) Set(key, value string) {
	os.Setenv(key, value)
}

// Get get env value by key
func (c *Config) Get(key string) string {
	return os.Getenv(key)
}

// Int get int env value
func (c *Config) Int(key string) int {
	value := c.Get(key)
	ret, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("[Error] can not load %s=%s as int\n", key, value)
		panic(err)
	}
	return ret
}
