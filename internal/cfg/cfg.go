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

func MustLoad(cfgPath string) Config {
	var result Config

	file, err := os.ReadFile(cfgPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = yaml.Unmarshal(file, &result)

	return result
}
