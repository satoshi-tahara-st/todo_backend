package application

import (
	"io"
	"os"

	"github.com/go-yaml/yaml"
)

//go:generate enumer -transform snake -json -yaml -type=Stage -trimprefix Stage
type Stage int

const (
	//ローカル
	StageLocal Stage = iota
	//検証環境
	StageStaging
	//本番環境
	StageProduction
)

// Database DBの設定
type Database struct {
	Driver string
	DBName string
	Ref    ConnectionSetting `yaml:"ref"`
	// Upd    ConnectionSetting `yaml:"upd"`
}

type ConnectionSetting struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`

	MaxOpenConns           int `yaml:"max_open_conns"`
	MaxIdleConns           int `yaml:"max_idle_conns"`
	ConnMaxLifetimeSeconds int `yaml:"conn_max_lifetime_seconds"`
}

type Config struct {
	Stage Stage               `yaml:"stage"`
	DB    map[string]Database `yaml:"db"`
}

// yamlFileを読み込んでConfigに変換します
func NewConfig(yamlFile string) (*Config, error) {
	file, err := os.Open(yamlFile)
	if err != nil {
		return nil, err
	}

	return newConfig(file)
}

func newConfig(yamlFile io.Reader) (*Config, error) {
	buf, err := io.ReadAll(yamlFile)
	if err != nil {
		return nil, err
	}

	var conf Config

	err = yaml.Unmarshal(buf, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}

func (c *Config) isLocal() bool {
	return c.Stage == StageLocal
}

func (c *Config) isDev() bool {
	return c.Stage == StageLocal || c.Stage == StageStaging
}
