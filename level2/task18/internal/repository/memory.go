package repository

import (
	"errors"
	"sync"
	"time"

	"github.com/Cladkoewka/wb-technoschool/level2/task18/internal/entity"
)

type MemoryRepository struct {
	mu     sync.RWMutex
	events map[int][]entity.Event
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{events: make(map[int][]entity.Event)}
}

func (r *MemoryRepository) Create(e entity.Event) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.events[e.UserID] = append(r.events[e.UserID], e)
	return nil
}

func (r *MemoryRepository) Update(e entity.Event) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	list := r.events[e.UserID]
	for i, ev := range list {
		if ev.Date.Equal(e.Date) {
			list[i] = e
			r.events[e.UserID] = list
			return nil
		}
	}
	return errors.New("event not found")
}

func (r *MemoryRepository) Delete(userID int, date time.Time) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	list := r.events[userID]
	for i, ev := range list {
		if ev.Date.Equal(date) {
			r.events[userID] = append(list[:i], list[i+1:]...)
			return nil
		}
	}
	return errors.New("event not found")
}

func (r *MemoryRepository) List(userID int, from, to time.Time) []entity.Event {
	r.mu.RLock()
	defer r.mu.RUnlock()
	res := []entity.Event{}
	for _, ev := range r.events[userID] {
		if (ev.Date.Equal(from) || ev.Date.After(from)) && ev.Date.Before(to) {
			res = append(res, ev)
		}
	}
	return res
}
