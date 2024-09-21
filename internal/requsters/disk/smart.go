package disk

import (
	"strings"

	"github.com/anatol/smart.go"
	"github.com/zcalusic/sysinfo"
)

func SmartDisks(devices []sysinfo.StorageDevice) map[string]smart.GenericAttributes {
	attrs := make(map[string]smart.GenericAttributes)
	for _, dev := range devices {
		name, smart := Smart("/dev/" + dev.Name)
		attrs[strings.Trim(name, "/dev/")] = *smart
	}
	return attrs
}

func Smart(disk string) (string, *smart.GenericAttributes) {
	dev, err := smart.Open(disk)
	if err != nil {
		return "", nil
	}
	defer dev.Close()
	attr, err := dev.ReadGenericAttributes()
	if err != nil {
		return "", nil
	}
	return disk, attr
}
