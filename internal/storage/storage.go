package storage

import (
	"context"
	"time"

	"github.com/Vladimirmoscow84/Events_Calendar/internal/model"
)

type store interface {
	CreateEvent(ctx context.Context, event *model.Event) (int, error)
	UpdateEvent(ctx context.Context, event *model.Event) error
	DeleteEvent(ctx context.Context, eventID int) error
	EventsForDay(ctx context.Context, userID int, day time.Time) ([]model.Event, error)
	EventsForWeek(ctx context.Context, userID int, day time.Time) ([]model.Event, error)
	EventsForMonth(ctx context.Context, userID int, day time.Time) ([]model.Event, error)
	DeleteOld(ctx context.Context, before time.Time) error
}

type Storage struct {
	store store
}

func New(s store) *Storage {
	return &Storage{
		store: s,
	}
}

func (s *Storage) CreateEvent(ctx context.Context, event *model.Event) (int, error) {
	return s.store.CreateEvent(ctx, event)
}

func (s *Storage) UpdateEvent(ctx context.Context, event *model.Event) error {
	return s.store.UpdateEvent(ctx, event)
}

func (s *Storage) DeleteEvent(ctx context.Context, eventID int) error {
	return s.store.DeleteEvent(ctx, eventID)
}

func (s *Storage) EventsForDay(ctx context.Context, userID int, day time.Time) ([]model.Event, error) {
	return s.store.EventsForDay(ctx, userID, day)
}

func (s *Storage) EventsForWeek(ctx context.Context, userID int, day time.Time) ([]model.Event, error) {
	return s.store.EventsForWeek(ctx, userID, day)
}

func (s *Storage) EventsForMonth(ctx context.Context, userID int, day time.Time) ([]model.Event, error) {
	return s.store.EventsForMonth(ctx, userID, day)
}

func (s *Storage) DeleteOld(ctx context.Context, before time.Time) error {
	return s.store.DeleteOld(ctx, before)
}
