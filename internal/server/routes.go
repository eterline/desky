package server

import (
	"net/http"
)

// Init private static content route: path /static/...
func (s *server) configPrivateStatic() {
	content := s.router.PathPrefix("/static/").Subrouter()
	content.PathPrefix("").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
}

// Init UI pages
func (s *server) configPagesRouter() {
	s.router.NotFoundHandler = http.HandlerFunc(s.Home)

	BuildSubRoute(s.router,
		Routes{
			"/ws":             {wsConnection, "GET"}, // ws connection for .../tty
			"/panel":          {s.Dasboard, "GET"},   // dashboard rendering page
			"/tty":            {s.Tty, "GET"},        // tty console page
			"/docker":         {s.Docker, "GET"},     // docker ct list page
			"/proxmox/{host}": {s.Proxmox, "GET"},    // proxmox devices and info page
			"/monitor":        {s.SysInfo, "GET"},    // system info page
		},
		"/dashboard",
	)
}

func (s *server) configApiRouter() {

	// Init base API subrouting: path /api/...
	api := BuildSubRoute(s.router,
		Routes{
			"/system":  {s.apiSystem, "GET"},
			"/systemd": {s.apiSystemdList, "GET"},
		},
		"/api",
	)

	// init handlers: path /api/proxmox/...
	BuildSubRoute(api,
		Routes{
			"/{host}/node":           {s.apiNodeInfo, "GET"}, // list pct devices
			"/{host}/pct":            {s.apiPctList, "GET"},  // list pct devices
			"/{host}/qm":             {s.apiQmList, "GET"},   // list qemu devices
			"/{host}/qm/{id}":        {s.apiQmInfo, "GET"},   // qemu device info
			"/{host}/pct/{id}":       {s.apiPctInfo, "GET"},  // pct device info
			"/{host}/pct/{id}/{cmd}": {s.apiPctExec, "POST"}, // reload|start|shutdown pct
			"/{host}/qm/{id}/{cmd}":  {s.apiQmExec, "POST"},  // reload|start|shutdown qemu
		},
		"/proxmox",
	)
}

// Serve Login page and public static files: path /public/...
func (s *server) configPublicRouter() {
	s.dirHandleFiles("/public/", "/public/", "./public")   //
	s.dirHandleFiles("/node/", "/node/", "./node_modules") // node modules tty js script
}
