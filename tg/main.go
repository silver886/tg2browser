package tg

import (
	"t2b/tg/msg"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Main kicks off the commnuication to telegram.
func Main() error {
	// Set logger prefix
	logMain := log.WithField("prefix", "main")

	// Get updates from Telegram
	if updates, err := telegramBot.GetUpdatesChan(tgbotapi.UpdateConfig{Timeout: 60}); err != nil {
		logMain.WithError(err).Errorln("Cannot get updates from Telegram")
		return err
	} else {
		for update := range updates {
			// Ignore non-message Updates
			if update.Message == nil {
				continue
			}

			// if m.ReplyToMessage != nil {
			// 	log.Traceln("Reply to message:", m.ReplyToMessage)
			// 	if m.ReplyToMessage.From != nil {
			// 		log.Traceln("Reply to author:", m.ReplyToMessage.From)
			// 	}
			// }
			if update.Message.Entities != nil {
				msgs := msg.GetEntities(update.Message)
				for _, m := range msgs {
					logMain.Infof(m.Text)
					send(m)
				}
			}
			if update.Message.Text != "" {
				m := msg.GetText(update.Message)
				logMain.Infof(m.Text)
				send(m)
			}
			// if update.Message.Audio != nil {
			// 	logMain.Traceln("Audio")
			// 	// tgBotAudio()
			// 	url, err := telegramBot.GetFileDirectURL(update.Message.Audio.FileID)
			// 	logMain.Debugln(
			// 		update.Message.Audio,
			// 		url,
			// 		err,
			// 	)
			// }
			// if update.Message.Document != nil {
			// 	logMain.Traceln("Document")
			// 	// tgBotDocument()
			// 	url, err := telegramBot.GetFileDirectURL(update.Message.Document.FileID)
			// 	logMain.Debugln(
			// 		update.Message.Document,
			// 		url,
			// 		err,
			// 	)
			// }
			// if update.Message.Animation != nil {
			// 	logMain.Traceln("Animation")
			// 	// tgBotAnimation()
			// 	url, err := telegramBot.GetFileDirectURL(update.Message.Animation.FileID)
			// 	logMain.Debugln(
			// 		update.Message.Animation,
			// 		url,
			// 		err,
			// 	)
			// }
			// if update.Message.Game != nil {
			// 	logMain.Traceln("Game")
			// 	// tgBotGame()
			// 	logMain.Debugln(
			// 		update.Message.Game,
			// 	)
			// }
			// if update.Message.Photo != nil {
			// 	logMain.Traceln("Photo")
			// 	// tgBotPhoto()
			// 	logMain.Debugln(update.Message.Photo)
			// 	for i, p := range *update.Message.Photo {
			// 		url, err := telegramBot.GetFileDirectURL(p.FileID)
			// 		logMain.Debugln(
			// 			i,
			// 			p,
			// 			url,
			// 			err,
			// 		)
			// 	}
			// }
			// if update.Message.Sticker != nil {
			// 	logMain.Traceln("Sticker")
			// 	// tgBotSticker()
			// 	url, err := telegramBot.GetFileDirectURL(update.Message.Sticker.FileID)
			// 	logMain.Debugln(
			// 		update.Message.Sticker,
			// 		url,
			// 		err,
			// 	)
			// }
			// if update.Message.Video != nil {
			// 	logMain.Traceln("Video")
			// 	// tgBotVideo()
			// 	url, err := telegramBot.GetFileDirectURL(update.Message.Video.FileID)
			// 	logMain.Debugln(
			// 		update.Message.Video,
			// 		url,
			// 		err,
			// 	)
			// }
			// if update.Message.Voice != nil {
			// 	logMain.Traceln("Voice")
			// 	// tgBotVoice()
			// 	url, err := telegramBot.GetFileDirectURL(update.Message.Voice.FileID)
			// 	logMain.Debugln(
			// 		update.Message.Voice,
			// 		url,
			// 		err,
			// 	)
			// }
			// if update.Message.VideoNote != nil {
			// 	logMain.Traceln("VideoNote")
			// 	// tgBotVideoNote()
			// 	url, err := telegramBot.GetFileDirectURL(update.Message.VideoNote.FileID)
			// 	logMain.Debugln(
			// 		update.Message.VideoNote,
			// 		url,
			// 		err,
			// 	)
			// }
			// if update.Message.Contact != nil {
			// 	logMain.Traceln("Contact")
			// 	// tgBotContact()
			// 	logMain.Debugln(update.Message.Contact)
			// }
			// if update.Message.Location != nil {
			// 	logMain.Traceln("Location")
			// 	// tgBotLocation()
			// 	logMain.Debugln(update.Message.Location)
			// }
			// if update.Message.Venue != nil {
			// 	logMain.Traceln("Venue")
			// 	// tgBotVenue()
			// 	logMain.Debugln(update.Message.Venue)
			// }
			// if update.Message.Invoice != nil {
			// 	logMain.Traceln("Invoice")
			// 	// tgBotInvoice()
			// 	logMain.Debugln(update.Message.Invoice)
			// }
			// if update.Message.SuccessfulPayment != nil {
			// 	logMain.Traceln("SuccessfulPayment")
			// 	// tgBotSuccessfulPayment()
			// 	logMain.Debugln(update.Message.SuccessfulPayment)
			// }
		}
	}
	return nil
}
