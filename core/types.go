package core

import (
	"context"
	"time"
)

type Session struct {
	Date     time.Time `json:"date"`
	Limit    Duration  `json:"limit"`
	Name     string    `json:"name"`
	Frames   int       `json:"frames"`
	Interval Duration  `json:"interval"`
}

type SessionWorker interface {
	Running() bool
	Session() Session
	CurrentImage() []byte
	Start(ctx context.Context, session Session, ready chan bool)
	Stop()
}

type ImagingProvider interface {
	CaptureImage() ([]byte, error)
}

type TransmissionProvider interface {
	Transmit(session Session, n int, img []byte) error
}
