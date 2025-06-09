package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	s := strings.Split(data, ",")

	if len(s) > 3 || len(s) < 3 {
		err := errors.New("Неверный формат")
		return 0, "", time.Duration(0), err
	}

	trainintgType := s[1]

	stepsCount, err1 := strconv.Atoi(s[0])
	if err1 != nil {
		return 0, "", time.Duration(0), err1
	}
	if stepsCount <= 0 {
		err1 = errors.New("неверные шаги")
		return 0, "", time.Duration(0), err1
	}

	timeForRun, err2 := time.ParseDuration(s[2])
	if err2 != nil {
		return 0, "", time.Duration(0), err2
	}
	if timeForRun.Hours() <= 0 {
		err2 = errors.New("неверная продолжительность")
		return 0, "", time.Duration(0), err2
	}

	return stepsCount, trainintgType, timeForRun, nil
}

func distance(steps int, height float64) float64 {
	lenStepForHeight := height * stepLengthCoefficient

	distance := (float64(steps) * lenStepForHeight) / mInKm

	return distance
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	speed := distance(steps, height) / duration.Hours()

	return speed
}

func TrainingInfo(data string, weight, height float64) (string, error) {

	var info string

	steps, trainingType, timeForRun, err := parseTraining(data)

	switch trainingType {
	case "Ходьба":
		distanceForWalking := distance(steps, height)
		speedForWalking := meanSpeed(steps, height, timeForRun)
		caloriesForWalking, err := WalkingSpentCalories(steps, weight, height, timeForRun)
		if err != nil {
			return "", nil
		}

		info = fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, timeForRun.Hours(), distanceForWalking, speedForWalking, caloriesForWalking)
		return info, nil

	case "Бег":
		distanceForRun := distance(steps, height)
		speedForRun := meanSpeed(steps, height, timeForRun)
		caloriesForRun, err := RunningSpentCalories(steps, weight, height, timeForRun)
		if err != nil {
			return "", nil
		}

		info = fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, timeForRun.Hours(), distanceForRun, speedForRun, caloriesForRun)
		return info, nil
	default:
		err = errors.New("неизвестный тип тренировки")
		return "", err
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		err := errors.New("Количество шагов равно 0 или отрицательно")
		return 0, err
	}
	if weight <= 0 {
		err := errors.New("Вес равен 0 или отрицательно")
		return 0, err
	}
	if duration <= 0 {
		err := errors.New("Продолжительность равна 0 или отрицательно")
		return 0, err
	}

	speedForSpent := meanSpeed(steps, height, duration)

	calories := (weight * speedForSpent * duration.Minutes()) / minInH

	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		err := errors.New("Количество шагов равно 0 или отрицательно")
		return 0, err
	}
	if weight <= 0 {
		err := errors.New("Вес равен 0 или отрицательно")
		return 0, err
	}
	if height <= 0 {
		err := errors.New("Рост равен 0 или отрицательно")
		return 0, err
	}
	if duration <= 0 {
		err := errors.New("Продолжительность равна 0 или отрицательно")
		return 0, err
	}

	speedForSpent := meanSpeed(steps, height, duration)

	calories := ((weight * speedForSpent * duration.Minutes()) / minInH) * walkingCaloriesCoefficient

	return calories, nil
}
