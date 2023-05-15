package Repositories

import (
	"context"
	"github.com/fiatjaf/relayer/v2/storage/postgresql"
	"github.com/nbd-wtf/go-nostr"
	"time"
)

type StorageWrapper struct {
	backend *postgresql.PostgresBackend
}

func (s *StorageWrapper) Init() error {
	return s.backend.Init()
}

func (s *StorageWrapper) QueryEvents(filter *nostr.Filter) ([]nostr.Event, error) {
	ctx := context.Background() // 或者其他 pass 到這裡的 context
	ch, err := s.backend.QueryEvents(ctx, filter)
	if err != nil {
		return nil, err
	}

	var events []nostr.Event
	for evt := range ch {
		events = append(events, *evt)
	}

	return events, nil
}

func (s *StorageWrapper) DeleteOldEvents() error {
	ctx := context.Background()
	_, err := s.backend.DB.ExecContext(ctx, `DELETE FROM event WHERE created_at < $1`, time.Now().AddDate(0, -3, 0).Unix()) // 3 months
	return err
}

func (s *StorageWrapper) DeleteEvent(id string, PublishKey string) error {
	ctx := context.Background()
	return s.backend.DeleteEvent(ctx, id, PublishKey)
}

func (s *StorageWrapper) SaveEvent(event *nostr.Event) error {
	ctx := context.Background()
	return s.backend.SaveEvent(ctx, event)
}
