package usecase

import (
	"time"

	"github.com/Cladkoewka/wb-technoschool/level2/task18/internal/entity"
)

type Repository interface {
	Create(e entity.Event) error
	Update(e entity.Event) error
	Delete(userID int, date time.Time) error
	List(userID int, from, to time.Time) []entity.Event
}

type CalendarUsecase struct {
	repo Repository
}

func NewCalendarUsecase(r Repository) *CalendarUsecase {
	return &CalendarUsecase{repo: r}
}

func (uc *CalendarUsecase) CreateEvent(e entity.Event) error {
	return uc.repo.Create(e)
}

func (uc *CalendarUsecase) UpdateEvent(e entity.Event) error {
	return uc.repo.Update(e)
}

func (uc *CalendarUsecase) DeleteEvent(userID int, date time.Time) error {
	return uc.repo.Delete(userID, date)
}

func (uc *CalendarUsecase) EventsForDay(userID int, date time.Time) []entity.Event {
	from := date.Truncate(24 * time.Hour)
	to := from.Add(24 * time.Hour)
	return uc.repo.List(userID, from, to)
}

func (uc *CalendarUsecase) EventsForWeek(userID int, date time.Time) []entity.Event {
	from := date.Truncate(24 * time.Hour)
	to := from.AddDate(0, 0, 7)
	return uc.repo.List(userID, from, to)
}

func (uc *CalendarUsecase) EventsForMonth(userID int, date time.Time) []entity.Event {
	from := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)
	to := from.AddDate(0, 1, 0)
	return uc.repo.List(userID, from, to)
}
