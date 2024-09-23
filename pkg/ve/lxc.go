package ve

import (
	"github.com/luthermonson/go-proxmox"
)

type LXC struct {
	ID     int
	Name   string
	CPUs   int
	Uptime uint64
	Status string
	Tags   string
	CT     *proxmox.Container
}

func (node *VENode) LXCList() ([]LXC, error) {
	cts, err := node.Containers(node.Context)
	if err != nil {
		return nil, err
	}
	var l []LXC
	for _, i := range cts {
		l = append(l, LXC{
			ID:     int(i.VMID),
			Name:   i.Name,
			CPUs:   i.CPUs,
			Uptime: i.Uptime,
			Status: i.Status,
			Tags:   i.Tags,
			CT:     i,
		})
	}
	return l, nil
}

func (node *VENode) LXCget(id int) (LXC, error) {
	ct, err := node.Container(node.Context, id)
	if err != nil {
		return LXC{}, err
	}
	return LXC{
		ID:     int(ct.VMID),
		Name:   ct.Name,
		CPUs:   ct.CPUs,
		Uptime: ct.Uptime,
		Status: ct.Status,
		Tags:   ct.Tags,
		CT:     ct,
	}, nil
}
