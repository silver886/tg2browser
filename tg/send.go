package tg

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func send(m *tgbotapi.Message, text string) {
	msg := tgbotapi.NewMessage(m.Chat.ID, text+", send mechanism still under developing . . .")
	msg.ReplyToMessageID = m.MessageID
	telegramBot.Send(msg)

}
