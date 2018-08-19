package psu

import (
	"math"
	"path/filepath"
)

// PowerSupplies represents all power supplies found on the system
type PowerSupplies struct {
	Mains     []*Mains
	Batteries []*Battery
}

// GetBatteriesEnergyNow returns EnergyNow for all batteries
func (psus *PowerSupplies) GetBatteriesEnergyNow() int64 {
	var nowSum int64
	for _, b := range psus.Batteries {
		nowSum = nowSum + b.GetEnergyNow()
	}
	return nowSum
}

// GetBatteriesEnergyFull returns EnergyFull for all batteries
func (psus *PowerSupplies) GetBatteriesEnergyFull() int64 {
	var fullSum int64
	for _, b := range psus.Batteries {
		fullSum = fullSum + b.GetEnergyFull()
	}
	return fullSum
}

// GetBatteriesPowerNow returns PowerNow for all batteries
func (psus *PowerSupplies) GetBatteriesPowerNow() int64 {
	var powerSum int64
	for _, b := range psus.Batteries {
		powerSum = powerSum + b.GetPowerNow()
	}
	return powerSum
}

// GetBatteriesCapacityPercent returns the capacity in percent for all batteries
func (psus *PowerSupplies) GetBatteriesCapacityPercent() float64 {
	return (float64(100) * float64(psus.GetBatteriesEnergyNow()) / float64(psus.GetBatteriesEnergyFull()))
}

// GetBatteriesCapacityTime returns the capacity in time left in hours, minutes
// and seconds for all batteries
func (psus *PowerSupplies) GetBatteriesCapacityTime() (int64, int64, int64) {
	p := float64(psus.GetBatteriesPowerNow())
	if p == 0 {
		return 0, 0, 0 // AC Power
	}

	e := float64(psus.GetBatteriesEnergyNow())
	f := float64(psus.GetBatteriesEnergyFull())
	var timeLeft float64

	switch psus.GetBatteriesStatus() {
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

// GetBatteriesStatus returns the Status of all batteries combined
func (psus *PowerSupplies) GetBatteriesStatus() string {
	var maxStatus string
	var maxStatusValue int

	statusMap := map[string]int{
		"Unknown":      0,
		"Full":         1,
		"Not Charging": 2,
		"Charging":     3,
		"Discharging":  4,
	}

	for _, b := range psus.Batteries {
		status := b.GetStatus()

		statusValue := statusMap[status]
		if statusValue > maxStatusValue {
			maxStatusValue = statusValue
		}
	}

	for k, v := range statusMap {
		if v == maxStatusValue {
			maxStatus = k
		}
	}

	return maxStatus
}

// GetPowerSupplies returns all power supplies found on the system
func GetPowerSupplies() (*PowerSupplies, error) {
	psus := &PowerSupplies{}

	paths, err := filepath.Glob(PowerSupplyPath + "*")
	if err != nil {
		return psus, err
	}

	for _, path := range paths {
		supplyName := filepath.Base(path)
		if supplyType, err := getPowerSupplyType(supplyName); err == nil {
			if supplyType == "Battery" {
				psus.Batteries = append(psus.Batteries, NewBattery(supplyName))
			} else if supplyType == "Mains" {
				psus.Mains = append(psus.Mains, NewMains(supplyName))
			}
		}
	}

	return psus, nil
}
