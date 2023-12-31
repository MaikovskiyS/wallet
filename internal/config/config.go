package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Psql       Postgres   `yaml:"postgres"`
	HttpServer HttpServer `yaml:"server"`
}

// Constructor parse yaml file and decode to struct
func New() (*Config, error) {
	f, err := os.Open("./cmd/config.yml")
	if err != nil {
		return &Config{}, fmt.Errorf("open file err: %w", err)
	}

	c := &Config{}

	err = yaml.NewDecoder(f).Decode(&c)
	if err != nil {
		return &Config{}, fmt.Errorf("decode cfg err: %w", err)
	}
	return c, nil
}

type Postgres struct {
	DbDriver      string
	Host          string
	Port          string
	UserName      string
	Password      string
	Schema        string
	SslMode       string
	MigrationPath string
}

func (p *Postgres) ConnString() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s", p.DbDriver, p.UserName, p.Password, p.Host, p.Port, p.Schema, p.SslMode)
}

func (p *Postgres) Driver() string {
	return p.DbDriver
}

type HttpServer struct {
	Port         string
	Writetimeout time.Duration
	Readtimeout  time.Duration
}

func (s *HttpServer) WriteTimeout() time.Duration {
	return s.Writetimeout * time.Second
}
func (s *HttpServer) ReadTimeout() time.Duration {
	return s.Readtimeout * time.Second
}
