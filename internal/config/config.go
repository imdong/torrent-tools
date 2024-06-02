package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config 基础信息
	Config struct {
		Http `yaml:"http"`
		Log  `yaml:"logger"`
		DB   `yaml:"db"`
	}

	// Http 服务相关
	Http struct {
		Port string `env-default:"3682" yaml:"web_port" env:"WEB_PORT"`
	}

	// Log 记录日志
	Log struct {
		Level string `env-default:"warn" yaml:"log_level" env:"LOG_LEVEL"`
	}

	DB struct {
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	// 如果文件不存在就完全使用默认配置
	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
