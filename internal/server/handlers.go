package server

import (
	"html/template"
	"log"
	"net/http"

	"github.com/eterline/desky/internal/config"
)

func (s *server) Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user := r.FormValue("username")
	pass := r.FormValue("password")
	if chekUser(s.configs, user, pass) {
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
	} else {
		log.Printf("| %s > Log In failed. Attempt username: %s", r.RemoteAddr, user)
	}
	t := template.Must(template.ParseFiles(s.templates.login))
	err := t.ExecuteTemplate(w, "login.html", s.configs)
	if err != nil {
		log.Println(EXEC_TEMPLATE_ERR, err.Error())
		return
	}
}

func (s *server) Logout(w http.ResponseWriter, r *http.Request) {
	clearSession(s)
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}

func clearSession(s *server) {
	log.Println("User session has been cleared. New CookieKey generated.")
	s.cookieKey = config.RandStringBytes(32)
}

func (s *server) Dasboard(w http.ResponseWriter, r *http.Request) {
	t := assemblyTemplates(s, s.templates.dashboard)
	data := initDashboard(s.configs)

	if err := templateExec(w, t, "index", data); err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
	}
}

func (s *server) Docker(w http.ResponseWriter, r *http.Request) {
	t := assemblyTemplates(s, s.templates.docker)
	data := initDocker(s.configs)

	if err := templateExec(w, t, "index", data); err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
	}
}

func (s *server) Proxmox(w http.ResponseWriter, r *http.Request) {
	t := assemblyTemplates(s, s.templates.proxmox)
	data := initProxmox(s.configs, s.proxmoxClient)

	if err := templateExec(w, t, "index", data); err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
	}
}

func (s *server) Tty(w http.ResponseWriter, r *http.Request) {
	t := assemblyTemplates(s, s.templates.tty)
	data := initTty(s.configs)

	if err := templateExec(w, t, "index", data); err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
	}
}

func (s *server) SysInfo(w http.ResponseWriter, r *http.Request) {
	t := assemblyTemplates(s, s.templates.monitor)
	data := initSysInfo(s.configs)

	if err := templateExec(w, t, "index", data); err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
	}
}

func (s *server) Home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/dashboard/panel", http.StatusTemporaryRedirect)
}
