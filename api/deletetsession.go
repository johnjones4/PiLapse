package api

import (
	"net/http"
)

func (a *API) deleteSessionHandler(w http.ResponseWriter, r *http.Request) {
	a.SessionWorker.Stop()
	w.WriteHeader(http.StatusOK)
}
