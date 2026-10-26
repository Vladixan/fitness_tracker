package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	check := (steps <= 0 || duration <= 0 || weight <= 0 || height <= 0)
	if check {
		return 0, errors.New("Ошибка входных данных!")
	}
	calorie := (weight * MeanSpeed(steps, height, duration) * duration.Minutes()) / minInH
	return calorie * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	check := (steps <= 0 || duration <= 0 || weight <= 0 || height <= 0)
	if check {
		return 0, errors.New("Ошибка входных данных!")
	}
	calorie := (weight * MeanSpeed(steps, height, duration) * duration.Minutes()) / minInH
	return calorie, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 || duration <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	return ((height * stepLengthCoefficient) * float64(steps)) / float64(mInKm)
}
