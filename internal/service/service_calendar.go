package service

import (
	"context"
	"errors"
	"time"

	"github.com/Vladimirmoscow84/Events_Calendar/internal/model"
)

var (
	ErrInvalidUserID  = errors.New("invalid user_id")
	ErrInvalidTitle   = errors.New("title cannot be empty")
	ErrInvalidDate    = errors.New("invalid date")
	ErrInvalidEventID = errors.New("invalid event_id")
)

func (s *Service) CreateEvent(ctx context.Context, event *model.Event) (int, error) {
	if event.UserID <= 0 {
		return 0, ErrInvalidUserID
	}
	if event.Title == "" {
		return 0, ErrInvalidTitle
	}
	if event.Date.IsZero() {
		return 0, ErrInvalidDate
	}

	return s.constructor.CreateEvent(ctx, event)
}

func (s *Service) UpdateEvent(ctx context.Context, event *model.Event) error {
	if event.EventID <= 0 {
		return ErrInvalidEventID
	}
	if event.UserID <= 0 {
		return ErrInvalidUserID
	}
	if event.Title == "" {
		return ErrInvalidTitle
	}
	if event.Date.IsZero() {
		return ErrInvalidDate
	}

	return s.constructor.UpdateEvent(ctx, event)
}

func (s *Service) DeleteEvent(ctx context.Context, eventID int) error {
	if eventID <= 0 {
		return ErrInvalidEventID
	}
	return s.constructor.DeleteEvent(ctx, eventID)
}

func (s *Service) EventsForDay(ctx context.Context, userID int, day time.Time) ([]model.Event, error) {
	if userID <= 0 {
		return nil, ErrInvalidUserID
	}
	return s.getter.EventsForDay(ctx, userID, day)
}

func (s *Service) EventsForWeek(ctx context.Context, userID int, day time.Time) ([]model.Event, error) {
	if userID <= 0 {
		return nil, ErrInvalidUserID
	}
	return s.getter.EventsForWeek(ctx, userID, day)
}

func (s *Service) EventsForMonth(ctx context.Context, userID int, day time.Time) ([]model.Event, error) {
	if userID <= 0 {
		return nil, ErrInvalidUserID
	}
	return s.getter.EventsForMonth(ctx, userID, day)
}
