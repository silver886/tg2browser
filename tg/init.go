package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/silver886/logger"
)

var (
	log         *logger.Entry
	telegramBot *tgbotapi.BotAPI
)

// Init set required resources for telegram.
func Init(l *logger.Logger, token string, debug bool) error {
	// Set module logger
	log = l.WithField("module", "tg")

	// Set logger prefix
	logInit := log.WithField("prefix", "init")

	// Set logger for Telegram bot
	logInit.Debugln("Set logger for Telegram bot . . .")
	tgbotapi.SetLogger(log.WithField("prefix", "telegramBot"))

	// Connect Telegram API
	logInit.Debugln("Connect Telegram API . . .")
	if bot, err := tgbotapi.NewBotAPI(token); err == nil {
		telegramBot = bot
	} else {
		return err
	}

	// Upgrade debug mode
	if debug {
		logInit.Debugln("Upgrade Telegram bot . . .")
		telegramBot.Debug = true
	}

	return nil
}
