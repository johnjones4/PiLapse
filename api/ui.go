package api

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"main/core"
	"net/http"
	"text/template"
	"time"
)

//go:embed templates/index.html
var indexTemplate string

var tmpl *template.Template

type uiSessionInfoRow struct {
	Key   string
	Value string
}

type uiParams struct {
	SessionInfo []uiSessionInfoRow
	Running     bool
}

func (a *API) uiHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	if r.Method == "POST" {
		err = r.ParseForm()
		if err != nil {
			a.handleError(w, err, http.StatusBadRequest)
			return
		}

		if r.Form.Get("stop") == "true" {
			a.SessionWorker.Stop()
		} else {
			interval, err := time.ParseDuration(r.Form.Get("interval"))
			if err != nil {
				a.handleError(w, err, http.StatusBadRequest)
				return
			}

			limit, err := time.ParseDuration(r.Form.Get("limit"))
			if err != nil {
				a.handleError(w, err, http.StatusBadRequest)
				return
			}

			name := r.Form.Get("name")
			if name == "" {
				a.handleError(w, errors.New("name required"), http.StatusBadRequest)
				return
			}

			session := core.Session{
				Date:     time.Now(),
				Interval: core.Duration{Duration: interval},
				Limit:    core.Duration{Duration: limit},
				Name:     name,
				Frames:   0,
			}
			ready := make(chan bool)
			go a.SessionWorker.Start(context.Background(), session, ready)
			<-ready
		}
	}

	if tmpl == nil {
		tmpl, err = template.New("index").Parse(indexTemplate)
		if err != nil {
			a.handleError(w, err, http.StatusInternalServerError)
			return
		}
	}

	session := a.SessionWorker.Session()
	params := uiParams{
		SessionInfo: []uiSessionInfoRow{
			{"Name", session.Name},
			{"Started", session.Date.Format(time.DateTime)},
			{"Time Elapsed", time.Since(session.Date).String()},
			{"Time Remaining", time.Until(session.Date.Add(session.Limit.Duration)).String()},
			{"Frames Captured", fmt.Sprint(session.Frames)},
			{"Frame Interval", session.Interval.String()},
		},
		Running: a.SessionWorker.Running(),
	}

	tmpl.Execute(w, params)
}
