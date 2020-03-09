package tg

import "errors"

// Stop terminates related telegram services.
func Stop() error {
	if telegramBot != nil {
		log.WithField("prefix", "stop").Debugln("Stop Telegram bot . . .")
		telegramBot.StopReceivingUpdates()
		telegramBot = nil
	} else {
		return errors.New("Telegram bot had stopped")
	}
	return nil
}
