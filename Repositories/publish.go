package Repositories

import (
	"context"
	"fmt"
	"github.com/nbd-wtf/go-nostr"
)

type Publish struct{}

func NewPublish() *Publish {
	return &Publish{}
}

func (rep *Publish) PublishRelays(keys map[string]string, ConnectURLs []string, Context string) (err error) {
	pub := keys["pk"]
	ev := nostr.Event{
		PubKey:    pub,
		CreatedAt: nostr.Now(),
		Kind:      nostr.KindTextNote,
		Tags:      nil,
		Content:   Context,
	}

	// calling Sign sets the event ID field and the event Sig field
	sk := keys["sk"]
	ev.Sign(sk)

	// publish the event to two relays
	ctx := context.Background()
	for _, url := range ConnectURLs {
		relay, err := nostr.RelayConnect(ctx, url)
		if err != nil {
			fmt.Println(err)
			continue
		}
		_, err = relay.Publish(ctx, ev)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("published to %s\n", url)

		err = relay.Close()
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return
}
