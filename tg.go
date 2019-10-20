package main

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func tgBot() {
	// Set logger prefix
	logTgBot := log.WithField("prefix", "tgBot")

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_BOT_KEY"))
	if err != nil {
		log.Panic(err)
	}

	tgbotapi.SetLogger(log)

	bot.Debug = true

	logTgBot.Debugf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		logTgBot.Debugf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
