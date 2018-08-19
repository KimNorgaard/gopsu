package psu

import (
	"path/filepath"
)

// PowerSupplyPath is where to find information about power supplies
const PowerSupplyPath string = "/sys/class/power_supply/"

func getPowerSupplyType(name string) (string, error) {
	supplyTypeFile := filepath.Join(PowerSupplyPath, name, "type")
	supplyType, err := stringFromFile(supplyTypeFile)
	return supplyType, err
}
