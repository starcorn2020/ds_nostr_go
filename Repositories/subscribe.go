package Repositories

import (
	"context"
	"fmt"
	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip19"
	"time"
)

type Subscribe struct{}

func NewSubscribe() *Subscribe {
	return &Subscribe{}
}

func (rep *Subscribe) SubARelay(keys map[string]string, ConnectURL string) (err error) {
	ctx := context.Background()
	relay, err := nostr.RelayConnect(ctx, ConnectURL)
	if err != nil {
		panic(err)
	}

	npub := keys["npub"]

	var filters nostr.Filters
	if _, v, err := nip19.Decode(npub); err == nil {
		pub := v.(string)
		filters = []nostr.Filter{{
			Kinds:   []int{nostr.KindTextNote},
			Authors: []string{pub},
			Limit:   1,
		}}
	} else {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	sub, err := relay.Subscribe(ctx, filters)
	if err != nil {
		panic(err)
	}

	for ev := range sub.Events {
		// handle returned event.
		// channel will stay open until the ctx is cancelled (in this case, context timeout)
		fmt.Println(ev.ID)
	}

	return nil
}
