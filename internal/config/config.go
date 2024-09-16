package config

import (
	"fmt"
	"log/slog"
	"tender-service/internal/repository/repo"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ServerAddress string        `env:"SERVER_ADDRESS" env-default:"0.0.0.0:8080"`
	StorageType   repo.Type     `env:"STORAGE_TYPE" env-default:"psql"`
	LogLevel      slog.Level    `env:"LOG_LEVEL" env-default:"INFO"`
	APIVersion    string        `env:"API_VERSION" env-default:"v1"`
	Timeout       time.Duration `env:"TIMEOUT" env-default:"10s"`
	Ticker        time.Duration `env:"TICKER" env-default:"60s"`

	PostgresUsername string `env:"POSTGRES_USERNAME"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresHost     string `env:"POSTGRES_HOST"`
	PostgresPort     int    `env:"POSTGRES_PORT" env-default:"5432"`
	PostgresDatabase string `env:"POSTGRES_DB"`
	PostgresJDBCURL  string `env:"POSTGRES_JDBC_URL"`
}

func MustLoad() (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}

	cfg.PostgresJDBCURL = fmt.Sprintf("jdbc:postgresql://%s:%d/%s",
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDatabase)
	return &cfg, nil
}

func (c *Config) DBConnectionString() string {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		c.PostgresUsername, c.PostgresPassword, c.PostgresHost, c.PostgresPort, c.PostgresDatabase)

	return connectionString
}
