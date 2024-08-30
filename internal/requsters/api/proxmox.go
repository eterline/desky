package api

import (
	"encoding/json"
	"sync"

	"github.com/eterline/desky/internal/requsters/system"
)

const (
	VM_LIST = `qm list  | awk '
NR > 1 { 
    printf "{\"VMID\": \"%s\", \"NAME\": \"%s\", \"STATUS\": \"%s\", \"MEM(MB)\": \"%s\", \"BOOTDISK(GB)\": \"%s\", \"PID\": \"%s\"}\n", $1, $2, $3, $4, $5, $6 
}' | jq -s '.'`

	LXC_LIST = `pct list  | awk '
NR > 1 { 
    printf "{\"VMID\": \"%s\", \"Status\": \"%s\", \"Lock\": \"%s\", \"Name\": \"%s\"}\n", $1, $2, $4, $3 
}' | jq -s '.'`
)

type VM struct {
	VMID   string `json:'VMID'`
	Name   string `json:'NAME'`
	Status string `json:'STATUS'`
	Mem    string `json:'MEM(MB)'`
	Disk   string `json:'BOOTDISK(GB)'`
	PID    string `json:'PID'`
}

type LXC struct {
	VMID   string `json:'VMID'`
	Status string `json:'Status'`
	Lock   string `json:'Lock'`
	Name   string `json:'Name'`
}

type VMList []VM
type LXCList []LXC

type Virt struct {
	VMList
	LXCList
}

var testVar = `
[
  {
    "VMID": "100",
    "Status": "running",
    "Lock": "",
    "Name": "nginx"
  },
  {
    "VMID": "110",
    "Status": "running",
    "Lock": "",
    "Name": "dns"
  },
  {
    "VMID": "115",
    "Status": "running",
    "Lock": "",
    "Name": "prometheus"
  },
  {
    "VMID": "130",
    "Status": "stopped",
    "Lock": "",
    "Name": "test-debian"
  },
  {
    "VMID": "800",
    "Status": "running",
    "Lock": "",
    "Name": "alert"
  }
]
`

func VirtHostRequest() (VMList, LXCList) {
	var (
		vm  VMList
		lxc LXCList
		wg  sync.WaitGroup
	)
	wg.Add(2)
	go func() {
		lxc.Parse()
		wg.Done()
	}()
	go func() {
		vm.Parse()
		wg.Done()
	}()
	wg.Wait()
	return vm, lxc
}

func (v *VMList) Parse() {
	out, err := system.ExecCmd(VM_LIST)
	if err != nil {
		v = nil
	}
	json.Unmarshal(out, &v)

}

func (v *LXCList) Parse() {
	out, err := system.ExecCmd(LXC_LIST)
	if err != nil {
		v = nil
	}
	json.Unmarshal([]byte(out), &v)
}
