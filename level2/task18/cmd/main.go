package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Cladkoewka/wb-technoschool/level2/task18/internal/handler"
	"github.com/Cladkoewka/wb-technoschool/level2/task18/internal/repository"
	"github.com/Cladkoewka/wb-technoschool/level2/task18/internal/usecase"
)

func main() {
	port := os.Getenv("CALENDAR_PORT")
	if port == "" {
		port = "8080"
	}

	repo := repository.NewMemoryRepository()
	uc := usecase.NewCalendarUsecase(repo)
	h := handler.NewHandler(uc)

	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", h.CreateEvent)
	mux.HandleFunc("/update_event", h.UpdateEvent)
	mux.HandleFunc("/delete_event", h.DeleteEvent)
	mux.HandleFunc("/events_for_day", h.EventsForDay)
	mux.HandleFunc("/events_for_week", h.EventsForWeek)
	mux.HandleFunc("/events_for_month", h.EventsForMonth)

	wrapped := handler.LoggingMiddleware(mux)

	log.Printf("Starting server on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, wrapped))
}
