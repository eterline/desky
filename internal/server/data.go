package server

import (
	"html/template"
	"net/http"
	"os"

	"github.com/eterline/desky/internal/applets"
	"github.com/eterline/desky/internal/config"
	"github.com/eterline/desky/internal/requsters/api"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	proxmox "github.com/luthermonson/go-proxmox"
	"github.com/zcalusic/sysinfo"
)

type (
	// Server structure data
	server struct {
		router        *mux.Router
		sessionStore  sessions.Store
		templates     paths
		configs       config.Settings
		cookieKey     string
		proxmoxClient *proxmox.Client
	}

	// html pages names
	paths struct {
		index     string
		login     string
		dashboard string
		docker    string
		proxmox   string
		monitor   string
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
		HostData   interface{}
		LXCs       interface{}
		VMs        interface{}
		Background string
		Auth       bool
	}

	ttyData struct {
		Host       string
		Background string
		Auth       bool
	}

	sysInfoData struct {
		Host       string
		Background string
		Auth       bool
		Info       sysinfo.SysInfo
		Systemd    any
		Smarts     any
	}

	ctxKey int8
)

const (
	SessionName       string = "DeskySession" // Name of Cookie
	ctxKeyUser        ctxKey = iota
	TEMPLATE_ERR             = "error parse template file:"
	EXEC_TEMPLATE_ERR        = "error execute template:"
)

func findHostname() string {
	host, err := os.Hostname()
	if err != nil {
		return "Unknown"
	}
	return host
}

func templateExec(w http.ResponseWriter, t *template.Template, templ string, data any) error {
	return t.ExecuteTemplate(w, templ, data)
}

func assemblyTemplates(s *server, templ ...string) *template.Template {
	templ = append(templ, s.templates.index)
	return template.Must(template.ParseFiles(templ...))
}
