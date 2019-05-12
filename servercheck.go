package main

import (
	"log"
	"net/http"
	"regexp"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

// RenderDomainInfo return json output
func RenderDomainInfo(domain string) render.Renderer {
	domainData := getDomainDataInSllLabs(domain)
	whoisData := getWhoisDomain(domain)
	server := dbGetServer(domain)

	// SSL Grades
	grades := []string{"A", "A+", "B", "B+", "C", "C+", "D", "D+", "E", "E+", "F", "F+"}

	savedGrade := indexOf(server.SslGrade, grades)
	newBestGrade := 0
	for _, v := range domainData.Endpoints {
		if indexOf(v.Grade, grades) > savedGrade {
			savedGrade = indexOf(v.Grade, grades)
		}
	}

	log.Output(1, whoisData.Changed+string(newBestGrade))

	return server
}

func start(w http.ResponseWriter, r *http.Request) {
	domain := r.URL.Query().Get("domain")
	var message string
	if domain != "" {
		var isValidDomain bool
		isValidDomain, _ = regexp.MatchString("^([a-z0-9]+(-[a-z0-9]+)*\\.)+[a-z]{2,}$", domain)
		if isValidDomain {
			render.Render(w, r, RenderDomainInfo(domain))
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
