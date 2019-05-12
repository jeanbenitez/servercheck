package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

//--
// Data model objects and persistence mocks:
//--

// ServerItem data model
type ServerItem struct {
	Address  string `json:"address"`
	SslGrade string `json:"ssl_grade"`
	Country  string `json:"country"`
	Owner    string `json:"owner"`
}

// ServerResponse data model
type ServerResponse struct {
	Domain           string `json:"domain"`
	ServersChanged   string `json:"servers_changed"`
	SslGrade         string `json:"ssl_grade"`
	PreviousSslGrade string `json:"previous_ssl_grade"`
	Logo             string `json:"logo"`
	IsDown           string `json:"is_down"`
	Servers          []ServerItem
}

// Render for Server struct
func (rd *ServerResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	// rd.Elapsed = 10
	return nil
}

// Business logic

var servers []*ServerResponse

func dbNewServer(server *ServerResponse) (string, error) {
	servers = append(servers, server)
	return server.Domain, nil
}

func dbGetServer(domain string) *ServerResponse {
	db := connect()
	rows, err := db.Query("SELECT domain, servers_changed, ssl_grade, previous_ssl_grade, logo, is_down FROM servers WHERE domain = $1", domain)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	server := ServerResponse{}
	for rows.Next() {
		err = rows.Scan(
			&server.Domain,
			&server.ServersChanged,
			&server.SslGrade,
			&server.PreviousSslGrade,
			&server.Logo,
			&server.IsDown,
		)
		if err != nil {
			log.Fatal(err)
		}
	}
	return &server
}

func dbUpdateServer(domain string, server *ServerResponse) (*ServerResponse, error) {
	for i, s := range servers {
		if s.Domain == domain {
			servers[i] = server
			return server, nil
		}
	}
	return nil, errors.New("server not found")
}

// Database connection logic

func connect() (db *sql.DB) {
	// Connect to the "bank" database.
	db, err := sql.Open("postgres", "postgresql://root@localhost:26257/servercheck?sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	return db
}
