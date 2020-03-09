package msg

import (
	"net/url"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// GetText return the text in one message.
func GetText(m *tgbotapi.Message) *tgbotapi.MessageConfig {
	// Fint URL in message text.
	if u, err := url.ParseRequestURI(m.Text); err == nil {
		msg := tgbotapi.NewMessage(m.Chat.ID, "You typed an URL: "+u.String())
		msg.ReplyToMessageID = m.MessageID
		return &msg
	}
	return nil
}
