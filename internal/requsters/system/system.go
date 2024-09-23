package system

import (
	"encoding/json"
	"os/exec"
)

var mainboardInfo = `
    local mainboard_name mainboard_vendor mainboard_version \

    mainboard_name="$(< /sys/devices/virtual/dmi/id/board_name)" \
    mainboard_vendor="$(< /sys/devices/virtual/dmi/id/board_vendor)" \
    mainboard_version="$(< /sys/devices/virtual/dmi/id/board_version)" \

    printf -v mainboard \
        "%s %s (%s)" \
        "${mainboard_name}" \
        "${mainboard_version}" \
        "${mainboard_vendor}"
`

type SysStats struct {
	Cpu struct {
		Load int `json:"load"`
	} `json:"cpu"`
	Mem struct {
		Current int `json:"current"`
		Total   int `json:"total"`
		Load    int `json:"load"`
	} `json:"mem"`
	Disk struct {
		Current int `json:"current"`
		Total   int `json:"total"`
		Used    int `json:"used"`
	} `json:"disk"`
	Uptime string `json:"uptime"`
}

func GetStats() SysStats {
	var stat SysStats
	res, err := ExecCmd("./internal/scripts/sys-stats.sh")
	if err != nil {
		return stat
	}
	err = json.Unmarshal(res, &stat)
	if err != nil {
		return stat
	}
	return stat
}

func BoardModel() string {
	board, err := ExecCmd("cat /sys/devices/virtual/dmi/id/board_name")
	if err != nil {
		return "UnknownModel"
	}
	return string(board)
}

func CpuModel() string {
	board, err := ExecCmd(`lscpu | sed -nr '/Model name/ s/.*:\s*(.*) @ .*/\1/p'`)
	if err != nil {
		return "UnknownModel"
	}
	return string(board)
}

func Uptime() string {
	time, err := ExecCmd("uptime -p")
	if err != nil {
		return "null"
	}
	return string(time)
}

func ExecCmd(cmd string) ([]byte, error) {
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return nil, err
	}
	return out, nil
}
