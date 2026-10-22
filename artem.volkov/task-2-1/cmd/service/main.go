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
	var departmentsCount int

	if _, err := fmt.Scan(&departmentsCount); err != nil {
		return
	}

	for range departmentsCount {
		var employeesCount int

		if _, err := fmt.Scan(&employeesCount); err != nil {
			return
		}

		manager := NewTemperatureManager()

		for range employeesCount {
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
