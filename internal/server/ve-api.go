package server

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/eterline/desky/pkg/ve"
	"github.com/gorilla/mux"
	"github.com/luthermonson/go-proxmox"
)

type Virtual interface {
	Start(ctx context.Context) (task *proxmox.Task, err error)
	Reboot(ctx context.Context) (task *proxmox.Task, err error)
}

func (s *server) apiQmExec(w http.ResponseWriter, r *http.Request) {
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

func (s *server) apiPctExec(w http.ResponseWriter, r *http.Request) {
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

func (s *server) apiPctList(w http.ResponseWriter, r *http.Request) {
	node, err := ve.Node(s.proxmoxClient, s.configs.Proxmox.Node, context.Background())
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
	}
	list, err := node.LXCList()
	if err != nil {
		s.error(w, r, http.StatusNotFound, err)
	}
	wrapJSON(w, struct{ VirtualMachines any }{list})
}

func (s *server) apiQmList(w http.ResponseWriter, r *http.Request) {
	node, err := ve.Node(s.proxmoxClient, s.configs.Proxmox.Node, context.Background())
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}
	list, err := node.VMList()
	if err != nil {
		s.error(w, r, http.StatusNotFound, err)
		return
	}
	wrapJSON(w, struct{ VirtualMachines any }{list})
}

func (s *server) apiPctInfo(w http.ResponseWriter, r *http.Request) {
	VMID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}
	node, err := ve.Node(s.proxmoxClient, s.configs.Proxmox.Node, context.Background())
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}
	info, err := node.VMget(VMID)
	if err != nil {
		s.error(w, r, http.StatusNotFound, err)
		return
	}
	wrapJSON(w, info)
}

func (s *server) apiQmInfo(w http.ResponseWriter, r *http.Request) {
	VMID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}
	node, err := ve.Node(s.proxmoxClient, s.configs.Proxmox.Node, context.Background())
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}
	info, err := node.VMget(VMID)
	if err != nil {
		s.error(w, r, http.StatusNotFound, err)
		return
	}
	wrapJSON(w, info)
}
func (s *server) apiNodeInfo(w http.ResponseWriter, r *http.Request) {
	node, err := ve.Node(s.proxmoxClient, s.configs.Proxmox.Node, context.Background())
	if err != nil {
		s.error(w, r, http.StatusInternalServerError, err)
		return
	}
	wrapJSON(w, node)
}
