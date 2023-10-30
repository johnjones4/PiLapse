package api

import "net/http"

func (a *API) imageHandler(w http.ResponseWriter, r *http.Request) {
	img := a.SessionWorker.CurrentImage()
	if img == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Content-type", "image/jpeg")
	w.Write(img)
}
