package server

import (
	"html/template"
	"log"
	"net/http"

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
		go system.ExecCmd("internal/scripts/tg-notify-login.sh")
		log.Printf("Log In to panel from: %s.", r.RemoteAddr)
		http.Redirect(w, r, "/dashboard/panel", http.StatusTemporaryRedirect)
		return
	}
	if username != "" && password != "" {
		log.Printf("| %s > Log In failed. Attempt username: %s", r.RemoteAddr, username)
		go system.ExecCmd("internal/scripts/tg-notify-failed.sh")
	}
	t := template.Must(template.ParseFiles(s.templates.login))
	err := t.ExecuteTemplate(w, "login.html", s.configs)
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
	t := template.Must(template.ParseFiles(s.templates.dashboard, s.templates.index))

	var data dashboardData
	data.setDashboardData(s.configs)

	templExecute(w, t, "index", data)
}

func (s *server) goDocker(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(s.templates.docker, s.templates.index))

	var data dockerData
	data.setDockerData(s.configs)

	templExecute(w, t, "index", data)
}

func (s *server) goProxmox(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(s.templates.proxmox, s.templates.index))

	var data proxmoxData
	data.setProxmoxData(s.configs)

	templExecute(w, t, "index", data)
}

func (s *server) goTty(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(s.templates.tty, s.templates.index))

	var data ttyData
	data.setTtyData(s.configs)

	templExecute(w, t, "index", data)
}

func (s *server) goSysInfo(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(s.templates.tty, s.templates.index))

	var data sysInfoData
	data.setSysInfoData(s.configs)

	templExecute(w, t, "index", data)
}

func (s *server) goHome(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/dashboard/panel", http.StatusTemporaryRedirect)
}

func (s *server) goNotFound(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(s.templates.notFound)
	errLog(err)
	err = t.ExecuteTemplate(w, s.templates.notFound, s.configs)
	if err != nil {
		log.Println(EXEC_TEMPLATE_ERR, EXEC_TEMPLATE_ERR)
		return
	}
}

func (d *proxmoxData) setProxmoxData(s config.Settings) {
	d.Host = findHostname()
	d.VMs, d.LXCs = api.VirtHostRequest()
	d.Background = s.Background
	d.Auth = s.Auth
}

func (d *dashboardData) setDashboardData(s config.Settings) {
	d.Host = findHostname()
	d.Board = system.BoardModel()
	d.Cpu = system.CpuModel()
	d.Background = s.Background
	d.Auth = s.Auth
}

func (d *dockerData) setDockerData(s config.Settings) {
	d.Host = findHostname()
	d.Containers = api.DockerContainers(s)
	d.Background = s.Background
	d.Auth = s.Auth
}

func (d *ttyData) setTtyData(s config.Settings) {
	d.Host = findHostname()
	d.Background = s.Background
	d.Auth = s.Auth
}

// TODO: Доделать вывод инфы о системе через вкладку System.
func (d *sysInfoData) setSysInfoData(s config.Settings) {
	d.Host = findHostname()
	d.Background = s.Background
	d.Auth = s.Auth
	d.Info.GetSysInfo()
}
