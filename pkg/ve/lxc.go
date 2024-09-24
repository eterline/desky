package ve

import (
	"context"
	"fmt"
	"sort"

	"github.com/luthermonson/go-proxmox"
)

type LXC struct {
	ID     int
	Name   string
	CPUs   int
	Uptime string
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
			Uptime: uptimeStr(i.Uptime),
			Status: i.Status,
			Tags:   i.Tags,
			CT:     i,
		})
	}

	sort.Slice(l, func(i, j int) (less bool) {
		return l[i].ID < l[j].ID
	})
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
		Uptime: uptimeStr(ct.Uptime),
		Status: ct.Status,
		Tags:   ct.Tags,
		CT:     ct,
	}, nil
}

func (ct LXC) Shutdown() {
	ct.CT.Shutdown(context.Background(), false, 0)
}

func uptimeStr(time uint64) string {
	days := time / 86400
	hours := time % 86400 / 3600
	mins := time % 86400 % 3600 / 60
	sec := time % 86400 % 3600 % 60

	return fmt.Sprintf("%vd|%vh|%vm|%vs", days, hours, mins, sec)
}
