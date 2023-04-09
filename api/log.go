package api

import (
	"log"
	"os"
)

func SetupLogger(config ApiConfig) {
	if config.LogFile == "" {
		config.LogFile = "seekr.log"
	}
	f, err := os.OpenFile(config.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		// try creating the file first
		f, err = os.Create(config.LogFile)
		if err != nil {
			log.Fatalf("error opening file: %s\nLog file: %s", err, config.LogFile)
		}
	}

	log.SetOutput(f)
	log.Printf("opening log file: %s", config.LogFile)
}

func CheckAndLog(err error, msg string, config ApiConfig) { // unused but may be needed in the future
	if err != nil {
		log.Printf("error: %v\n%s", err, msg)
	}
}
