package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/eterline/desky/internal/config"
)

type Container struct {
	Command      string `json:'Command'`
	CreatedAt    string `json:'CreatedAt'`
	ID           string `json:'ID'`
	Image        string `json:'Image'`
	Labels       string `json:'Labels'`
	LocalVolumes string `json:'LocalVolumes'`
	Mounts       string `json:'Mounts'`
	Names        string `json:'Names'`
	Networks     string `json:'Networks'`
	Port         string `json:'Port'`
	Ports        string `json:'Ports'`
	RunningFor   string `json:'RunningFor'`
	Size         string `json:'Size'`
	State        string `json:'State'`
	Status       string `json:'Status'`
}

type ContainerList []Container

func dockerRequest(s config.Settings) ([]byte, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", s.Docker.URL, nil)
	if err != nil {
		return *new([]byte), err
	}
	req.Header = http.Header{
		"Password": {s.Docker.Key},
	}
	res, err := client.Do(req)
	if err != nil {
		return *new([]byte), err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return *new([]byte), err
	}
	return body, nil
}

func DockerContainers(s config.Settings) ContainerList {
	var list ContainerList
	res, err := dockerRequest(s)
	if err != nil {
		return list
	}
	err = json.Unmarshal(res, &list)
	if err != nil {
		return list
	}
	return list
}
