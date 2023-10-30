package telegramtransmitter

import (
	"bytes"
	"fmt"
	"main/core"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type TelegramTransmitterConfiguraiton struct {
	SendInterval core.Duration `json:"sendInterval"`
	ChatId       int           `json:"chatId"`
	BotId        string        `json:"botId"`
}

func (c *TelegramTransmitterConfiguraiton) IsZero() bool {
	return c.SendInterval.Duration == 0 && c.ChatId == 0 && c.BotId == ""
}

type TelegramTransmitter struct {
	TelegramTransmitterConfiguraiton
	Log      logrus.FieldLogger
	lastSend time.Time
}

func (t *TelegramTransmitter) Transmit(session core.Session, n int, img []byte) error {
	t.Log.Debug("Transmitting via telegram")
	defer t.Log.Debug("Done transmitting via telegram")

	t.Log.Debug("Time since last Telegram: ", time.Since(t.lastSend), t.SendInterval.Duration)

	if time.Since(t.lastSend) < t.SendInterval.Duration {
		return nil
	}

	body := &bytes.Buffer{}

	mwriter := multipart.NewWriter(body)

	w, err := mwriter.CreateFormFile("photo", "photo.jpg")
	if err != nil {
		return err
	}

	_, err = w.Write(img)
	if err != nil {
		return err
	}

	err = mwriter.Close()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto?chat_id=%d", t.BotId, t.ChatId)
	req, err := http.NewRequest(http.MethodPost, uri, bytes.NewReader(body.Bytes()))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", mwriter.FormDataContentType())
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("bad response status from telegram: %d", res.StatusCode)
	}

	t.lastSend = time.Now()

	return nil
}
