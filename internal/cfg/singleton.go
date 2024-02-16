package cfg

import "sync"

var instance *Config
var once sync.Once

func GetInstance() *Config {
	once.Do(loadConfig)
	return instance
}
