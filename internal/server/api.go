package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/eterline/desky/internal/requsters/system"
	"github.com/eterline/desky/internal/requsters/systemd"
)

const (
	serverErr  = "{\"message\":\"InternalServerError\"}"
	requestOK  = "{\"message\":\"OK\"}"
	requestBad = "{\"message\":\"BadRequest\"}"
)

var errBadCommand = errors.New("Uncorrect command")

func (s *server) apiSystem(w http.ResponseWriter, r *http.Request) {
	wrapJSON(w, system.GetStats())
}

func (s *server) apiSystemdList(w http.ResponseWriter, r *http.Request) {
	wrapJSON(w, systemd.UnitsList())
}

func wrapJSON(w http.ResponseWriter, v any) error {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		return err
	}
	return nil
}
