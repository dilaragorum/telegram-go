package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

const projeGroupChatID = int64(-1001549822450)

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
	message := tgbotapi.NewMessage(projeGroupChatID, text)
	_, err := t.bot.Send(message)
	return err
}

func (t telegramClient) GetMessagesAndReply() {
	update := tgbotapi.NewUpdate(0)
	update.Timeout = 5

	updates := t.bot.GetUpdatesChan(update)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID

		message := tgbotapi.NewMessage(chatID, "Sana da selam")
		message.ReplyToMessageID = update.Message.MessageID

		t.bot.Send(message)
	}
}
