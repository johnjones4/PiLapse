package sessionworker

import (
	"context"
	"main/core"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

type SessionWorkerParams struct {
	Log   logrus.FieldLogger
	Img   core.ImagingProvider
	Trans []core.TransmissionProvider
}

type frame struct {
	session core.Session
	img     []byte
}

type SessionWorker struct {
	SessionWorkerParams
	session      core.Session
	cancel       context.CancelFunc
	currentImage []byte
	lock         sync.Mutex
	frames       chan frame
}

func NewSessionWorker(params SessionWorkerParams) *SessionWorker {
	return &SessionWorker{
		SessionWorkerParams: params,
		frames:              make(chan frame, 1024),
	}
}

func (w *SessionWorker) StartBackground(ctx context.Context) {
	w.Log.Info("Starting session background worker")
	for {
		select {
		case <-ctx.Done():
			return
		case f := <-w.frames:
			w.Log.Infof("Handling update frame %d", w.session.Frames)
			w.lock.Lock()
			w.session.Frames++
			w.currentImage = f.img
			w.lock.Unlock()
			for _, t := range w.Trans {
				err := t.Transmit(f.session, f.session.Frames, f.img)
				if err != nil {
					w.Log.Error(err)
				}
			}
		}
	}
}

func (w *SessionWorker) Running() bool {
	w.lock.Lock()
	defer w.lock.Unlock()
	return w.cancel != nil
}

func (w *SessionWorker) Session() core.Session {
	w.lock.Lock()
	defer w.lock.Unlock()
	return w.session
}

func (w *SessionWorker) CurrentImage() []byte {
	w.lock.Lock()
	defer w.lock.Unlock()
	if w.currentImage == nil {
		return nil
	}
	image := make([]byte, len(w.currentImage))
	copy(image, w.currentImage)
	return image
}

func (w *SessionWorker) Stop() {
	w.lock.Lock()
	defer w.lock.Unlock()
	if w.cancel != nil {
		w.cancel()
		w.cancel = nil
	}
}

func (w *SessionWorker) Start(ctx context.Context, session core.Session, ready chan bool) {
	w.Log.Info("Starting capture session")
	ticker := time.NewTicker(session.Interval.Duration)
	limit := time.NewTicker(session.Limit.Duration)
	ctx1, cancel := context.WithCancel(ctx)
	w.cancel = cancel
	w.session = session
	ready <- true
	var last time.Time
	for {
		select {
		case <-ctx1.Done():
			w.Log.Info("Capturing killed")
			return
		case <-limit.C:
			w.Log.Info("Capturing completed on time")
			return
		case <-ticker.C:
			w.Log.Info("Starting capture")
			w.Log.Debugf("Time since last capture %s (Interval: %s)", time.Since(last).String(), w.session.Interval.String())
			w.Log.Debugf("Time since start %s (Running for %s)", time.Since(w.session.Date).String(), w.session.Limit.String())
			last = time.Now()
			img, err := w.Img.CaptureImage()
			if err != nil {
				w.Log.Error(err)
				continue
			}
			w.frames <- frame{
				session: w.session,
				img:     img,
			}
			w.Log.Info("Completed capture")
		}
	}
}
