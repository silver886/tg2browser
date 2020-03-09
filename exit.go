package main

import (
	"os"
	"os/signal"
	"syscall"
	"t2b/tg"

	"github.com/sirupsen/logrus"
)

func exitInit() error {
	c := make(chan os.Signal, 1)
	signal.Notify(c,
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	go func() {
		for sig := range c {
			log.WithFields(logrus.Fields{
				"prefix": "signal",
				"signal": sig,
			}).Debugln("Get OS signal")
			exit(7)
		}
	}()
	return nil
}

func exit(code int) {
	// Set logger prefix
	logExit := log.WithField("prefix", "exit")

	logExit.WithField("exit_code", code).Infoln("Exiting program . . .")

	// Stop Telegram bot
	tg.Stop()

	// Exit program with exit code
	logExit.WithField("exit_code", code).Debugln("Exit . . .")
	// Wait logger completed
	log.Wait()
	os.Exit(code)
}
