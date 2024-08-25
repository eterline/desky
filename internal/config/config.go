package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type (
	Settings struct {
		Address struct {
			Ip   string `yaml:"ip"`
			Port string `yaml:"port"`
		} `yaml:"address"`
		User struct {
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		} `yaml:"user"`
		Tls struct {
			Crt string `yaml:"crt"`
			Key string `yaml:"key"`
		} `yaml:"tls"`
		SessionStoreKey string `yaml:"sessionkey"`
	}
)

func ParseSettings() Settings {
	file, err := os.Open("settings.yml")
	if err != nil {
		log.Fatal(err.Error())
	}

	var cfg Settings
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	return cfg
}
