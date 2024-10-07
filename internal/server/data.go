package server

import (
	"github.com/eterline/desky/internal/applets"
	"github.com/eterline/desky/internal/config"
	"github.com/eterline/desky/internal/requsters/systemd"
	opnsenseapi "github.com/eterline/opnsense-api"
	"github.com/gorilla/mux"
	"github.com/luthermonson/go-proxmox"
	"github.com/sirupsen/logrus"
	"github.com/zcalusic/sysinfo"
)

type (
	// Server structure data
	server struct {
		router         *mux.Router
		templates      map[string]string
		configs        config.Settings
		proxmoxClient  []*proxmox.Client
		opnsenseClient opnsenseapi.OpnsenseClient
		logger         *logrus.Logger
	}

	// html pages names
	paths struct {
		index      string
		dashboard  string
		docker     string
		proxmox    string
		monitor    string
		tty        string
		ProxmNodes []config.ProxmNode
	}

	dashboardData struct {
		Apps       applets.Apps
		Host       string
		Board      string
		Cpu        string
		Background string
		ProxmNodes []config.ProxmNode
	}

	proxmoxData struct {
		Host       string
		HostData   interface{}
		LXCs       interface{}
		VMs        interface{}
		Background string
		ProxmNodes []config.ProxmNode
	}

	ttyData struct {
		Host       string
		Background string
		ProxmNodes []config.ProxmNode
	}

	sysInfoData struct {
		Host       string
		Background string
		Info       sysinfo.SysInfo
		Systemd    any
		Smarts     any
		ProxmNodes []config.ProxmNode
	}

	systemdData struct {
		Host       string
		Background string
		Systemd    systemd.Systemctl
		ProxmNodes []config.ProxmNode
	}

	opnsenseData struct {
		Host       string
		Background string
		Firmware   opnsenseapi.FirmwareInfo
		Syslog     opnsenseapi.SyslogStats
		Wireguard  opnsenseapi.WgService
		ProxmNodes []config.ProxmNode
	}

	ctxKey int8
)

const (
	SessionName       string = "DeskySession" // Name of Cookie
	ctxKeyUser        ctxKey = iota
	TEMPLATE_ERR             = "error parse template file:"
	EXEC_TEMPLATE_ERR        = "error execute template:"
)
