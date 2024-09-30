package config

import (
	"log"
	"math/rand"
	"os"

	"gopkg.in/yaml.v3"
)

var printDesky = `
===============================
______             _           
|  _  \           | |          
| | | |  ___  ___ | | __ _   _ 
| | | | / _ \/ __|| |/ /| | | |
| |/ / |  __/\__ \|   < | |_| |
|___/   \___||___/|_|\_\ \__, |
                          __/ |
                         |___/
===============================   
 - Desky succsessfully started on: %s:%s
 - Proxmox module is up: %v
 - Docker module is up: %v
===============================                   
`

type (
	Settings struct {
		Address struct {
			Ip   string `yaml:"ip"`
			Port string `yaml:"port"`
		} `yaml:"Address"`
		Background string `yaml:"Background"`
		Tls        struct {
			Enable bool   `yaml:"enable"`
			Crt    string `yaml:"crt"`
			Key    string `yaml:"key"`
		} `yaml:"TLS"`
		SessionStoreKey string `yaml:"Sessionkey"`
		Proxmox         struct {
			Up    bool        `yaml:"up"`
			Nodes []ProxmNode `yaml:"Nodes"`
		} `yaml:"Proxmox"`
		Docker struct {
			Up  bool   `yaml:"up"`
			URL string `yaml:"url"`
			Key string `yaml:"key"`
		} `yaml:"Docker"`
		Notifications struct {
			Gotify struct {
				URL string `yaml:"url"`
				KEY string `yaml:"key"`
			} `yaml:"Gotify"`
		} `yaml:"Notifications"`
	}

	ProxmNode struct {
		Node     string `yaml:"node"`
		User     string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"hostname"`
		Port     int    `yaml:"port"`
	}
)

const letterBytes = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func ParseSettings() Settings {
	file, err := os.Open("init/settings.yml")
	if err != nil {
		log.Fatal(err.Error())
	}

	var cfg Settings
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	if !cfg.Proxmox.Up {
		cfg.Proxmox.Up = false
	}
	if !cfg.Docker.Up {
		cfg.Docker.Up = false
	}
	if !cfg.Tls.Enable {
		cfg.Tls.Enable = false
	}
	return cfg
}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (cfg *Settings) PrintLogo() {
	log.Printf(printDesky, cfg.Address.Ip, cfg.Address.Port, cfg.Proxmox.Up, cfg.Docker.Up)
}
