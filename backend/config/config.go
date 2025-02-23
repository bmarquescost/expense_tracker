package config

import "os"

type Config struct {
    ServerAddr string
}

func LoadConfig() *Config {
    return &Config{
        ServerAddr: getEnv("SERVER_ADDR", ":8080"),
    }
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}