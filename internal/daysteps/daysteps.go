package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-4-sprint-final/internal/spentcalories"
)

var (
	StepLength = 0.65 // длина шага в метрах
)

func parsePackage(data string) (int, time.Duration, error) {
	// ваш код ниже
	slice := strings.Split(data, ",")
	if len(slice) == 2 {
		steps, err1 := strconv.Atoi(slice[0])
		if err1 != nil {
			return 0, 0, err1
		}
		duration, err2 := time.ParseDuration(slice[1])
		if err2 != nil {
			return 0, 0, err2
		}
		return steps, duration, nil
	} else {
		return 0, 0, nil
	}
}

// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.
func DayActionInfo(data string, weight, height float64) string {
	// ваш код ниже
	steps, duration, err := parsePackage(data)
	if err != nil {
		return ""
	}
	if steps <= 0 {
		return ""
	}
	dist := float64(steps) * StepLength / 1000
	calories := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	return fmt.Sprintf("Количество шагов: %d. \n Дистанция составила %.2f км. \n Вы сожгли %.2f ккал. \n", steps, dist, calories)
}
