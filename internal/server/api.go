package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/eterline/desky/internal/requsters/system"
	"github.com/eterline/desky/internal/requsters/systemd"
	"github.com/gorilla/mux"
)

const (
	serverErr  = "{\"message\":\"InternalServerError\"}"
	requestOK  = "{\"message\":\"OK\"}"
	requestBad = "{\"message\":\"BadRequest\"}"
)

var errBadCommand = errors.New("Uncorrect command")

func (s *server) apiSystem(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	wrapJSON(w, system.GetStats())
}

func (s *server) apiSystemdList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	wrapJSON(w, systemd.UnitsList())
}

func (s *server) apiSystemdStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stat, err := systemd.ServiceStatus(vars["service"])
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		fmt.Fprintf(w, requestBad)
		return
	}
	type status struct {
		Descripiton string `json:"description"`
	}
	w.WriteHeader(http.StatusOK)
	wrapJSON(w, status{stat})
}

func (s *server) apiSystemdExec(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cmd := vars["cmd"]
	service := vars["service"]
	err := systemd.ServiceExecute(service, cmd)
	if err != nil {
		s.error(w, r, http.StatusBadRequest, err)
		fmt.Fprintf(w, requestBad)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, requestOK)
}

func wrapJSON(w http.ResponseWriter, v any) error {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		return err
	}
	return nil
}
