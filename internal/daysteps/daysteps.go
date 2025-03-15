package daysteps

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
// StepLength = 0.65
)

// создайте структуру DaySteps
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (ds *DaySteps) Parse(datastring string) (err error) {
	s := strings.Split(datastring, ",")
	if len(s) != 2 {
		return errors.New("некорректная строка - формат строки должен быть \"678,0h50m\"")
	}
	ds.Steps, err = strconv.Atoi(s[0])
	if err != nil {
		return fmt.Errorf("conversion error: %w", err)
	}
	ds.Duration, err = time.ParseDuration(s[1])
	if err != nil {
		return errors.New("некорректная строка - требуется продолжительность в формате \"3h00m\"")
	}
	return nil
}

// создайте метод ActionInfo()
func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Duration <= 0 {
		return "", errors.New("продолжительность должна быть больше нуля")
	}
	distance := spentenergy.Distance(ds.Steps)
	spentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`
Количество шагов: %v.
Дистанция составила %.2f км.
Вы сожгли %.2f ккал.
`, ds.Steps, distance, spentCalories), nil
}
