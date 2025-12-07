package app

import (
	"github.com/Vladimirmoscow84/Events_Calendar/internal/service"
	"github.com/Vladimirmoscow84/Events_Calendar/internal/storage"
	"github.com/Vladimirmoscow84/Events_Calendar/internal/storage/inmemory"
)

func Run() {
	memStore := inmemory.New()

	store := storage.New(memStore)

	svc, err := service.New(store, store)
	if err != nil {

	}
}
