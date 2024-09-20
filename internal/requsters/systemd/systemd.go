package systemd

import (
	"encoding/json"

	"github.com/eterline/desky/internal/requsters/system"
)

type Service struct {
	Unit   string `json:"unit_file"`
	State  string `json:"state"`
	Preset string `json:"preset"`
}

type Systemctl []Service

func UnitsList() Systemctl {
	res, err := system.ExecCmd("systemctl list-unit-files --type=service -o json | jq .")
	if err != nil {
		return nil
	}
	var list Systemctl
	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil
	}
	return list
}
