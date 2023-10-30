package main

import (
	"context"
	"encoding/json"
	"main/api"
	"main/core"
	"main/impl/ftptransmitter"
	"main/impl/picamera"
	"main/impl/sessionworker"
	"main/impl/telegramtransmitter"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

type configuration struct {
	FTP      ftptransmitter.FTPTransmitterConfiguration           `json:"ftp"`
	Telegram telegramtransmitter.TelegramTransmitterConfiguraiton `json:"telegram"`
}

func main() {
	log := logrus.New()

	loglevel, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		loglevel = logrus.DebugLevel
	}
	log.SetLevel(loglevel)

	bytes, err := os.ReadFile(os.Getenv("CONFIG_FILE"))
	if err != nil {
		log.Panic(err)
	}

	var config configuration
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		log.Panic(err)
	}

	transmiters := make([]core.TransmissionProvider, 0)

	if !config.FTP.IsZero() {
		transmiters = append(transmiters, &ftptransmitter.FTPTransmitter{
			FTPTransmitterConfiguration: config.FTP,
			Log:                         log,
		})
	}
	if !config.Telegram.IsZero() {
		transmiters = append(transmiters, &telegramtransmitter.TelegramTransmitter{
			TelegramTransmitterConfiguraiton: config.Telegram,
			Log:                              log,
		})
	}

	worker := sessionworker.NewSessionWorker(sessionworker.SessionWorkerParams{
		Log: log,
		Img: &picamera.PiCamera{
			Log: log,
		},
		Trans: transmiters,
	})
	go worker.StartBackground(context.Background())

	api := api.New(api.APIParams{
		Log:           log,
		SessionWorker: worker,
	})

	http.ListenAndServe(os.Getenv("HTTP_HOST"), api)
}
