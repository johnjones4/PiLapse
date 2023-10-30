package api

import (
	"main/core"
	"net/http"
)

type sessionResponse struct {
	Session core.Session `json:"session,omitempty"`
	Running bool         `json:"running"`
}

func (a *API) getSessionHandler(w http.ResponseWriter, r *http.Request) {
	req := sessionResponse{
		Session: a.SessionWorker.Session(),
		Running: a.SessionWorker.Running(),
	}
	a.jsonResponse(w, req)
}
