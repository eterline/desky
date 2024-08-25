package server

import (
	"github.com/eterline/desky/internal/applets"
	"github.com/eterline/desky/internal/requsters/system"
)

type dashboardData struct {
	Apps    applets.Apps
	SysData system.SystemStats
}
