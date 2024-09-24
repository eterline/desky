package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/eterline/desky/internal/requsters/system"
	"github.com/eterline/desky/pkg/ve"
	"github.com/gorilla/mux"
	"github.com/luthermonson/go-proxmox"
)

const (
	serverErr  = "{\"message\":\"InternalServerError\"}"
	requestOK  = "{\"message\":\"OK\"}"
	requestBad = "{\"message\":\"BadRequest\"}"
)

var errBadCommand = errors.New("Uncorrect command")

type Virtual interface {
	Start(ctx context.Context) (task *proxmox.Task, err error)
	Reboot(ctx context.Context) (task *proxmox.Task, err error)
}

func (s *server) apiQm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	host := strings.Split(s.configs.Proxmox.Host, ".")[0]

	node, err := ve.Node(s.proxmoxClient, host, context.Background())
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		fmt.Fprintf(w, serverErr)
		return
	}

	vm, err := node.VirtualMachine(
		context.Background(),
		idFromStr(vars["id"]),
	)
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		fmt.Fprintf(w, serverErr)
		return
	}
	err = execVirt(vm, vars["cmd"])
	if err != nil {
		s.error(w, r, http.StatusBadRequest, err)
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
		s.error(w, r, http.StatusInternalServerError, err)
		fmt.Fprintf(w, serverErr)
		return
	}

	pct, err := node.Container(context.Background(), idFromStr(vars["id"]))
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		fmt.Fprintf(w, serverErr)
		return
	}

	err = execVirt(pct, vars["cmd"])
	if err != nil {
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

func execVirt(virt Virtual, cmd string) error {
	var err error
	switch cmd {
	case "start":
		_, err = virt.Start(context.Background())
	case "reboot":
		_, err = virt.Reboot(context.Background())
	case "shutdown":
		switch v := virt.(type) {
		case *proxmox.Container:
			_, err = v.Shutdown(context.Background(), false, 0)
		case *proxmox.VirtualMachine:
			_, err = v.Shutdown(context.Background())
		}
	default:
		return errBadCommand
	}
	if err != nil {
		return err
	}
	return nil
}
