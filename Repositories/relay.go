package Repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/fiatjaf/relayer"
	"github.com/fiatjaf/relayer/v2/storage/postgresql"
	"github.com/kelseyhightower/envconfig"
	"github.com/nbd-wtf/go-nostr"
)

type Relay struct {
	PostgresDatabase string `envconfig:"POSTGRESQL_DATABASE"`
	storage          *StorageWrapper
	ctx              context.Context
}

func (r *Relay) Name() string {
	return "BasicRelay"
}

func (r *Relay) Storage() relayer.Storage {
	return r.storage
}

func (r *Relay) OnInitialized(s *relayer.Server) {
	log.Printf("Server initialized: %v", s)
}

func (r *Relay) Init() error {
	return r.InitWithContext(r.ctx)
}

func (r *Relay) InitWithContext(ctx context.Context) error {
	err := envconfig.Process("", r)
	if err != nil {
		return fmt.Errorf("couldn't process envconfig: %w", err)
	}

	r.storage = &StorageWrapper{
		backend: &postgresql.PostgresBackend{
			DatabaseURL: r.PostgresDatabase,
		},
	}

	if err := r.storage.backend.Init(); err != nil {
		return fmt.Errorf("couldn't initialize storage: %w", err)
	}

	go func() {
		ticker := time.NewTicker(60 * time.Minute)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				// Assuming that your PostgresBackend has a method called DeleteOldEvents.
				if err := r.storage.DeleteOldEvents(); err != nil {
					log.Printf("could not delete old events: %v", err)
				}
			}
		}
	}()

	return nil
}

func (r *Relay) AcceptEvent(evt *nostr.Event) bool {
	jsonb, _ := json.Marshal(evt)
	return len(jsonb) <= 10000
}
