package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Cladkoewka/wb-technoschool/level2/task18/internal/entity"
	"github.com/Cladkoewka/wb-technoschool/level2/task18/internal/usecase"
)

type Handler struct {
	uc *usecase.CalendarUsecase
}

func NewHandler(uc *usecase.CalendarUsecase) *Handler {
	return &Handler{uc: uc}
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var input entity.EventDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeJSON(w, 400, map[string]string{"error": "bad request"})
		return
	}
	date, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		writeJSON(w, 400, map[string]string{"error": "invalid date"})
		return
	}
	e := entity.Event{
		UserID: input.UserID,
		Date:   date,
		Text:   input.Text,
	}
	if err := h.uc.CreateEvent(e); err != nil {
		writeJSON(w, 503, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]string{"result": "created"})
}

func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	var input entity.EventDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeJSON(w, 400, map[string]string{"error": "bad request"})
		return
	}
	date, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		writeJSON(w, 400, map[string]string{"error": "invalid date"})
		return
	}
	e := entity.Event{
		UserID: input.UserID,
		Date:   date,
		Text:   input.Text,
	}
	if err := h.uc.UpdateEvent(e); err != nil {
		writeJSON(w, 503, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]string{"result": "updated"})
}

func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		writeJSON(w, 400, map[string]string{"error": "invalid date"})
		return
	}
	if err := h.uc.DeleteEvent(userID, date); err != nil {
		writeJSON(w, 503, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, 200, map[string]string{"result": "deleted"})
}

func (h *Handler) EventsForDay(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		writeJSON(w, 400, map[string]string{"error": "invalid date"})
		return
	}
	events := h.uc.EventsForDay(userID, date)
	writeJSON(w, 200, events)
}

func (h *Handler) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		writeJSON(w, 400, map[string]string{"error": "invalid date"})
		return
	}
	events := h.uc.EventsForWeek(userID, date)
	writeJSON(w, 200, events)
}

func (h *Handler) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		writeJSON(w, 400, map[string]string{"error": "invalid date"})
		return
	}
	events := h.uc.EventsForMonth(userID, date)
	writeJSON(w, 200, events)
}
