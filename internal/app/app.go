package app

import (
	"context"
	"log"
	"time"

	"github.com/Vladimirmoscow84/Events_Calendar/internal/handlers"
	"github.com/Vladimirmoscow84/Events_Calendar/internal/middleware"
	"github.com/Vladimirmoscow84/Events_Calendar/internal/service"
	"github.com/Vladimirmoscow84/Events_Calendar/internal/storage"
	"github.com/Vladimirmoscow84/Events_Calendar/internal/storage/inmemory"
	"github.com/Vladimirmoscow84/Events_Calendar/internal/worker"
	"github.com/wb-go/wbf/config"
	"github.com/wb-go/wbf/ginext"
)

func Run() {

	cfg := config.New()
	err := cfg.LoadEnvFiles(".env")
	if err != nil {
		log.Fatalf("[app] error of loading cfg: %v", err)
	}
	cfg.EnableEnv("")
	serverAddr := cfg.GetString("SERVER_ADDRESS")

	memStore := inmemory.New()

	store := storage.New(memStore)
	log.Println("[app] store initialized successfully")

	svc, err := service.New(store, store)
	if err != nil {
		log.Fatalf("[app] failed to create service")
	}
	log.Println("[app] service initialized successfully")

	logCh := make(chan middleware.Logger, 50)
	middleware.RunLogger(logCh)
	engine := ginext.New("release")
	engine.Use(middleware.LoggerMiddleware(logCh))

	router, err := handlers.New(engine, svc, svc)
	if err != nil {
		log.Fatalf("[app] failed to create router: %v", err)
	}
	log.Println("[app] router initialized successfully")

	router.Routes()

	ctx := context.Background()

	worker.RunCleaner(ctx, worker.CleanerConfig{
		Interval: 1 * time.Minute,
	}, store)

	log.Printf("[app] server running on %s", serverAddr)
	err = engine.Run(serverAddr)
	if err != nil {
		log.Fatalf("[app] server failed: %v", err)
	}

}
