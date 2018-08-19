package psu

import (
	"math"
	"path/filepath"
	"strconv"
)

// Battery represents a battery
type Battery struct {
	Name string
}

// GetStatus returns the Status of the battery
func (b *Battery) GetStatus() string {
	status, _ := stringFromFile(filepath.Join(PowerSupplyPath, b.Name, "status"))
	return status
}

// GetEnergyNow returns EnergyNow of the battery
func (b *Battery) GetEnergyNow() int64 {
	chargeNow, _ := stringFromFile(filepath.Join(PowerSupplyPath, b.Name, "energy_now"))
	i, _ := strconv.ParseInt(chargeNow, 10, 64)
	return i
}

// GetEnergyFull returns EnergyFull of the battery
func (b *Battery) GetEnergyFull() int64 {
	chargeFull, _ := stringFromFile(filepath.Join(PowerSupplyPath, b.Name, "energy_full"))
	i, _ := strconv.ParseInt(chargeFull, 10, 64)
	return i
}

// GetPowerNow returns PowerNow of the battery
func (b *Battery) GetPowerNow() int64 {
	powerNow, _ := stringFromFile(filepath.Join(PowerSupplyPath, b.Name, "power_now"))
	i, _ := strconv.ParseInt(powerNow, 10, 64)
	return i
}

// GetCapacityPercent returns the capacity in percent of the battery
func (b *Battery) GetCapacityPercent() float64 {
	return float64(100) * float64(b.GetEnergyNow()) / float64(b.GetEnergyFull())
}

// GetCapacityTime returns the time left in hours and minutes of the battery
func (b *Battery) GetCapacityTime() (int64, int64, int64) {
	p := float64(b.GetPowerNow())
	if p == 0 {
		return 0, 0, 0 // AC Power
	}

	e := float64(b.GetEnergyNow())
	f := float64(b.GetEnergyFull())
	var timeLeft float64

	switch b.GetStatus() {
	case "Charging":
		timeLeft = (f - e) / p
	case "Discharging":
		timeLeft = e / p
	}

	hoursLeft, hoursFractLeft := math.Modf(timeLeft)
	minutesLeft, minFractLeft := math.Modf(hoursFractLeft * 60.0)
	secondsLeft := minFractLeft * 60.0
	return int64(hoursLeft), int64(minutesLeft), int64(secondsLeft)
}

// NewBattery returns a new Battery
func NewBattery(name string) *Battery {
	bat := &Battery{
		Name: name,
	}
	return bat
}
