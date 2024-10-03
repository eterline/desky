package server

import (
	"github.com/eterline/desky/internal/applets"
	"github.com/eterline/desky/internal/config"
	"github.com/eterline/desky/internal/requsters/disk"
	"github.com/eterline/desky/internal/requsters/system"
	"github.com/eterline/desky/internal/requsters/systemd"
	"github.com/luthermonson/go-proxmox"
	"github.com/zcalusic/sysinfo"
)

func initProxmox(s config.Settings, proxm []*proxmox.Client, nodeId string) (proxmoxData, error) {
	// Setting hostname, ignores .domain: 'host.domain.lan' --> 'host'
	node, err := findProxmoxHost(s, nodeId, proxm)
	if err != nil {
		return proxmoxData{}, err
	}

	pcts, err := node.LXCList()
	if err != nil {
		return proxmoxData{}, err
	}
	vms, err := node.VMList()
	if err != nil {
		return proxmoxData{}, err
	}
	return proxmoxData{
		Host:       findHostname(),
		HostData:   node.Data(),
		VMs:        vms,
		LXCs:       pcts,
		Background: s.Background,
		ProxmNodes: s.Proxmox.Nodes,
	}, nil
}

func initDashboard(s config.Settings) dashboardData {
	return dashboardData{
		Apps:       applets.ParseApps(),
		Host:       findHostname(),
		Board:      system.BoardModel(),
		Cpu:        system.CpuModel(),
		Background: s.Background,
		ProxmNodes: s.Proxmox.Nodes,
	}
}

func initTty(s config.Settings) ttyData {
	return ttyData{
		Host:       findHostname(),
		Background: s.Background,
		ProxmNodes: s.Proxmox.Nodes,
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
		ProxmNodes: s.Proxmox.Nodes,
	}
}
