package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("Ошибка! Количество входных данных не соответствует заявленному.")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return errors.New("Ошибка! Неверный формат шагов.")
	}

	if steps <= 0 {
		return errors.New("Ошибка! Количество шагов не является положительным числом.")
	}

	delta, err := time.ParseDuration(parts[1])
	if err != nil {
		return errors.New("Ошибка! Неверный формат продолжительности.")
	}

	if delta <= 0 {
		return errors.New("Ошибка! Продолжительность должна быть больше нуля.")
	}

	ds.Steps = steps
	ds.Duration = delta
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calorie, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)

	if err != nil {
		return "", err
	}

	out := fmt.Sprintf("Количество шагов: %d.\n"+
		"Дистанция составила %.2f км.\n"+
		"Вы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calorie,
	)
	return out, nil
}
