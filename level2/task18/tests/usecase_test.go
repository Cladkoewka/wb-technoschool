package tests

import (
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
