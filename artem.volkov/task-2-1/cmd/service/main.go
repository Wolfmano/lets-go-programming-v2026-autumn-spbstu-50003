package main

import "fmt"

const (
	MinTemperature = 15
	MaxTemperature = 30

	LessOrEqual    = "<="
	GreaterOrEqual = ">="
)

type TemperatureManager struct {
	minTemp int
	maxTemp int
}

func NewTemperatureManager() TemperatureManager {
	return TemperatureManager{
		minTemp: MinTemperature,
		maxTemp: MaxTemperature,
	}
}

func (tm *TemperatureManager) UpdateTemperature(sign string, temp int) {
	switch sign {
	case LessOrEqual:
		if temp < tm.maxTemp {
			tm.maxTemp = temp
		}
	case GreaterOrEqual:
		if temp > tm.minTemp {
			tm.minTemp = temp
		}
	}
}

func (tm *TemperatureManager) GetTemperature() int {
	if tm.minTemp > tm.maxTemp {
		return -1
	}
	return tm.minTemp
}

func main() {
	var n, k int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}
	if _, err := fmt.Scan(&k); err != nil {
		return
	}
	for i := 0; i < n; i++ {
		manager := NewTemperatureManager()
		for j := 0; j < k; j++ {
			var sign string
			var temp int
			if _, err := fmt.Scan(&sign); err != nil {
				return
			}
			if _, err := fmt.Scan(&temp); err != nil {
				return
			}
			manager.UpdateTemperature(sign, temp)
			fmt.Println(manager.GetTemperature())
		}
	}
}
