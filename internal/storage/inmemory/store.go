package inmemory

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/Vladimirmoscow84/Events_Calendar/internal/model"
)

var ErrNotFound = errors.New("[store] event not found")

type Store struct {
	mu     sync.RWMutex
	LastID int
	events map[int]model.Event
}

func New() *Store {
	return &Store{
		LastID: 0,
		events: make(map[int]model.Event),
	}
}

// Create создает событие
func (s *Store) Create(ctx context.Context, event *model.Event) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.LastID++

	event.EventID = s.LastID
	s.events[event.EventID] = *event

	return event.EventID, nil
}

// Update обновляет событие
func (s *Store) Update(ctx context.Context, event *model.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.events[event.EventID]
	if !ok {
		return ErrNotFound
	}
	s.events[event.EventID] = *event
	return nil
}

// Delete удаляет событие
func (s *Store) Delete(ctx context.Context, eventID int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.events[eventID]
	if !ok {
		return ErrNotFound
	}

	delete(s.events, eventID)
	return nil

}

// EventsForWeek возвращает события запланированные на конкретный день
func (s *Store) EventsForDay(ctx context.Context, userID int, day time.Time) ([]model.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []model.Event
	y, m, d := day.Date()

	for _, e := range s.events {
		ey, em, ed := e.Date.Date()
		if e.UserID == userID && y == ey && m == em && d == ed {
			result = append(result, e)
		}
	}
	return result, nil
}

// EventsForWeek возвращает события запланированные на неделю
func (s *Store) EventsForWeek(ctx context.Context, userID int, day time.Time) ([]model.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []model.Event
	year, week := day.ISOWeek()

	for _, e := range s.events {
		ey, ew := e.Date.ISOWeek()
		if e.UserID == userID && year == ey && week == ew {
			result = append(result, e)
		}
	}
	return result, nil
}

// EventsForWeek возвращает события запланированные на месяц
func (s *Store) EventsForMonth(ctx context.Context, userID int, day time.Time) ([]model.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	y, m, _ := day.Date()
	var result []model.Event

	for _, e := range s.events {
		ey, em, _ := e.Date.Date()
		if e.UserID == userID && y == ey && m == em {
			result = append(result, e)
		}
	}

	return result, nil
}

// DeleteOld удаляет прошедшие события в зависимости от заданной даты
func (s *Store) DeleteOld(ctx context.Context, before time.Time) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for id, e := range s.events {
		if e.Date.Before(before) {
			delete(s.events, id)
		}
	}

	return nil
}
