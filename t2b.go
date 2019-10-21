package main

import "sync"

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

	// Create wait group for services
	wg := &sync.WaitGroup{}

	// Telegram bot
	wg.Add(1)
	go func() {
		for {
			// Initialize Telegram bot
			logMain.Debugln("Initialize Telegram bot . . .")
			if err := telegramInit(); err != nil {
				logMain.WithError(err).Errorln("Cannot initialize Telegram bot")
				continue
			}

			// Start Telegram bot
			logMain.Infoln("Start Telegram bot . . .")
			telegram()

			// Stop Telegram bot
			logMain.Infoln("Stop Telegram bot . . .")
			telegramStop()
		}
	}()

	wg.Wait()
	exit(0)
}
