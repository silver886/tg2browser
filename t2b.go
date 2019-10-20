package main

import (
	"sync"
)

const (
	appName    = "T2B"
	appVersion = "0.1.0"

	appTitle = appName + " " + appVersion
)

func main() {
	// Set logger prefix
	logMain := log.WithField("prefix", "main")

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		for {
			logMain.Infoln("Start Telegram bot . . .")
			tgBot()
			logMain.Infoln("Telegram bot has stopped")
		}
		wg.Done()
	}()

	wg.Wait()
	exit(0)
}
