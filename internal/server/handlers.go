package server

import (
	"html/template"
	"log"
	"net/http"

	"github.com/eterline/desky/internal/config"
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
	} else {
		log.Printf("| %s > Log In failed. Attempt username: %s", r.RemoteAddr, username)
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
	t := assemblyTempls(s.templates.dashboard)

	data := initDashboard(s.configs)
	templExecute(w, t, "index", data)
}

func (s *server) goDocker(w http.ResponseWriter, r *http.Request) {
	t := assemblyTempls(s.templates.docker)
	templExecute(w, t, "index", initDocker(s.configs))
}

func (s *server) goProxmox(w http.ResponseWriter, r *http.Request) {
	t := assemblyTempls(s.templates.proxmox)
	d := initProxmox(s.configs)
	templExecute(w, t, "index", d)
}

func (s *server) goTty(w http.ResponseWriter, r *http.Request) {
	t := assemblyTempls(s.templates.tty)
	templExecute(w, t, "index", initTty(s.configs))
}

func (s *server) goSysInfo(w http.ResponseWriter, r *http.Request) {
	t := assemblyTempls(s.templates.tty)
	templExecute(w, t, "index", initSysInfo(s.configs))
}

func (s *server) goHome(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/dashboard/panel", http.StatusTemporaryRedirect)
}

func assemblyTempls(templ ...string) *template.Template {
	templ = append(templ, "templates/index.html")
	return template.Must(template.ParseFiles(templ...))
}
