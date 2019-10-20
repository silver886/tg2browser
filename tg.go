package main

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func tgInit() error {
	// Set logger prefix
	logTgInit := log.WithField("prefix", "tgInit")

	// Set logger for Telegram bot
	logTgInit.Debugln("Set logger for Telegram bot")
	tgbotapi.SetLogger(log)

	// Start Telegram bot
	logTgInit.Debugln("Start Telegram bot")
	if bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_BOT_KEY")); err == nil {
		tgBot = bot
	} else {
		return err
	}
	return nil
}

func tg() {
	// Set logger prefix
	logTgBot := log.WithField("prefix", "tgBot")

	// Start get updates from Telegram
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	if updates, err := tgBot.GetUpdatesChan(u); err != nil {
		logTgBot.WithError(err).Errorln("Cannot get updates from Telegram")
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
				logTgBot.Traceln("Text")
				// tgBotText()
				logTgBot.Debugln(update.Message.Text)
			}
			if update.Message.Audio != nil {
				logTgBot.Traceln("Audio")
				// tgBotAudio()
				url, err := tgBot.GetFileDirectURL(update.Message.Audio.FileID)
				logTgBot.Debugln(
					update.Message.Audio,
					url,
					err,
				)
			}
			if update.Message.Document != nil {
				logTgBot.Traceln("Document")
				// tgBotDocument()
				url, err := tgBot.GetFileDirectURL(update.Message.Document.FileID)
				logTgBot.Debugln(
					update.Message.Document,
					url,
					err,
				)
			}
			if update.Message.Animation != nil {
				logTgBot.Traceln("Animation")
				// tgBotAnimation()
				url, err := tgBot.GetFileDirectURL(update.Message.Animation.FileID)
				logTgBot.Debugln(
					update.Message.Animation,
					url,
					err,
				)
			}
			if update.Message.Game != nil {
				logTgBot.Traceln("Game")
				// tgBotGame()
				logTgBot.Debugln(
					update.Message.Game,
				)
			}
			if update.Message.Photo != nil {
				logTgBot.Traceln("Photo")
				// tgBotPhoto()
				logTgBot.Debugln(update.Message.Photo)
				for i, p := range *update.Message.Photo {
					url, err := tgBot.GetFileDirectURL(p.FileID)
					logTgBot.Debugln(
						i,
						p,
						url,
						err,
					)
				}
			}
			if update.Message.Sticker != nil {
				logTgBot.Traceln("Sticker")
				// tgBotSticker()
				url, err := tgBot.GetFileDirectURL(update.Message.Sticker.FileID)
				logTgBot.Debugln(
					update.Message.Sticker,
					url,
					err,
				)
			}
			if update.Message.Video != nil {
				logTgBot.Traceln("Video")
				// tgBotVideo()
				url, err := tgBot.GetFileDirectURL(update.Message.Video.FileID)
				logTgBot.Debugln(
					update.Message.Video,
					url,
					err,
				)
			}
			if update.Message.Voice != nil {
				logTgBot.Traceln("Voice")
				// tgBotVoice()
				url, err := tgBot.GetFileDirectURL(update.Message.Voice.FileID)
				logTgBot.Debugln(
					update.Message.Voice,
					url,
					err,
				)
			}
			if update.Message.VideoNote != nil {
				logTgBot.Traceln("VideoNote")
				// tgBotVideoNote()
				url, err := tgBot.GetFileDirectURL(update.Message.VideoNote.FileID)
				logTgBot.Debugln(
					update.Message.VideoNote,
					url,
					err,
				)
			}
			if update.Message.Contact != nil {
				logTgBot.Traceln("Contact")
				// tgBotContact()
				logTgBot.Debugln(update.Message.Contact)
			}
			if update.Message.Location != nil {
				logTgBot.Traceln("Location")
				// tgBotLocation()
				logTgBot.Debugln(update.Message.Location)
			}
			if update.Message.Venue != nil {
				logTgBot.Traceln("Venue")
				// tgBotVenue()
				logTgBot.Debugln(update.Message.Venue)
			}
			if update.Message.Invoice != nil {
				logTgBot.Traceln("Invoice")
				// tgBotInvoice()
				logTgBot.Debugln(update.Message.Invoice)
			}
			if update.Message.SuccessfulPayment != nil {
				logTgBot.Traceln("SuccessfulPayment")
				// tgBotSuccessfulPayment()
				logTgBot.Debugln(update.Message.SuccessfulPayment)
			}
		}
	}
}
