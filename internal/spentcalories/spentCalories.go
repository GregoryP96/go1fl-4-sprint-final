package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep   = 0.65  // средняя длина шага.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.
)

func parseTraining(data string) (int, string, time.Duration, error) {

	arr := strings.Split(data, ",")
	if len(arr) != 3 {
		return 0, "", 0, fmt.Errorf("the length of the slice is not equal to three")
	}

	numOfSteps, err := strconv.Atoi(arr[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("the first value of a slice must be an integer")
	}

	duration, err := time.ParseDuration(arr[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid duration template")
	}
	return numOfSteps, arr[1], duration, nil
}

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий (число шагов при ходьбе и беге).
func distance(steps int) float64 {

	return (float64(steps) * lenStep) / mInKm
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий(число шагов при ходьбе и беге).
// duration time.Duration — длительность тренировки.
func meanSpeed(steps int, duration time.Duration) float64 {

	if duration < 0 {
		return 0
	}
	dist := distance(steps)
	return dist / (duration.Hours())
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
//
// Параметры:
//
// data string - строка с данными.
// weight, height float64 — вес и рост пользователя.
func TrainingInfo(data string, weight, height float64) string {

	steps, typeOfTraining, duration, err := parseTraining(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	dist := distance(steps)
	speed := meanSpeed(steps, duration)
	var spentCalories float64

	switch typeOfTraining {
	case "Ходьба":
		spentCalories = WalkingSpentCalories(steps, weight, height, duration)
	case "Бег":
		spentCalories = RunningSpentCalories(steps, weight, duration)
	default:
		return "неизвестный тип тренировки"
	}
	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		typeOfTraining, duration.Hours(), dist, speed, spentCalories)
	return result
}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// steps int - количество шагов.
// weight float64 — вес пользователя.
// duration time.Duration — длительность тренировки.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {

	avrSpeed := meanSpeed(steps, duration)
	return ((runningCaloriesMeanSpeedMultiplier * avrSpeed) - runningCaloriesMeanSpeedShift) * weight
}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// steps int - количество шагов.
// duration time.Duration — длительность тренировки.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) float64 {

	avrSpeed := meanSpeed(steps, duration)
	return ((walkingCaloriesWeightMultiplier * weight) + (avrSpeed*avrSpeed/height)*walkingSpeedHeightMultiplier) * duration.Hours() * minInH
}
