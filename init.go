package main

import (
	"flag"

	"github.com/silver886/logger"
	"github.com/sirupsen/logrus"
)

var (
	log *logger.Logger

	debug bool
)

func init() {
	// Define arguments
	flag.BoolVar(&debug, "debug", false, "Enable debug mode")

	// Create logger
	log, _ = logger.New(appName,
		logrus.Level(len(logrus.AllLevels)-1),
		[]logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
			logrus.WarnLevel,
			logrus.InfoLevel,
		},
		false,
	)

	logInit := log.WithField("prefix", "init")

	logInit.Debugln("Create logger")

	// Setup OS signal handler
	logInit.Debugln("Setup OS signal handler . . .")
	if err := exitInit(); err != nil {
		logInit.WithError(err).Errorln("Cannot setup OS signal handler")
		exit(9)
	}
}

func flagInit() error {
	// Set logger prefix
	logFlagInit := log.WithField("prefix", "flagInit")

	// Parse arguments
	flag.Parse()

	// Upgrade debug mode
	if debug {
		logFlagInit.Debugln("Upgrade logger . . .")
		log.Wait()
		log.Config(
			logrus.Level(len(logrus.AllLevels)-1),
			logrus.AllLevels,
			false,
		)
	}

	logFlagInit.WithFields(logrus.Fields{
		"debug": debug,
	}).Debugln("Parse arguments")

	return nil
}
