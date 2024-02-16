package cfg

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Build   string `yaml:"build"`
	Storage `yaml:"storage"`
	Server  `yaml:"server"`
}

type Storage struct {
	ConnString string `yaml:"connection_string"`
}

type Server struct {
	Port string `yaml:"port"`
}

func loadConfig() {
	var result Config

	cfgPath := os.Getenv("CFG_PATH")
	file, err := os.ReadFile(cfgPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = yaml.Unmarshal(file, &result)

	instance = &result
}
