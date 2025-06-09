package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	s := strings.Split(data, ",")

	if len(s) > 2 || len(s) < 2 {
		err := errors.New("Неверный формат")
		return 0, time.Duration(0), err
	}

	stepsCount, err1 := strconv.Atoi(s[0])
	if err1 != nil {
		return 0, time.Duration(0), err1
	}
	if stepsCount <= 0 {
		return 0, time.Duration(0), err1
	}

	timeForRun, err2 := time.ParseDuration(s[1])
	if err2 != nil {
		return 0, time.Duration(0), err2
	}
	if timeForRun <= 0 {
		return 0, time.Duration(0), err2
	}

	return stepsCount, timeForRun, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	stepsCount, timeForRun, err := parsePackage(data)
	if err != nil {
		return ""
	}
	if stepsCount <= 0 {
		return ""
	}
	distance := float64(stepsCount) * stepLength
	calories

	info := fmt.Printf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", stepsCount, distance, calories)
	return info
}
