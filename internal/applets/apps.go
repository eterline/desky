package applets

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type (
	App struct {
		Name string `json:"name"`
		Icon string `json:"icon"`
		Url  string `json:"url"`
		Desc string `json:"desciprion"`
	}

	Apps []App
)

func ParseApps() Apps {
	file, err := os.Open("init/apps.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	var appList Apps
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&appList)
	if err != nil {
		log.Panic(err.Error())
	}
	return appList
}
