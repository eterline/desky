package server

import (
	"net/http"
)

func (s *server) Dasboard(w http.ResponseWriter, r *http.Request) {
	t := assemblyTemplates(s, s.templates.dashboard)
	data := initDashboard(s.configs)

	if err := templateExec(w, t, "index", data); err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
	}
}

func (s *server) Docker(w http.ResponseWriter, r *http.Request) {
	if s.configs.Docker.Up {
		t := assemblyTemplates(s, s.templates.docker)
		data := initDocker(s.configs)

		if err := templateExec(w, t, "index", data); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}
		return
	}
	s.Home(w, r)
}

func (s *server) Proxmox(w http.ResponseWriter, r *http.Request) {
	if s.configs.Proxmox.Up {
		t := assemblyTemplates(s, s.templates.proxmox)
		data := initProxmox(s.configs, s.proxmoxClient)

		if err := templateExec(w, t, "index", data); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
		}
		return
	}
	s.Home(w, r)
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
