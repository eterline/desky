package server

import (
	"github.com/eterline/desky/internal/applets"
	"github.com/eterline/desky/internal/config"
	"github.com/eterline/desky/internal/requsters/api"
	"github.com/eterline/desky/internal/requsters/system"
	"github.com/eterline/desky/internal/requsters/systemd"
	"github.com/zcalusic/sysinfo"
)

func initProxmox(s config.Settings) proxmoxData {
	vm, lxc := api.VirtHostRequest()
	return proxmoxData{
		Host:       findHostname(),
		VMs:        vm,
		LXCs:       lxc,
		Background: s.Background,
		Auth:       s.Auth,
	}
}

func initDashboard(s config.Settings) dashboardData {
	return dashboardData{
		Apps:       applets.ParseApps(),
		Host:       findHostname(),
		Board:      system.BoardModel(),
		Cpu:        system.CpuModel(),
		Background: s.Background,
		Auth:       s.Auth,
	}
}

func initDocker(s config.Settings) dockerData {
	return dockerData{
		Host:       findHostname(),
		Containers: api.DockerContainers(s),
		Background: s.Background,
		Auth:       s.Auth,
	}
}

func initTty(s config.Settings) ttyData {
	return ttyData{
		Host:       findHostname(),
		Background: s.Background,
		Auth:       s.Auth,
	}
}

func initSysInfo(s config.Settings) sysInfoData {
	var inf sysinfo.SysInfo
	inf.GetSysInfo()
	return sysInfoData{
		Host:       findHostname(),
		Background: s.Background,
		Auth:       s.Auth,
		Info:       inf,
		Systemd:    systemd.UnitsList(),
	}
}
