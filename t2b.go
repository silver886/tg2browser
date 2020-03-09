package main

import (
	"os"
	"t2b/tg"
)

const (
	appName    = "T2B"
	appVersion = "0.1.0"

	appTitle = appName + " " + appVersion
)

func main() {
	// Set logger prefix
	logMain := log.WithField("prefix", "main")

	// Parse flags
	if err := flagInit(); err != nil {
		logMain.WithError(err).Errorln("Cannot parse flags")
		exit(11)
	}

	// Telegram bot
	for {
		// Initialize Telegram bot
		logMain.Debugln("Initialize Telegram bot . . .")
		if err := tg.Init(log, os.Getenv("TG_BOT_KEY"), debug); err != nil {
			logMain.WithError(err).Errorln("Cannot initialize Telegram bot")
			continue
		}

		// Start Telegram bot
		logMain.Infoln("Start Telegram bot . . .")
		tg.Main()

		// Stop Telegram bot
		logMain.Infoln("Stop Telegram bot . . .")
		tg.Stop()
	}
}
