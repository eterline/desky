package server

import (
	"context"
	"strings"

	"github.com/eterline/desky/internal/applets"
	"github.com/eterline/desky/internal/config"
	"github.com/eterline/desky/internal/requsters/api"
	"github.com/eterline/desky/internal/requsters/disk"
	"github.com/eterline/desky/internal/requsters/system"
	"github.com/eterline/desky/internal/requsters/systemd"
	"github.com/eterline/desky/pkg/ve"
	"github.com/luthermonson/go-proxmox"
	"github.com/zcalusic/sysinfo"
)

func initProxmox(s config.Settings, proxm *proxmox.Client) proxmoxData {
	// Setting hostname, ignores .domain: 'host.domain.lan' --> 'host'
	hst := strings.Split(s.Proxmox.Host, ".")[0]

	node, err := ve.Node(proxm, hst, context.Background())
	if err != nil {
		return proxmoxData{}
	}
	pcts, err := node.LXCList()
	if err != nil {
		return proxmoxData{}
	}
	vms, err := node.VMList()
	if err != nil {
		return proxmoxData{}
	}
	return proxmoxData{
		Host:       findHostname(),
		HostData:   node.Data(),
		VMs:        vms,
		LXCs:       pcts,
		Background: s.Background,
	}
}

func initDashboard(s config.Settings) dashboardData {
	return dashboardData{
		Apps:       applets.ParseApps(),
		Host:       findHostname(),
		Board:      system.BoardModel(),
		Cpu:        system.CpuModel(),
		Background: s.Background,
	}
}

func initDocker(s config.Settings) dockerData {
	return dockerData{
		Host:       findHostname(),
		Containers: api.DockerContainers(s),
		Background: s.Background,
	}
}

func initTty(s config.Settings) ttyData {
	return ttyData{
		Host:       findHostname(),
		Background: s.Background,
	}
}

func initSysInfo(s config.Settings) sysInfoData {
	var inf sysinfo.SysInfo
	inf.GetSysInfo()
	smarts := disk.SmartDisks(inf.Storage)

	return sysInfoData{
		Host:       findHostname(),
		Background: s.Background,
		Info:       inf,
		Systemd:    systemd.UnitsList(),
		Smarts:     smarts,
	}
}
