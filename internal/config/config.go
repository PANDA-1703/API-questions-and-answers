package config

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Mode       string
		HttpServer *HttpServerConfig
		Handler    *HandlerConfig
		Service    *ServiceConfig
		Postgres   *PostgresConfig
	}

	PostgresConfig struct {
		Host     string
		User     string
		Password string
		DBName   string
		Port     int
	}

	HttpServerConfig struct {
		Port           int
		ReadTimeout    time.Duration
		WriteTimeout   time.Duration
		MaxHeaderBytes int
	}

	HandlerConfig struct {
		RequestTimeout time.Duration
		StreamTimeout  time.Duration
	}

	ServiceConfig struct{}
)

func Init(configPath string, testMode bool) (*Config, error) {
	jsonCfg := viper.New()
	jsonCfg.AddConfigPath(filepath.Dir(configPath))
	jsonCfg.SetConfigName(filepath.Base(configPath))

	if err := jsonCfg.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("config/Init/jsonCfg.ReadInCinfig: %w", err)
	}

	envCfg := viper.New()
	envCfg.SetConfigFile(".env")

	if err := envCfg.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("config/Init/envCfg.ReadInCinfig: %w", err)
	}

	cfg := &Config{
		Mode: jsonCfg.GetString("mode"),
		HttpServer: &HttpServerConfig{
			Port:           jsonCfg.GetInt("httpServer.port"),
			ReadTimeout:    jsonCfg.GetDuration("httpServer.readTimeout"),
			WriteTimeout:   jsonCfg.GetDuration("httpServer.writeTimeout"),
			MaxHeaderBytes: jsonCfg.GetInt("httpServer.maxHeaderBytes"),
		},
		Handler: &HandlerConfig{
			RequestTimeout: jsonCfg.GetDuration("handler.requestTimeout"),
			StreamTimeout:  jsonCfg.GetDuration("handler.streamTimeout"),
		},
		Service: &ServiceConfig{},
		Postgres: &PostgresConfig{
			Host:     envCfg.GetString("POSTGRES_HOST"),
			User:     envCfg.GetString("POSTGRES_USER"),
			Password: envCfg.GetString("POSTGRES_PASSWORD"),
			DBName:   envCfg.GetString("POSTGRES_DB"),
			Port:     envCfg.GetInt("POSTGRES_PORT"),
		},
	}
	return cfg, nil
}

func (p *PostgresConfig) PgSource() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable pool_max_conns=32",
		p.Host, p.Port, p.User, p.Password, p.DBName)
}
