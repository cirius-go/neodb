package gormdb

import (
	"fmt"
	"net/url"
)

// PostgresConfig represents the configuration for Postgres database.
type PostgresConfig struct {
	Host     string `envPrefix:"HOST"`
	Port     int    `envPrefix:"PORT"`
	Username string `envPrefix:"USERNAME"`
	Password string `envPrefix:"PASSWORD"`
	Database string `envPrefix:"NAME"`
	Params   string `envPrefix:"PARAMS"`
}

// BuildPostgresConnStr builds a PostgreSQL connection string from a Config struct.
func BuildPostgresConnStr(cfg PostgresConfig) (string, error) {
	queries, err := url.ParseQuery(cfg.Params)
	if err != nil {
		return "", fmt.Errorf("invalid db params '%s': %w", cfg.Params, err)
	}
	if queries.Get("sslmode") == "" {
		queries.Set("sslmode", "disable")
	}
	if queries.Get("connect_timeout") == "" {
		queries.Set("connect_timeout", "5")
	}
	u := url.URL{
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		User:     url.UserPassword(cfg.Username, cfg.Password),
		Path:     cfg.Database,
		RawQuery: queries.Encode(),
	}
	return u.String(), nil
}

// // NewPostgresDialector and returns dialector if not error.
// func NewPostgresDialector(cfg PostgresConfig) (gorm.Dialector, error) {
// 	dsn, err := BuildPostgresConnStr(cfg)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return postgres.New(postgres.Config{
// 		DSN:                  dsn,
// 		PreferSimpleProtocol: true,
// 	}), nil
// }
