package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

type telegramClient struct {
	bot *tgbotapi.BotAPI
}

func NewTelegramClient() *telegramClient {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("PROJECT_MESSAGEDG_TOKEN"))
	if err != nil {
		log.Fatalf("Error initializing telegram %v", err)
	}
	return &telegramClient{
		bot: bot,
	}
}

func (t telegramClient) SendMessage(text string) error {
	chatId := int64(-1001549822450)
	message := tgbotapi.NewMessage(chatId, text)
	_, err := t.bot.Send(message)
	return err
}

func (t telegramClient) GetMessagesAndReply() {
	update := tgbotapi.NewUpdate(0)
	update.Timeout = 60

	updates := t.bot.GetUpdatesChan(update)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID

		message := tgbotapi.NewMessage(chatID, "Sana da selam")
		message.ReplyToMessageID = update.Message.MessageID // Hangi mesaja cevap veriyor
		t.bot.Send(message)
	}
}
