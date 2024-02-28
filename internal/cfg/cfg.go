package cfg

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Build  string `yaml:"build"`
	Server `yaml:"server"`
	Cache  `yaml:"cache"`
}

type Server struct {
	Port string `yaml:"port"`
}

type Cache struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

func loadConfig() {
	var result Config

	cfgPath := os.Getenv("CFG_PATH")
	file, err := os.ReadFile(cfgPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = yaml.Unmarshal(file, &result)
	if err != nil {
		log.Fatal(err.Error())
	}

	instance = &result
}
