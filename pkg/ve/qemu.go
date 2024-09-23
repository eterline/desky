package ve

import (
	"github.com/luthermonson/go-proxmox"
)

type VM struct {
	ID     int
	Name   string
	CPUs   int
	Uptime uint64
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
			Uptime: i.Uptime,
			Status: i.Status,
			Tags:   i.Tags,
			VM:     i,
		})
	}
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
		Uptime: vm.Uptime,
		Status: vm.Status,
		Tags:   vm.Tags,
		VM:     vm,
	}, nil
}
