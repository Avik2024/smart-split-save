package config

import (
	"log"
	"os"
)

type Config struct {
	Port       string
	DBUrl      string
	JWTSecret  string
	RedisAddr  string
	KafkaBrokers string
}

func Load() *Config {
	c := &Config{
		Port: os.Getenv("PORT"),
		DBUrl: os.Getenv("DATABASE_URL"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		RedisAddr: os.Getenv("REDIS_ADDR"),
		KafkaBrokers: os.Getenv("KAFKA_BROKERS"),
	}
	if c.Port == "" { c.Port = "4000" }
	if c.DBUrl == "" { log.Fatal("DATABASE_URL required") }
	if c.JWTSecret == "" { log.Fatal("JWT_SECRET required") }
	return c
}
