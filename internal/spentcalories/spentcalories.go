package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	//lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	parsed := strings.Split(data, ",")

	if len(parsed) != 3 {
		return 0, "", 0, errors.New("Длина слайса не равна 3")
	}

	steps, err := strconv.Atoi(parsed[0])
	if err != nil {
		return 0, "", 0, errors.New("Не удалось обработать данные в INT")
	}

	if steps < 1 {
		return 0, "", 0, errors.New("Количество шагов должно быть больше 0")
	}

	duration, err := time.ParseDuration(parsed[2])
	if err != nil {
		return 0, "", 0, errors.New("Не удалось обработать данные в Time.duration")
	}

	if duration <= 0 {
		return 0, "", 0, errors.New("Продолжительность должна быть больше 0")
	}

	if strings.TrimSpace(parsed[1]) == "" {
		return 0, "", 0, errors.New("Вид активности отсутствует")
	}

	return steps, parsed[1], duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	return (height * stepLengthCoefficient) * float64(steps) / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}

	distance := distance(steps, height)

	if distance == 0 {
		return 0
	}

	return distance / duration.Hours()

}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, typeTraining, duration, err := parseTraining(data)

	if err != nil {
		log.Println(err)
		return "", err
	}

	var caloriesBurned float64

	switch typeTraining {
	case "Бег":
		caloriesBurned, err = RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println(err)
			return "", err
		}

	case "Ходьба":
		caloriesBurned, err = WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
	default:
		log.Println("неизвестный тип тренировки")
		return "", errors.New("неизвестный тип тренировки")
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", typeTraining, duration.Hours(), distance(steps, height), meanSpeed(steps, height, duration), caloriesBurned), nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps < 1 {
		return 0, errors.New("Количество шагов должно быть больше 0")
	}

	if duration <= 0 {
		return 0, errors.New("Продолжительность должна быть больше 0")
	}

	if weight <= 0 || height <= 0 {
		return 0, errors.New("Вес или рост должны быть больше 0")
	}

	avgSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	return (weight * avgSpeed * durationInMinutes) / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps < 1 {
		return 0, errors.New("Количество шагов должно быть больше 0")
	}

	if duration <= 0 {
		return 0, errors.New("Продолжительность должна быть больше 0")
	}

	if weight <= 0 || height <= 0 {
		return 0, errors.New("Вес или рост должны быть больше 0")
	}

	avgSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	return (weight * avgSpeed * durationInMinutes) / minInH * walkingCaloriesCoefficient, nil
}

