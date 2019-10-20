package main

import (
	"errors"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func telegramInit() error {
	// Set logger prefix
	logTelegramInit := log.WithField("prefix", "telegramInit")

	// Set logger for Telegram bot
	logTelegramInit.Debugln("Set logger for Telegram bot . . .")
	tgbotapi.SetLogger(log.WithField("prefix", "telegramBot"))

	// Connect Telegram API
	logTelegramInit.Debugln("Connect Telegram API . . .")
	if bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_BOT_KEY")); err == nil {
		telegramBot = bot
	} else {
		return err
	}

	// Upgrade debug mode
	if debug {
		logTelegramInit.Debugln("Upgrade Telegram bot . . .")
		telegramBot.Debug = true
	}

	return nil
}

func telegramStop() error {
	if telegramBot != nil {
		log.WithField("prefix", "telegramStop").Debugln("Stop Telegram bot . . .")
		telegramBot.StopReceivingUpdates()
		telegramBot = nil
	} else {
		return errors.New("Telegram bot has stopped")
	}
	return nil
}

func telegram() {
	// Set logger prefix
	logTelegram := log.WithField("prefix", "telegram")

	// Get updates from Telegram
	if updates, err := telegramBot.GetUpdatesChan(tgbotapi.UpdateConfig{Timeout: 60}); err != nil {
		logTelegram.WithError(err).Errorln("Cannot get updates from Telegram")
		exit(116)
	} else {
		for update := range updates {
			if update.Message == nil { // ignore any non-Message Updates
				continue
			}

			if update.Message.ReplyToMessage != nil {
				log.Traceln("Reply to message:", update.Message.ReplyToMessage)
				if update.Message.ReplyToMessage.From != nil {
					log.Traceln("Reply to author:", update.Message.ReplyToMessage.From)
				}
			}

			if update.Message.Text != "" {
				logTelegram.Traceln("Text")
				// tgBotText()
				logTelegram.Debugln(update.Message.Text)
			}
			if update.Message.Audio != nil {
				logTelegram.Traceln("Audio")
				// tgBotAudio()
				url, err := telegramBot.GetFileDirectURL(update.Message.Audio.FileID)
				logTelegram.Debugln(
					update.Message.Audio,
					url,
					err,
				)
			}
			if update.Message.Document != nil {
				logTelegram.Traceln("Document")
				// tgBotDocument()
				url, err := telegramBot.GetFileDirectURL(update.Message.Document.FileID)
				logTelegram.Debugln(
					update.Message.Document,
					url,
					err,
				)
			}
			if update.Message.Animation != nil {
				logTelegram.Traceln("Animation")
				// tgBotAnimation()
				url, err := telegramBot.GetFileDirectURL(update.Message.Animation.FileID)
				logTelegram.Debugln(
					update.Message.Animation,
					url,
					err,
				)
			}
			if update.Message.Game != nil {
				logTelegram.Traceln("Game")
				// tgBotGame()
				logTelegram.Debugln(
					update.Message.Game,
				)
			}
			if update.Message.Photo != nil {
				logTelegram.Traceln("Photo")
				// tgBotPhoto()
				logTelegram.Debugln(update.Message.Photo)
				for i, p := range *update.Message.Photo {
					url, err := telegramBot.GetFileDirectURL(p.FileID)
					logTelegram.Debugln(
						i,
						p,
						url,
						err,
					)
				}
			}
			if update.Message.Sticker != nil {
				logTelegram.Traceln("Sticker")
				// tgBotSticker()
				url, err := telegramBot.GetFileDirectURL(update.Message.Sticker.FileID)
				logTelegram.Debugln(
					update.Message.Sticker,
					url,
					err,
				)
			}
			if update.Message.Video != nil {
				logTelegram.Traceln("Video")
				// tgBotVideo()
				url, err := telegramBot.GetFileDirectURL(update.Message.Video.FileID)
				logTelegram.Debugln(
					update.Message.Video,
					url,
					err,
				)
			}
			if update.Message.Voice != nil {
				logTelegram.Traceln("Voice")
				// tgBotVoice()
				url, err := telegramBot.GetFileDirectURL(update.Message.Voice.FileID)
				logTelegram.Debugln(
					update.Message.Voice,
					url,
					err,
				)
			}
			if update.Message.VideoNote != nil {
				logTelegram.Traceln("VideoNote")
				// tgBotVideoNote()
				url, err := telegramBot.GetFileDirectURL(update.Message.VideoNote.FileID)
				logTelegram.Debugln(
					update.Message.VideoNote,
					url,
					err,
				)
			}
			if update.Message.Contact != nil {
				logTelegram.Traceln("Contact")
				// tgBotContact()
				logTelegram.Debugln(update.Message.Contact)
			}
			if update.Message.Location != nil {
				logTelegram.Traceln("Location")
				// tgBotLocation()
				logTelegram.Debugln(update.Message.Location)
			}
			if update.Message.Venue != nil {
				logTelegram.Traceln("Venue")
				// tgBotVenue()
				logTelegram.Debugln(update.Message.Venue)
			}
			if update.Message.Invoice != nil {
				logTelegram.Traceln("Invoice")
				// tgBotInvoice()
				logTelegram.Debugln(update.Message.Invoice)
			}
			if update.Message.SuccessfulPayment != nil {
				logTelegram.Traceln("SuccessfulPayment")
				// tgBotSuccessfulPayment()
				logTelegram.Debugln(update.Message.SuccessfulPayment)
			}
		}
	}
}
