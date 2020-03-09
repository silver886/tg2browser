package msg

import (
	"net/url"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Entity(m *tgbotapi.Message) (string, string) {
	for _, e := range *m.Entities {
		entityRaw := string([]rune(m.Text)[e.Offset : e.Offset+e.Length])
		switch e.Type {
		case "url":
			if _, err := url.ParseRequestURI(entityRaw); err != nil {
				entityRaw = "http://" + entityRaw
			}
			return "url", entityRaw
		case "mention":
			return "mention", "https://t.me/" + entityRaw[1:]
		case "email":
			return "email", entityRaw
		case "code":
			return "code", entityRaw
		case "pre":
			return "pre", entityRaw
		case "text_link":
			return "text_link", e.URL
		case "text_mention":
			return "text_mention", e.User.String()
		}
	}
	return "", ""
}
