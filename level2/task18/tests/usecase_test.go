package tests

import (
	"strconv"
	"testing"
	"time"

	"github.com/Cladkoewka/wb-technoschool/level2/task18/internal/entity"
	"github.com/Cladkoewka/wb-technoschool/level2/task18/internal/repository"
	"github.com/Cladkoewka/wb-technoschool/level2/task18/internal/usecase"
)

func TestCreateAndGetEvent(t *testing.T) {
	repo := repository.NewMemoryRepository()
	uc := usecase.NewCalendarUsecase(repo)

	date := time.Now().Truncate(24 * time.Hour)
	e := entity.Event{UserID: 1, Date: date, Text: "Test event"}

	if err := uc.CreateEvent(e); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	events := uc.EventsForDay(1, date)
	if len(events) != 1 || events[0].Text != "Test event" {
		t.Fatalf("expected 1 event, got %+v", events)
	}
}

func TestUpdateEvent(t *testing.T) {
	repo := repository.NewMemoryRepository()
	uc := usecase.NewCalendarUsecase(repo)

	date := time.Now().Truncate(24 * time.Hour)
	e := entity.Event{UserID: 1, Date: date, Text: "Original"}
	_ = uc.CreateEvent(e)

	e.Text = "Updated"
	if err := uc.UpdateEvent(e); err != nil {
		t.Fatalf("unexpected error on update: %v", err)
	}

	events := uc.EventsForDay(1, date)
	if events[0].Text != "Updated" {
		t.Fatalf("expected updated text, got %+v", events[0].Text)
	}
}

func TestDeleteEvent(t *testing.T) {
	repo := repository.NewMemoryRepository()
	uc := usecase.NewCalendarUsecase(repo)

	date := time.Now().Truncate(24 * time.Hour)
	e := entity.Event{UserID: 1, Date: date, Text: "To be deleted"}
	_ = uc.CreateEvent(e)

	if err := uc.DeleteEvent(1, date); err != nil {
		t.Fatalf("unexpected error on delete: %v", err)
	}

	events := uc.EventsForDay(1, date)
	if len(events) != 0 {
		t.Fatalf("expected 0 events, got %+v", events)
	}
}

func TestEventsForWeekAndMonth(t *testing.T) {
	repo := repository.NewMemoryRepository()
	uc := usecase.NewCalendarUsecase(repo)

	baseDate := time.Date(2025, 9, 1, 0, 0, 0, 0, time.UTC)
	for i := 0; i < 10; i++ {
		e := entity.Event{UserID: 1, Date: baseDate.AddDate(0, 0, i), Text: "Event" + strconv.Itoa(i)}
		_ = uc.CreateEvent(e)
	}

	weekEvents := uc.EventsForWeek(1, baseDate)
	if len(weekEvents) != 7 {
		t.Fatalf("expected 7 events for week, got %d", len(weekEvents))
	}

	monthEvents := uc.EventsForMonth(1, baseDate)
	if len(monthEvents) != 10 {
		t.Fatalf("expected 10 events for month, got %d", len(monthEvents))
	}
}

func TestDeleteNonexistentEvent(t *testing.T) {
	repo := repository.NewMemoryRepository()
	uc := usecase.NewCalendarUsecase(repo)

	date := time.Now().Truncate(24 * time.Hour)
	err := uc.DeleteEvent(1, date)
	if err == nil {
		t.Fatalf("expected error when deleting nonexistent event")
	}
}
