package Repositories

import (
	"fmt"

	nostr "github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip19"
)

type Keys struct{}

func NewKeys() *Keys {
	return &Keys{}
}

func (rep *Keys) GenKey() (keys map[string]string, err error) {
	// Gen Public & Private Key
	sk := nostr.GeneratePrivateKey()
	pk, err := nostr.GetPublicKey(sk)
	if err != nil {
		return nil, err
	}
	nsec, err := nip19.EncodePrivateKey(sk)
	if err != nil {
		return nil, err
	}
	npub, err := nip19.EncodePublicKey(pk)
	if err != nil {
		return nil, err
	}

	keys = make(map[string]string)
	keys["sk"] = sk
	keys["pk"] = pk
	keys["nsec"] = nsec
	keys["npub"] = npub

	fmt.Printf("sk: %s\n", sk)
	fmt.Println(nsec)
	fmt.Printf("pk: %s\n", pk)
	fmt.Println(npub)

	return keys, nil
}
