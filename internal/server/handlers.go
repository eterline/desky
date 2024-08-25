package server

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/eterline/desky/internal/applets"
	"github.com/eterline/desky/internal/requsters/system"
)

type ctxKey int8

const (
	SessionName       string = "DeskySession"
	ctxKeyUser        ctxKey = iota
	TEMPLATE_ERR             = "error parse template file:"
	EXEC_TEMPLATE_ERR        = "error execute template:"
)

func (s *server) goLogin(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if chekUser(s.configs, username, password) {
		session, err := s.sessionStore.Get(r, SessionName)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		session.Values["uID"] = "1"
		err = s.sessionStore.Save(r, w, session)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		http.Redirect(w, r, "/dashboard/panel", http.StatusTemporaryRedirect)
		return
	}

	t, err := template.ParseFiles(s.templates.login)
	if err != nil {
		log.Println(TEMPLATE_ERR, err)
		return
	}
	err = t.ExecuteTemplate(w, s.templates.login, nil)
	if err != nil {
		log.Println(EXEC_TEMPLATE_ERR, EXEC_TEMPLATE_ERR)
		return
	}
}

func (s *server) goDasboard(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(s.templates.dashboard)
	if err != nil {
		log.Println(TEMPLATE_ERR, err)
		return
	}

	var data dashboardData

	data.Apps = applets.ParseApps()
	data.SysData = system.GetStats()
	err = t.ExecuteTemplate(w, s.templates.dashboard, data)
	if err != nil {
		log.Println(EXEC_TEMPLATE_ERR, EXEC_TEMPLATE_ERR)
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
	err = t.ExecuteTemplate(w, s.templates.notFound, nil)
	if err != nil {
		log.Println(EXEC_TEMPLATE_ERR, EXEC_TEMPLATE_ERR)
		return
	}
}

func (s *server) apiSystem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	st := system.GetStats()
	json.NewEncoder(w).Encode(st)
}
