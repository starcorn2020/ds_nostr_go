package repositories

import (
	"fmt"

	nostr "github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip19"
)

type tools struct {
}

func NewTools() *tools {
	return &tools{}
}

func (rep *tools) GenKey() {
	// 產生 SK, PK

	sk := nostr.GeneratePrivateKey()
	pk, _ := nostr.GetPublicKey(sk)
	nsec, _ := nip19.EncodePrivateKey(sk)
	npub, _ := nip19.EncodePublicKey(pk)

	fmt.Println("sk:", sk)
	fmt.Println("pk:", pk)
	fmt.Println(nsec)
	fmt.Println(npub)
}
