package server

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/eterline/desky/internal/config"
	"github.com/gorilla/mux"
	"github.com/luthermonson/go-proxmox"
)

type Routes map[string]Handle
type Handle struct {
	http.HandlerFunc
	string
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	log.Println(err)
	http.Error(w, err.Error(), code)
	s.respond(w, r, code)
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int) {
	w.WriteHeader(code)
}

func checkUser(s config.Settings, u string, p string) bool {
	h := sha256.New()
	h.Write([]byte(p))
	hex.EncodeToString(h.Sum(nil))
	return s.User.Username == u && s.User.Password == ToHash(p)
}

func checkAPIKey(s config.Settings, key string) bool {
	return s.APIKey == key
}

func ToHash(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func errLog(err error) {
	if err != nil {
		log.Println(TEMPLATE_ERR, err.Error())
	}
}

// Build subrouter with middlewares
// <-- baseRouter, map["route"]{http.HandlerFunc, "GET POST ANY"}, "/main", [...]middleware
func BuildSubRoute(router *mux.Router, routes Routes, path string, middlewares ...mux.MiddlewareFunc) *mux.Router {
	sub := router.PathPrefix(path).Subrouter()
	for route, handler := range routes {
		methods := strings.Split(handler.string, " ")
		sub.HandleFunc(route, handler.HandlerFunc).Methods(methods...)
	}

	if len(middlewares) > 0 {
		sub.Use(middlewares...)
	}
	return sub
}

func (s *server) dirHandleFiles(path, route, dir string) *mux.Route {
	filesDir := http.StripPrefix(route, http.FileServer(http.Dir(dir)))
	return s.router.PathPrefix(path).Handler(filesDir)
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

func idFromStr(id string) int {
	res, err := strconv.Atoi(id)
	if err != nil {
		return 0
	}
	return res
}
