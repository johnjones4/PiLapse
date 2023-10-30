package api

import (
	"main/core"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

type APIParams struct {
	Log           logrus.FieldLogger
	SessionWorker core.SessionWorker
}

type API struct {
	mux *chi.Mux
	APIParams
}

func (a *API) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	a.mux.ServeHTTP(res, req)
}

func New(params APIParams) *API {
	a := &API{
		mux:       chi.NewRouter(),
		APIParams: params,
	}

	a.mux.Use(middleware.RequestID)
	a.mux.Use(middleware.RealIP)
	a.mux.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: params.Log}))
	a.mux.Use(middleware.Recoverer)

	a.mux.Route("/api", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
		})

		r.Get("/image", a.imageHandler)
		r.Route("/session", func(r chi.Router) {
			r.Get("/", a.getSessionHandler)
			r.Post("/", a.postSessionHandler)
			r.Delete("/", a.deleteSessionHandler)
		})
	})

	return a
}
