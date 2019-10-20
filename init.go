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

	log.WithField("prefix", "init").Debugln("Create logger")
}

func flagInit() error {
	// Set logger prefix
	logFlagInit := log.WithField("prefix", "flagInit")

	// Parse arguments
	flag.Parse()

	if debug {
		log.Wait()
		log.Config(
			logrus.TraceLevel,
			logrus.AllLevels,
			false,
		)
		logFlagInit.Debugln("Update logger")
	}

	logFlagInit.WithFields(logrus.Fields{
		"debug": debug,
	}).Debugln("Parse arguments")

	return nil
}
