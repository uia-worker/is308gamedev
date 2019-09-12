package main

import (
	"fmt"
	pubnub "github.com/pubnub/go"
)

func main() {
	config := pubnub.NewConfig()
	config.SubscribeKey = "sub-c-7acb63f4-d555-11e9-9e70-eeb3a5fbbe72"
	config.PublishKey = "pub-c-d63a94d6-6450-4608-833f-1360e4f64065"
	pn := pubnub.NewPubNub(config)
	fmt.Println("Velkommen!")
	newLobby("", "", pn)
}