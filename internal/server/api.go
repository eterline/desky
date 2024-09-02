package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eterline/desky/internal/requsters/system"
	"github.com/gorilla/mux"
)

func (s *server) apiQm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	switch vars["cmd"] {
	case "start":
		go system.ExecCmd(fmt.Sprintf("qm start %s", id))
	case "shutdown":
		go system.ExecCmd(fmt.Sprintf("qm shutdown %s", id))
	case "reboot":
		go system.ExecCmd(fmt.Sprintf("qm restart %s", id))
	default:
		s.error(w, r, http.StatusBadRequest, nil)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "{\"message\":\"OK\"}")
}

func (s *server) apiPct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	switch vars["cmd"] {
	case "start":
		go system.ExecCmd(fmt.Sprintf("pct start %s", id))
	case "shutdown":
		go system.ExecCmd(fmt.Sprintf("pct shutdown %s", id))
	case "reboot":
		go system.ExecCmd(fmt.Sprintf("pct restart %s", id))
	default:
		s.error(w, r, http.StatusBadRequest, nil)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "{\"message\":\"OK\"}")
}

func (s *server) apiSystem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	st := system.GetStats()
	json.NewEncoder(w).Encode(st)
}
