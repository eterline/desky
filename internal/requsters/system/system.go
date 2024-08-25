package system

import "os"

type SystemStats struct {
	Host    string
	CpuLoad int
	MemLoad int
}

func GetStats() SystemStats {
	host, err := os.Hostname()
	if err != nil {
		host = "Error"
	}

	cpu := 0

	mem := 0

	return SystemStats{
		Host:    host,
		CpuLoad: cpu,
		MemLoad: mem,
	}
}
