package api

import (
	"context"
	"main/core"
	"net/http"
	"time"
)

func (a *API) postSessionHandler(w http.ResponseWriter, r *http.Request) {
	if !a.SessionWorker.Running() {
		var req core.Session
		err := a.readJson(r, &req)
		if err != nil {
			a.handleError(w, err, http.StatusBadRequest)
			return
		}
		session := core.Session{
			Date:     time.Now(),
			Interval: req.Interval,
			Limit:    req.Limit,
			Name:     req.Name,
			Frames:   0,
		}
		a.Log.Debug(session)
		ready := make(chan bool)
		go a.SessionWorker.Start(context.Background(), session, ready)
		<-ready
	}
	a.getSessionHandler(w, r)
}
