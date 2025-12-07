package service

import (
	"context"
	"errors"
	"time"

	"github.com/Vladimirmoscow84/Events_Calendar/internal/model"
)

type eventsConstructor interface {
	CreateEvent(ctx context.Context, event *model.Event) (int, error)
	UpdateEvent(ctx context.Context, event *model.Event) error
	DeleteEvent(ctx context.Context, eventID int) error
}

type eventsGetter interface {
	EventsForDay(ctx context.Context, userID int, day time.Time) ([]model.Event, error)
	EventsForWeek(ctx context.Context, userID int, day time.Time) ([]model.Event, error)
	EventsForMonth(ctx context.Context, userID int, day time.Time) ([]model.Event, error)
}

type Service struct {
	constructor eventsConstructor
	getter      eventsGetter
}

func New(eCtr eventsConstructor, eGtr eventsGetter) (*Service, error) {
	if eCtr == nil || eGtr == nil {
		return nil, errors.New("[service] storage is nil")
	}
	return &Service{
		constructor: eCtr,
		getter:      eGtr,
	}, nil
}
