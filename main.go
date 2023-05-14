package main

import (
	"context"
	"fmt"
	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip19"
	"github.com/nbd-wtf/go-nostr/nip42"
	"time"
)

func main() {
	url := "ws://localhost:7447"

	// Once the connection is initiated the server will send "AUTH" with the challenge string.
	relay, err := nostr.RelayConnect(context.Background(), url)
	if err != nil {
		panic(err)
	}

	// Initialize test user.
	sk := nostr.GeneratePrivateKey()
	pub, _ := nostr.GetPublicKey(sk)
	npub, _ := nip19.EncodePublicKey(pub)

	// Relay.Challenges channel will receive the "AUTH" command.
	challenge := <-relay.Challenges

	// Create the auth event to send back.
	// The user will be authenticated as pub.
	event := nip42.CreateUnsignedAuthEvent(challenge, pub, url)
	event.Sign(sk)

	// Set-up context with 3 second time out.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Send the event by calling relay.Auth.
	// Returned status is either success, fail, or sent (if no reply given in the 3 second timeout).
	auth_status, err := relay.Auth(ctx, event)

	fmt.Printf("authenticated as %s: %s\n", npub, auth_status)
}
