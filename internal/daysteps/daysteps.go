package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-4-sprint-final/internal/spentcalories"
)

var (
	StepLength  = 0.65 // длина шага в метрах
	ErrProgramm = errors.New("ошибка в ходе выполнения программы")
)

func parsePackage(data string) (int, time.Duration, error) {
	// ваш код ниже
	slice := strings.Split(data, ",")
	if len(slice) == 2 {
		steps, err := strconv.Atoi(slice[0])
		//if err1 != nil {
		//	return 0, 0, err1
		//}
		duration, err := time.ParseDuration(slice[1])
		if err != nil {
			return 0, 0, err
		}
		return steps, duration, nil
	}
	return 0, 0, ErrProgramm

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
		fmt.Println(err)
		return ""
	}
	if steps <= 0 {
		return ""
	}
	dist := float64(steps) * StepLength / 1000
	calories := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	return fmt.Sprintf("Количество шагов: %d. \n Дистанция составила %.2f км. \n Вы сожгли %.2f ккал. \n", steps, dist, calories)
}
