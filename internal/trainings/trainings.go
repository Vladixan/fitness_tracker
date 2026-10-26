package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("Ошибка! Количество входных данных не соответствует заявленному.")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return errors.New("Ошибка! Неверный формат шагов.")
	}

	if steps <= 0 {
		return errors.New("Ошибка! Количество шагов не является положительным числом.")
	}

	delta, err := time.ParseDuration(parts[2])
	if err != nil {
		return errors.New("Ошибка! Неверный формат продолжительности.")
	}

	if delta <= 0 {
		return errors.New("Ошибка! Продолжительность должна быть больше нуля.")
	}

	t.Steps = steps
	t.TrainingType = parts[1]
	t.Duration = delta
	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	var calorie float64
	var err error

	switch t.TrainingType {
	case "Ходьба":
		calorie, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		calorie, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("Ошибка! Неизвестный тип тренировки.")
	}

	if err != nil {
		return "", err
	}

	out := fmt.Sprintf("Тип тренировки: %s\n"+
		"Длительность: %.2f ч.\n"+
		"Дистанция: %.2f км.\n"+
		"Скорость: %.2f км/ч\n"+
		"Сожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		speed,
		calorie,
	)
	return out, nil
}
