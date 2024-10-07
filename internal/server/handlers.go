package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *server) Dasboard(w http.ResponseWriter, r *http.Request) {
	t := assemblyTemplates(s, s.templates["dashboard"])
	data := initDashboard(s.configs)

	if err := templateExec(w, t, "index", data); err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
	}
}

func (s *server) Proxmox(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if s.configs.Proxmox.Up {
		t := assemblyTemplates(s, s.templates["proxmox"])
		data, err := initProxmox(s.configs, s.proxmoxClient, vars["host"])
		if err != nil {
			log.Println(err)
			s.Home(w, r)
			return
		}
		if err := templateExec(w, t, "index", data); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		return

	}
	s.Home(w, r)
}

func (s *server) SysInfo(w http.ResponseWriter, r *http.Request) {
	t := assemblyTemplates(s, s.templates["monitor"])
	data := initSysInfo(s.configs)

	if err := templateExec(w, t, "index", data); err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
	}
}

func (s *server) Systemd(w http.ResponseWriter, r *http.Request) {
	t := assemblyTemplates(s, s.templates["systemd"])
	data := initSystemd(s.configs)

	if err := templateExec(w, t, "index", data); err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
	}
}

func (s *server) Tty(w http.ResponseWriter, r *http.Request) {
	t := assemblyTemplates(s, s.templates["tty"])
	data := initTty(s.configs)

	if err := templateExec(w, t, "index", data); err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
	}
}

func (s *server) Opnsense(w http.ResponseWriter, r *http.Request) {
	t := assemblyTemplates(s, s.templates["opnsense"])
	data := initOpnsense(s.configs, s.opnsenseClient)

	if err := templateExec(w, t, "index", data); err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
	}
}

func (s *server) Home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/dashboard/panel", http.StatusTemporaryRedirect)
}
