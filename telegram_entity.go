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
			send(m, entityRaw+" is a valid URL")

			logTelegramEntity.Infoln("Valid URL:", entityRaw)
		case "mention":
			send(m, "https://t.me/"+entityRaw[1:]+" is a valid user name")

			logTelegramEntity.Infoln("User ID:", entityRaw[1:])
		case "email":
			send(m, entityRaw+" is a valid email")

			logTelegramEntity.Infoln("Email:", entityRaw)
		case "code":
			send(m, entityRaw+" is a code line")

			logTelegramEntity.Infoln("Code line:", entityRaw)
		case "pre":
			send(m, entityRaw+" is a code block")

			logTelegramEntity.Infoln("Code block:", entityRaw)
		case "text_link":
			send(m, e.URL+" is a valid URL")

			logTelegramEntity.Debugln("I found a text_link:", e.URL)
		case "text_mention":
			send(m, e.User.String()+" is a valid user")

			logTelegramEntity.Infoln("User:", e.User)
		}
	}
}
