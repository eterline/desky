package server

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/eterline/desky/internal/applets"
	"github.com/eterline/desky/internal/config"
	"github.com/eterline/desky/internal/requsters/api"
	"github.com/eterline/desky/internal/requsters/system"
)

func (s *server) goLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	if chekUser(s.configs, username, password) {
		session, err := s.sessionStore.Get(r, SessionName)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		session.Values["uID"] = s.cookieKey
		err = s.sessionStore.Save(r, w, session)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		log.Printf("Log In to panel from: %s.", r.RemoteAddr)
		http.Redirect(w, r, "/dashboard/panel", http.StatusTemporaryRedirect)
		return
	}

	log.Printf("| %s > Log In failed. Attempt username: %s", r.RemoteAddr, username)

	t, err := template.ParseFiles(s.templates.login)
	if err != nil {
		log.Println(TEMPLATE_ERR, err.Error())
		return
	}
	err = t.ExecuteTemplate(w, s.templates.login, s.configs)
	if err != nil {
		log.Println(EXEC_TEMPLATE_ERR, err.Error())
		return
	}
}

func (s *server) goLogOut(w http.ResponseWriter, r *http.Request) {
	clearSession(s)
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}

// Updates uID and makes Log Out user.
func clearSession(s *server) {
	log.Println("User session has been cleared. New CookieKey generated.")
	s.cookieKey = config.RandStringBytes(32)
}

func (s *server) goDasboard(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(s.templates.dashboard)
	if err != nil {
		log.Println(TEMPLATE_ERR, err.Error())
		return
	}
	var data dashboardData
	data.setDashboardData(s.configs)

	err = t.ExecuteTemplate(w, s.templates.dashboard, data)
	if err != nil {
		log.Println(EXEC_TEMPLATE_ERR, err.Error())
		return
	}
}

func (s *server) goDocker(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(s.templates.docker)
	if err != nil {
		log.Println(TEMPLATE_ERR, err.Error())
		return
	}

	var data dockerData
	data.setDockerData(s.configs)
	err = t.ExecuteTemplate(w, s.templates.docker, data)
	if err != nil {
		log.Println(EXEC_TEMPLATE_ERR, err)
		return
	}
}

func (s *server) goHome(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/dashboard/panel", http.StatusTemporaryRedirect)
}

func (s *server) goNotFound(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(s.templates.notFound)
	if err != nil {
		log.Println(TEMPLATE_ERR, err)
		return
	}
	err = t.ExecuteTemplate(w, s.templates.notFound, s.configs)
	if err != nil {
		log.Println(EXEC_TEMPLATE_ERR, EXEC_TEMPLATE_ERR)
		return
	}
}

func (s *server) goProxmox(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(s.templates.proxmox)
	if err != nil {
		log.Println(TEMPLATE_ERR, err.Error())
		return
	}

	var data proxmoxData
	data.setProxmoxData(s.configs)
	err = t.ExecuteTemplate(w, s.templates.proxmox, data)
	if err != nil {
		log.Println(EXEC_TEMPLATE_ERR, err)
		return
	}
}

func (s *server) apiSystem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	st := system.GetStats()
	json.NewEncoder(w).Encode(st)
}

func (d *proxmoxData) setProxmoxData(s config.Settings) {
	host, err := os.Hostname()
	if err != nil {
		d.Host = "Error"
	} else {
		d.Host = host
	}
	d.VMs, d.LXCs = api.VirtHostRequest()
	d.Background = s.Background
}

func (d *dashboardData) setDashboardData(s config.Settings) {
	d.Apps = applets.ParseApps()
	host, err := os.Hostname()
	if err != nil {
		d.Host = "Error"
	} else {
		d.Host = host
	}
	d.Board = system.BoardModel()
	d.Cpu = system.CpuModel()
	d.Background = s.Background
}

func (d *dockerData) setDockerData(s config.Settings) {
	host, err := os.Hostname()
	if err != nil {
		d.Host = "Error"
	} else {
		d.Host = host
	}
	d.Containers = api.DockerContainers(s)
	d.Background = s.Background
}
