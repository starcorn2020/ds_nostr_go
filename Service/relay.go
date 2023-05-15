package Service

import (
	"ds_nostr_go/Repositories"
	"github.com/fiatjaf/relayer"
	"log"
)

type Relay struct{}

func NewRelay() *Relay {
	return &Relay{}
}

func (rep *Relay) InitRelay(address string, port string) {
	r := Repositories.Relay{}
	if err := r.Init(); err != nil {
		log.Fatalf("failed to initialize relay: %v", err)
		return
	}

	fullAddress := address + ":" + port
	server := relayer.NewServer(fullAddress, &r)

	if err := server.Start(); err != nil {
		log.Fatalf("server terminated: %v", err)
	}
}
