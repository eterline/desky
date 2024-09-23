package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/eterline/desky/internal/requsters/system"
	"github.com/eterline/desky/pkg/ve"
	"github.com/gorilla/mux"
)

const (
	serverErr  = "{\"message\":\"InternalServerError\"}"
	requestOK  = "{\"message\":\"OK\"}"
	requestBad = "{\"message\":\"BadRequest\"}"
)

func (s *server) apiQm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	host := strings.Split(s.configs.Proxmox.Host, ".")[0]
	node, err := ve.Node(s.proxmoxClient, host, context.Background())
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, serverErr)
		return
	}

	vm, err := node.VirtualMachine(
		context.Background(),
		idFromStr(vars["id"]),
	)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, serverErr)
		return
	}

	switch vars["cmd"] {
	case "start":
		go vm.Start(context.Background())
	case "shutdown":
		go vm.Shutdown(context.Background())
	case "reboot":
		go vm.Reboot(context.Background())
	default:
		s.error(w, r, http.StatusBadRequest, nil)
		fmt.Fprintf(w, requestBad)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, requestOK)
}

func (s *server) apiPct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	host := strings.Split(s.configs.Proxmox.Host, ".")[0]
	node, err := ve.Node(s.proxmoxClient, host, context.Background())
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, serverErr)
		return
	}

	pct, err := node.Container(
		context.Background(),
		idFromStr(vars["id"]),
	)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, serverErr)
		return
	}

	switch vars["cmd"] {
	case "start":
		go pct.Start(context.Background())
	case "shutdown":
		go pct.Shutdown(context.Background(), false, 0)
	case "reboot":
		go pct.Reboot(context.Background())
	default:
		s.error(w, r, http.StatusBadRequest, nil)
		fmt.Fprintf(w, requestBad)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, requestOK)
}

func (s *server) apiSystem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	st := system.GetStats()
	json.NewEncoder(w).Encode(st)
}

func idFromStr(id string) int {
	res, err := strconv.Atoi(id)
	if err != nil {
		return 0
	}
	return res
}
