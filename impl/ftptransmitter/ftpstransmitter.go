package ftptransmitter

import (
	"bytes"
	"fmt"
	"main/core"
	"path"
	"time"

	"github.com/flytam/filenamify"
	"github.com/secsy/goftp"
	"github.com/sirupsen/logrus"
)

type FTPTransmitterConfiguration struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Path     string `json:"path"`
}

func (c *FTPTransmitterConfiguration) IsZero() bool {
	return c.Host == ""
}

type FTPTransmitter struct {
	FTPTransmitterConfiguration
	Log logrus.FieldLogger
}

func (t *FTPTransmitter) Transmit(session core.Session, n int, img []byte) error {
	t.Log.Debug("Transmitting via ftp")
	defer t.Log.Debug("Done transmitting via ftp")

	client, err := goftp.DialConfig(goftp.Config{
		User:     t.Username,
		Password: t.Password,
	}, t.Host)
	if err != nil {
		return err
	}

	name, err := filenamify.Filenamify(session.Name, filenamify.Options{})
	if err != nil {
		return err
	}

	projectFolder := path.Join(t.Path, fmt.Sprintf("%s_%s", session.Date.Format(time.DateOnly), name))
	_, err = client.Stat(projectFolder)
	if err != nil {
		if ftperr, ok := err.(goftp.Error); ok && ftperr.Code() == 550 {
			_, err = client.Mkdir(projectFolder)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	err = client.Store(path.Join(projectFolder, fmt.Sprintf("%010d.jpg", n)), bytes.NewReader(img))
	if err != nil {
		return err
	}

	err = client.Close()
	if err != nil {
		return err
	}

	return nil
}
