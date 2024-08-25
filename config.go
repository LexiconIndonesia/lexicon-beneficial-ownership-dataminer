package main

import (
	"fmt"
	"os"
	"strconv"
)

func loadEnvString(key string, result *string) {
	s, ok := os.LookupEnv(key)

	if !ok {
		return
	}
	*result = s
}

func loadEnvUint(key string, result *uint) {
	s, ok := os.LookupEnv(key)

	if !ok {
		return
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return
	}
	*result = uint(n)
}

/* Configuration */

/* PgSQL Configuration */
type pgSqlConfig struct {
	Host     string `json:"host"`
	Port     uint   `json:"port"`
	Database string `json:"database"`
	SslMode  string `json:"ssl_mode"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func (p pgSqlConfig) ConnStr() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s database=%s sslmode=%s", p.Host, p.Port, p.User, p.Password, p.Database, p.SslMode)
}

func defaultPgSql() pgSqlConfig {
	return pgSqlConfig{
		Host:     "localhost",
		Port:     5432,
		Database: "database",
		User:     "",
		Password: "",
		SslMode:  "disable",
	}
}

func (p *pgSqlConfig) loadFromEnv(dbName string) {
	loadEnvString("POSTGRES_HOST", &p.Host)
	loadEnvUint("POSTGRES_PORT", &p.Port)
	loadEnvString(dbName, &p.Database)
	loadEnvString("POSTGRES_SSLMODE", &p.SslMode)
	loadEnvString("POSTGRES_USERNAME", &p.User)
	loadEnvString("POSTGRES_PASSWORD", &p.Password)
}

/* Listen Configuration */

type listenConfig struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}

func (l listenConfig) Addr() string {
	return fmt.Sprintf("%s:%d", l.Host, l.Port)
}

func defaultListenConfig() listenConfig {
	return listenConfig{
		Host: "127.0.0.1",
		Port: 8080,
	}
}

func (l *listenConfig) loadFromEnv() {
	loadEnvString("LISTEN_HOST", &l.Host)
	loadEnvUint("LISTEN_PORT", &l.Port)
}

type config struct {
	Listen        listenConfig `json:"listen"`
	CrawlerDB     pgSqlConfig  `json:"crawler_db"`
	BeneficialDB  pgSqlConfig  `json:"beneficial_db"`
	BackendApiKey string       `json:"api_key"`
	ServerSalt    string       `json:"salt"`
}

func (c *config) loadFromEnv() {
	c.Listen.loadFromEnv()
	c.CrawlerDB.loadFromEnv("POSTGRES_CRAWLER_DB_NAME")
	c.BeneficialDB.loadFromEnv("POSTGRES_BENEFICIAL_OWNERSHIP_DB_NAME")

}

func defaultConfig() config {
	return config{
		Listen:       defaultListenConfig(),
		CrawlerDB:    defaultPgSql(),
		BeneficialDB: defaultPgSql(),
	}
}
