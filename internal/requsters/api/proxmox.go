package api

import (
	"encoding/json"
	"sync"
)

const (
	VM_LIST = `qm list  | awk '
NR > 1 { 
    printf "{\"VMID\": \"%s\", \"NAME\": \"%s\", \"STATUS\": \"%s\", \"MEM(MB)\": \"%s\", \"BOOTDISK(GB)\": \"%s\", \"PID\": \"%s\"}\n", $1, $2, $3, $4, $5, $6 
}' | jq -s '.'`

	LXC_LIST = `pct list  | awk '
NR > 1 { 
    printf "{\"VMID\": \"%s\", \"Status\": \"%s\", \"Lock\": \"%s\", \"Name\": \"%s\"}\n", $1, $2, $3, $4 
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
	// out, err := system.ExecCmd(VM_LIST)
	// if err != nil {
	// 	v = nil
	// }
	out := []byte(testVm)
	json.Unmarshal(out, &v)

}

func (v *LXCList) Parse() {
	// out, err := system.ExecCmd(LXC_LIST)
	// if err != nil {
	// 	v = nil
	// }
	out := []byte(testLXC)
	json.Unmarshal(out, &v)
}

const (
	testVm = `
[
  {
    "VMID": "250",
    "NAME": "NAS",
    "STATUS": "running",
    "MEM(MB)": "6144",
    "BOOTDISK(GB)": "48.00",
    "PID": "2449"
  },
  {
    "VMID": "300",
    "NAME": "Win-Server",
    "STATUS": "stopped",
    "MEM(MB)": "4092",
    "BOOTDISK(GB)": "48.00",
    "PID": "0"
  }
]
	`
	testLXC = `
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
    "Status": "running",
    "Lock": "",
    "Name": "test-debian"
  },
  {
    "VMID": "500",
    "Status": "running",
    "Lock": "",
    "Name": "uptimekuma"
  }
]
	`
)
