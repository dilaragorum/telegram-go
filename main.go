package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	apiToken := os.Getenv("PROJECT_MESSAGEDG_TOKEN")
	bot, err := tgbotapi.NewBotAPI(apiToken)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	update := tgbotapi.NewUpdate(0)
	update.Timeout = 5

	updates := bot.GetUpdatesChan(update)

	go func() {
		for range time.Tick(5 * time.Second) {
			sendMessageToChatGroup(apiToken)
		}
	}()

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}

func sendMessageToChatGroup(token string) {

	chatId := "-1001549822450"
	text := "Selam"

	urlString := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s",
		token, chatId, text)
	_, err := http.Post(urlString, "text/plain", strings.NewReader(text))

	//Error handler için middleware ekle.
	//Go httpclient'a middleware ekleme ve error mesajlarını third party
	//tool'a gönderme(Senty-ELK)
	if err != nil {
		fmt.Println(err.Error())
	}
}
