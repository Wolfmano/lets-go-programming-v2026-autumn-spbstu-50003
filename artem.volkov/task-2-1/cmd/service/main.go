package main

import (
	"fmt"
	"os"
)

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
		fmt.Fprintln(os.Stderr, fmt.Errorf("ошибка чтения количества отделов: %w", err))

		return
	}

	for range departmentsCount {
		var employeesCount int

		if _, err := fmt.Scan(&employeesCount); err != nil {
			fmt.Fprintln(os.Stderr, fmt.Errorf("ошибка чтения количества сотрудников: %w", err))

			return
		}

		manager := NewTemperatureManager()

		for range employeesCount {
			var sign string

			var temp int

			if _, err := fmt.Scan(&sign); err != nil {
				fmt.Fprintln(os.Stderr, fmt.Errorf("ошибка чтения знаков: %w", err))

				return
			}

			if _, err := fmt.Scan(&temp); err != nil {
				fmt.Fprintln(os.Stderr, fmt.Errorf("ошибка чтения температуры: %w", err))

				return
			}

			manager.UpdateTemperature(sign, temp)

			fmt.Println(manager.GetTemperature())
		}
	}
}
