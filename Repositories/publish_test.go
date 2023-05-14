package Repositories

import (
	"fmt"
	"log"
	"testing"
)

func TestPublishRelays(t *testing.T) {
	// 拿 Keys
	tools := NewTools()
	keys, err := tools.ReadKeysJson()
	if err != nil {
		log.Fatal(err)
	}

	// 推送內容
	urls := []string{"wss://relay.nekolicio.us"}
	text := "ss"
	publish := NewPublish()
	err = publish.PublishRelays(keys, urls, text)

	// 測試結果
	if err == nil {
		fmt.Println("Test Pass")
	} else {
		t.Errorf("Test Fail")
	}
}
