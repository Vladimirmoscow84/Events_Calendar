package worker

import (
	"context"
	"log"
	"time"

	"github.com/Vladimirmoscow84/Events_Calendar/internal/storage"
)

type CleanerConfig struct {
	Interval time.Duration // интервал запуска
}

func RunCleaner(ctx context.Context, cfg CleanerConfig, store *storage.Storage) {
	go func() {
		ticker := time.NewTicker(cfg.Interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				before := time.Now()
				if err := store.DeleteOld(ctx, before); err != nil {
					log.Printf("[worker] error cleaning old events: %v", err)
				} else {
					log.Printf("[worker] cleaned events before %s", before.Format(time.RFC3339))
				}
			case <-ctx.Done():
				log.Println("[worker] cleaner stopped")
				return
			}
		}
	}()
}
