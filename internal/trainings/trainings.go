package trainings

import (
	"errors"
	"fmt"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

const (
	walking = "Ходьба"
	running = "Бег"
)

// создайте структуру Training
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (t *Training) Parse(datastring string) (err error) {
	s := strings.Split(datastring, ",")
	if len(s) != 3 {
		return errors.New("некорректная строка - формат строки должен быть \"3456,Ходьба,3h00m\"")
	}
	t.Steps, err = strconv.Atoi(s[0])
	if err != nil {
		return errors.New("некорректное значение - необходимо целое число")
	}
	switch s[1] {
	case walking:
		t.TrainingType = s[1]
	case running:
		t.TrainingType = s[1]
	default:
		return fmt.Errorf("некорректный тип тренировки (допустимые значения '%s' или '%s'", walking, running)
	}
	t.Duration, err = time.ParseDuration(s[2])
	if err != nil {
		return errors.New("некорректная строка - требуется продолжительность в формате \"3h00m\"")
	}
	return nil
}

// создайте метод ActionInfo()
func (t Training) ActionInfo() (string, error) {
	if t.Duration <= 0 {
		return "", errors.New("продолжительность должна быть больше нуля")
	}
	distance := spentenergy.Distance(t.Steps)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Duration)
	var spentCalories float64
	var err error
	switch t.TrainingType {
	case walking:
		spentCalories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	case running:
		spentCalories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Duration)
		if err != nil {
			return "", err
		}
	default:
		return "неизвестный тип тренировки", errors.New("неизвестный тип тренировки")
	}
	return fmt.Sprintf(`
Тип тренировки: %v
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f
`, t.TrainingType, t.Duration.Hours(), distance, meanSpeed, spentCalories), nil
}
