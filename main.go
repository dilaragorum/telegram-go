package main

import (
	"time"
)

func main() {
	telegramClient := NewTelegramClient()

	go func() {
		for range time.Tick(5 * time.Second) {
			telegramClient.SendMessage("Selam")
		}
	}()

	telegramClient.GetMessagesAndReply()
}
