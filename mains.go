package psu

import (
	"path/filepath"
	"strconv"
)

// Mains represents an AC power source
type Mains struct {
	Name string
}

// GetOnline gets the online status of the AC power source
func (m *Mains) GetOnline() bool {
	online, _ := stringFromFile(filepath.Join(PowerSupplyPath, m.Name, "online"))
	onlineB, _ := strconv.ParseBool(online)
	return onlineB
}

// NewMains returns a new Mains
func NewMains(name string) *Mains {
	mains := &Mains{
		Name: name,
	}
	return mains
}
