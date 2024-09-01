package server

import (
	"github.com/eterline/desky/internal/applets"
	"github.com/eterline/desky/internal/config"
	"github.com/eterline/desky/internal/requsters/api"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type (
	// Server structure data
	server struct {
		router       *mux.Router
		sessionStore sessions.Store
		templates    paths
		configs      config.Settings
		cookieKey    string
	}

	// html pages names
	paths struct {
		index     string
		login     string
		dashboard string
		notFound  string
		docker    string
		proxmox   string
		tty       string
	}

	dashboardData struct {
		Apps       applets.Apps
		Host       string
		Board      string
		Cpu        string
		Background string
		Auth       bool
	}

	dockerData struct {
		Host       string
		Containers api.ContainerList
		Background string
		Auth       bool
	}

	proxmoxData struct {
		Host       string
		LXCs       api.LXCList
		VMs        api.VMList
		Background string
		Auth       bool
	}

	ctxKey int8
)

const (
	SessionName       string = "DeskySession" // Name of Cookie
	ctxKeyUser        ctxKey = iota
	TEMPLATE_ERR             = "error parse template file:"
	EXEC_TEMPLATE_ERR        = "error execute template:"
)
