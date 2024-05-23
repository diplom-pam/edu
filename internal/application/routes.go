package application

import (
	"net/http"

	"github.com/diplom-pam/edu/internal/api/handlers/createquestion"
	"github.com/diplom-pam/edu/internal/api/handlers/getquestion"
	"github.com/go-chi/chi/v5"
)

func (app *Application) registerRoutes(r chi.Router) {
	r.Route("/edu", func(r chi.Router) {
		r.Use(app.middlewareCORS)

		r.Get("/question/{id}", func(rw http.ResponseWriter, req *http.Request) { getquestion.Handler(rw, req, app.store) })
		r.Post("/question", func(rw http.ResponseWriter, req *http.Request) { createquestion.Handler(rw, req, app.store) })

		//r.Route("", func(r chi.Router) {
		//	r.Use(app.middlewareAuth)
		//
		//	r.Route("/reports", func(r chi.Router) {
		//		r.Post("/dsp/clicklog", func(rw http.ResponseWriter, req *http.Request) { dspclicklog.Handler(rw, req, app.store) })
		//	})
		//})
	})
}
