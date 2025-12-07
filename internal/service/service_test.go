package service

import (
	"context"
	"testing"
	"time"

	"github.com/Vladimirmoscow84/Events_Calendar/internal/model"
	"github.com/Vladimirmoscow84/Events_Calendar/internal/storage/inmemory"
)

func setupService(t *testing.T) *Service {
	memStore := inmemory.New()
	svc, err := New(memStore, memStore)
	if err != nil {
		t.Fatalf("failed to create service: %v", err)
	}
	return svc
}

func TestCreateUpdateDeleteEvent(t *testing.T) {
	svc := setupService(t)
	ctx := context.Background()

	event := &model.Event{
		UserID: 1,
		Title:  "Test Event",
		Notice: "Notice",
		Date:   time.Now(),
	}

	eventID, err := svc.CreateEvent(ctx, event)
	if err != nil {
		t.Fatalf("CreateEvent failed: %v", err)
	}

	if eventID == 0 {
		t.Fatal("expected eventID > 0")
	}

	event.Title = "Updated Event"
	event.EventID = eventID
	if err := svc.UpdateEvent(ctx, event); err != nil {
		t.Fatalf("UpdateEvent failed: %v", err)
	}

	if err := svc.DeleteEvent(ctx, eventID); err != nil {
		t.Fatalf("DeleteEvent failed: %v", err)
	}

	events, _ := svc.EventsForDay(ctx, event.UserID, event.Date)
	if len(events) != 0 {
		t.Fatalf("expected 0 events after deletion, got %d", len(events))
	}
}

func TestEventsForDayWeekMonth(t *testing.T) {
	svc := setupService(t)
	ctx := context.Background()

	now := time.Now()
	userID := 1

	events := []*model.Event{
		{UserID: userID, Title: "Day Event", Notice: "", Date: now},
		{UserID: userID, Title: "Week Event", Notice: "", Date: now.AddDate(0, 0, 2)},
		{UserID: userID, Title: "Month Event", Notice: "", Date: now.AddDate(0, 0, 10)},
	}

	for _, e := range events {
		if _, err := svc.CreateEvent(ctx, e); err != nil {
			t.Fatalf("CreateEvent failed: %v", err)
		}
	}

	dayEvents, err := svc.EventsForDay(ctx, userID, now)
	if err != nil {
		t.Fatalf("EventsForDay failed: %v", err)
	}
	if len(dayEvents) != 1 {
		t.Fatalf("expected 1 event for day, got %d", len(dayEvents))
	}

	weekEvents, err := svc.EventsForWeek(ctx, userID, now)
	if err != nil {
		t.Fatalf("EventsForWeek failed: %v", err)
	}
	if len(weekEvents) != 2 {
		t.Fatalf("expected 2 events for week, got %d", len(weekEvents))
	}

	monthEvents, err := svc.EventsForMonth(ctx, userID, now)
	if err != nil {
		t.Fatalf("EventsForMonth failed: %v", err)
	}
	if len(monthEvents) != 3 {
		t.Fatalf("expected 3 events for month, got %d", len(monthEvents))
	}
}
