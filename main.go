package main

import (
	"ds_nostr_go/Repositories"
	"ds_nostr_go/Service"
	"os"
)

var (
	address string
	port    string
)

func init() {

	// Initialize the config
	tool := Repositories.Tools{}
	configs, err := tool.ReadKeysJson()
	if err != nil {
		panic(err)
	}

	// Set the config values as environment variables
	for key, value := range configs {
		os.Setenv(key, value)
	}
}

func main() {

	// Get the address and port from the environment
	address = os.Getenv("address")
	port = os.Getenv("port")

	// Initialize the relay service
	relay := Service.NewRelay()
	relay.InitRelay(address, port)

}
