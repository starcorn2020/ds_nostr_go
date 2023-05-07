package main

import (
	"ds_nostr_go/Repositories"
	"fmt"
	"log"
)

func main() {

	// 拿 Keys
	tools := Repositories.NewTools()
	keys, err := tools.ReadKeysJson()
	if err != nil {
		log.Fatal(err)
	}

	// 推送內容
	urls := []string{"wss://relay.nekolicio.us"}
	text := "Hello World"
	publish := Repositories.NewPublish()
	err = publish.PublishRelays(keys, urls, text)
	if err == nil {
		fmt.Println("Test Pass")
	}
}
