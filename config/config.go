package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	DownMessage  string `yaml:"down_message"`
	WebhookUrl   string `yaml:"webhook_url"`
	Port         string `yaml:"port"`
	CrashMessage string `yaml:"crash_message"`
	StoppageTime int    `yaml:"stoppage_time"`
	Mongo        Mongo  `yaml:"mongo"`
}
type Mongo struct {
	Uri         string   `yaml:"uri"`
	Database    string   `yaml:"database"`
	Collections []string `yaml:"collections"`
}

func Init(configPath string) *Config {
	// Create config structure
	c := &Config{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	//defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&c); err != nil {
		panic(err)
	}

	return c
}
