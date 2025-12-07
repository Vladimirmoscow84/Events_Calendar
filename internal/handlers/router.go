package handlers

import (
	"context"
	"errors"
	"time"

	"github.com/Vladimirmoscow84/Events_Calendar/internal/model"
	"github.com/wb-go/wbf/ginext"
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

type Router struct {
	Engine      *ginext.Engine
	Constructor eventsConstructor
	Getter      eventsGetter
}

func New(e *ginext.Engine, ctr eventsConstructor, gtr eventsGetter) (*Router, error) {
	if e == nil || ctr == nil || gtr == nil {
		return nil, errors.New("[handlers] invalid Router parametrs")
	}
	return &Router{
		Engine:      e,
		Constructor: ctr,
		Getter:      gtr,
	}, nil
}

func (r *Router) Routes() {
	r.Engine.POST("/create_event", r.CreateEventHandler)
	r.Engine.POST("/update_event", r.UpdateEventHandler)
	r.Engine.POST("/delete_event", r.DeleteEventHandler)
	r.Engine.GET("/events_for_day", r.EventsForDayHandler)
	r.Engine.GET("/events_for_week", r.EventsForWeekHandler)
	r.Engine.GET("/events_for_month", r.EventsForMonthHandler)
}
