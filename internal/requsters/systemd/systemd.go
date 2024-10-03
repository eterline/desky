package systemd

import (
	"encoding/json"
	"errors"
	"os/exec"
	"slices"

	"github.com/eterline/desky/internal/requsters/system"
)

var (
	IncorrectCommandErr = errors.New("Incorrect systemd command")
	UnableStatus        = errors.New("Unable to get status")
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

func ServiceStatus(service string) (string, error) {
	c := exec.Command("systemctl", "status", service)
	output, err := c.CombinedOutput()

	if err != nil {
		return string(output), UnableStatus
	}
	return string(output), nil
}

func ServiceExecute(service, cmd string) error {
	cmds := []string{"start", "stop", "enable", "disable", "restart"}

	if slices.Contains(cmds, cmd) {
		c := exec.Command("systemctl", cmd, service)
		err := c.Run()
		if err != nil {
			return err
		}
		return nil
	}
	return IncorrectCommandErr
}
