package picamera

import (
	"os"
	"os/exec"
	"sync"

	"github.com/sirupsen/logrus"
)

const (
	file = "/tmp/image.jpg"
)

type PiCamera struct {
	lock sync.Mutex
	Log  logrus.FieldLogger
}

func (c *PiCamera) CaptureImage() ([]byte, error) {
	c.Log.Debug("Capturing image")
	defer c.Log.Debug("Done capturing image")

	c.lock.Lock()
	defer c.lock.Unlock()
	cm := exec.Command("libcamera-still", "-o", file)
	err := cm.Run()
	if err != nil {
		return nil, err
	}

	bytes, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	os.Remove(file)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
