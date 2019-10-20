package main

import (
	"net/url"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func telegramEntity(m *tgbotapi.Message) {
	// Set logger prefix
	logTelegramEntity := log.WithField("prefix", "telegramEntity")

	for i, e := range *m.Entities {
		entityRaw := string([]rune(m.Text)[e.Offset : e.Offset+e.Length])
		logTelegramEntity.Tracef("%v: type: %v, text: %v (%#v)\n", i, e.Type, entityRaw, e)
		switch e.Type {
		case "url":
			if _, err := url.ParseRequestURI(entityRaw); err != nil {
				entityRaw = "http://" + entityRaw
			}

			msg := tgbotapi.NewMessage(m.Chat.ID, entityRaw+" is a valid URL, send mechanism still under developing . . .")
			msg.ReplyToMessageID = m.MessageID
			telegramBot.Send(msg)

			logTelegramEntity.Infoln("Valid URL:", entityRaw)
			// case "mention":
			// 	logTelegramEntity.Debugln("I found a mention:", entityRaw)
			// case "hashtag":
			// 	logTelegramEntity.Debugln("I found a hashtag:", entityRaw)
			// case "email":
			// 	logTelegramEntity.Debugln("I found a email:", entityRaw)
			// case "phone_number":
			// 	logTelegramEntity.Debugln("I found a phone_number:", entityRaw)
			// case "code":
			// 	logTelegramEntity.Debugln("I found a code:", entityRaw)
			// case "pre":
			// 	logTelegramEntity.Debugln("I found a pre:", entityRaw)
			// case "text_link":
			// 	logTelegramEntity.Debugln("I found a text_link:", e.URL)
			// case "text_mention":
			// 	logTelegramEntity.Debugln("I found a text_mention:", e.User)
		}
	}
}
