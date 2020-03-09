package tg

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func send(m *tgbotapi.MessageConfig) {
	m.Text += ", send mechanism still under developing . . ."
	telegramBot.Send(m)
}
