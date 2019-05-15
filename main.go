package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jeanbenitez/servercheck/handlers"
)

// RenderDomainInfo return json output
// func RenderDomainInfo(domain string) render.Renderer {
// 	domainData := getDomainDataInSllLabs(domain)
// 	whoisData := getWhoisDomain(domain)
// 	server := dbGetServer(domain)

// 	// SSL Grades
// 	grades := []string{"A", "A+", "B", "B+", "C", "C+", "D", "D+", "E", "E+", "F", "F+"}

// 	savedGrade := utils.IndexOf(server.SslGrade, grades)
// 	newBestGrade := 0
// 	for _, v := range domainData.Endpoints {
// 		if utils.IndexOf(v.Grade, grades) > savedGrade {
// 			savedGrade = utils.IndexOf(v.Grade, grades)
// 		}
// 	}

// 	log.Output(1, whoisData.Changed+string(newBestGrade))

// 	return server
// }

// func start(w http.ResponseWriter, r *http.Request) {
// 	domain := r.URL.Query().Get("domain")
// 	var message string
// 	if domain != "" {
// 		var isValidDomain bool
// 		isValidDomain, _ = regexp.MatchString("^([a-z0-9]+(-[a-z0-9]+)*\\.)+[a-z]{2,}$", domain)
// 		if isValidDomain {
// 			render.Render(w, r, RenderDomainInfo(domain))
// 		} else {
// 			message = "Invalid domain \"" + domain + "\""
// 		}
// 	} else {
// 		message = "Domain not found"
// 	}
// 	w.Write([]byte(message))
// }

func main() {
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	goPort := os.Getenv("GOPORT")

	// DB Connection
	dataSource := "postgresql://" + dbUser + ":" + dbPass + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=disable"

	conn, err := sql.Open("postgres", dataSource)
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	r := chi.NewRouter()
	// r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// r.Use(middleware.URLFormat)
	// r.Use(render.SetContentType(render.ContentTypeJSON))

	dHandler := handlers.NewDomainHandler(conn)
	r.Route("/", func(rt chi.Router) {
		rt.Mount("/domains", domainRouter(dHandler))
	})

	fmt.Println("Server listen at localhost:" + goPort)
	http.ListenAndServe(":"+goPort, r)
}

// A completely separate router for posts routes
func domainRouter(dHandler *handlers.Domain) http.Handler {
	r := chi.NewRouter()
	r.Get("/", dHandler.Fetch)
	r.Get("/{domain:^([a-z0-9]+(-[a-z0-9]+)*\\.)+[a-z]{2,}$}", dHandler.GetByDomain)
	// r.Post("/", dHandler.Create)
	// r.Put("/{id:[0-9]+}", dHandler.Update)
	// r.Delete("/{id:[0-9]+}", dHandler.Delete)

	return r
}
