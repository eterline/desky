package server

import (
	"html/template"
	"net/http"
	"os"

	"github.com/eterline/desky/internal/applets"
	"github.com/eterline/desky/internal/config"
	"github.com/eterline/desky/internal/requsters/api"
	"github.com/gorilla/mux"
	"github.com/luthermonson/go-proxmox"
	"github.com/sirupsen/logrus"
	"github.com/zcalusic/sysinfo"
)

type (
	// Server structure data
	server struct {
		router        *mux.Router
		templates     paths
		configs       config.Settings
		proxmoxClient []*proxmox.Client
		ProxmNodes    config.Settings
		logger        *logrus.Logger
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

	dockerData struct {
		Host       string
		Containers api.ContainerList
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
