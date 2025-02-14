package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"go1fl-4-sprint-final/internal/spentcalories"
)

var (
	StepLength = 0.65 // длина шага в метрах
)

func parsePackage(data string) (int, time.Duration, error) {
	// ваш код ниже
	arr := strings.Split(data, ",")
	if len(arr) != 2 {
		return 0, 0, fmt.Errorf("The length of the slice is not equal to two!")
	}

	numOfSteps, err := strconv.Atoi(arr[0])
	if err != nil {
		return 0, 0, fmt.Errorf("The first value of a slice must be an integer!")
	}

	duration, err := time.ParseDuration(arr[1])
	if err != nil {
		return 0, 0, fmt.Errorf("Invalid duration template!")
	}
	return numOfSteps, duration, nil
}

// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.
func DayActionInfo(data string, weight, height float64) string {
	// ваш код ниже
	numOfSteps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if numOfSteps < 0 {
		return ""
	}
	dist := float64(numOfSteps) * StepLength
	numOfCalories := spentcalories.WalkingSpentCalories(numOfSteps, weight, height, duration)
	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		numOfSteps, dist, numOfCalories)
	return result
}
