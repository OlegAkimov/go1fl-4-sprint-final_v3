package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	parsed := strings.Split(data, ",")

	if len(parsed) != 2 {
		return 0, 0, errors.New("Длина слайса не равна 2")
	}

	steps, err := strconv.Atoi(parsed[0])
	if err != nil {
		return 0, 0, errors.New("Не удалось обработать данные в INT")
	}

	if steps < 1 {
		return 0, 0, errors.New("Количество шагов должно быть больше 0")
	}

	duration, err := time.ParseDuration(parsed[1])
	if err != nil {
		return 0, 0, errors.New("Не удалось обработать данные в Time.duration")
	}

	if duration <= 0 {
		return 0, 0, errors.New("Продолжительность должна быть больше 0")
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		return ""
	}

	distance := float64(steps) * stepLength
	kilometers := distance / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, kilometers, calories)
}