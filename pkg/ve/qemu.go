package ve

import (
	"context"
	"sort"

	"github.com/luthermonson/go-proxmox"
)

type VM struct {
	ID     int
	Name   string
	CPUs   int
	Uptime string
	Status string
	Tags   string
	VM     *proxmox.VirtualMachine
}

func (node *VENode) VMList() ([]VM, error) {
	vms, err := node.VirtualMachines(node.Context)
	if err != nil {
		return nil, err
	}
	var l []VM
	for _, i := range vms {
		l = append(l, VM{
			ID:     int(i.VMID),
			Name:   i.Name,
			CPUs:   i.CPUs,
			Uptime: uptimeStr(i.Uptime),
			Status: i.Status,
			Tags:   i.Tags,
			VM:     i,
		})
	}

	sort.Slice(l, func(i, j int) (less bool) {
		return l[i].ID < l[j].ID
	})
	return l, nil
}

func (node *VENode) VMget(id int) (VM, error) {
	vm, err := node.VirtualMachine(node.Context, id)
	if err != nil {
		return VM{}, err
	}
	return VM{
		ID:     int(vm.VMID),
		Name:   vm.Name,
		CPUs:   vm.CPUs,
		Uptime: uptimeStr(vm.Uptime),
		Status: vm.Status,
		Tags:   vm.Tags,
		VM:     vm,
	}, nil
}

func (ct VM) Shutdown() {
	ct.VM.Shutdown(context.Background())
}
