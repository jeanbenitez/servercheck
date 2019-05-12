package main

import (
	"net/http"
	"regexp"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// RenderDomainInfo return json output
func RenderDomainInfo(domain string) render.Renderer {
	return dbGetServer(domain)
}

func start(w http.ResponseWriter, r *http.Request) {
	domain := r.URL.Query().Get("domain")
	var message string
	if domain != "" {
		var isValidDomain bool
		isValidDomain, _ = regexp.MatchString("^([a-z0-9]+(-[a-z0-9]+)*\\.)+[a-z]{2,}$", domain)
		if isValidDomain {
			if err := render.Render(w, r, RenderDomainInfo(domain)); err != nil {
				render.Render(w, r, ErrRender(err))
				return
			}
		} else {
			message = "Invalid domain \"" + domain + "\""
		}
	} else {
		message = "Domain not found"
	}
	w.Write([]byte(message))
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", start)
	http.ListenAndServe(":3000", r)
}
