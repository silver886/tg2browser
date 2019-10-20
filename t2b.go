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
			logMain.Infoln("Start Telegram bot . . .")
			tg()
			logMain.Infoln("Telegram bot has stopped")
		}
		wg.Done()
	}()

	wg.Wait()
	exit(0)
}
