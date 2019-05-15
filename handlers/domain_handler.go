package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"

	"github.com/jeanbenitez/servercheck/controllers"
	"github.com/jeanbenitez/servercheck/interfaces"
	"github.com/jeanbenitez/servercheck/models"
	"github.com/jeanbenitez/servercheck/services"
	"github.com/jeanbenitez/servercheck/utils"
)

// NewDomainHandler ...
func NewDomainHandler(db *sql.DB) *Domain {
	return &Domain{
		domainController: controllers.NewControllerDomain(db),
		serverController: controllers.NewControllerServer(db),
	}
}

// Domain ...
type Domain struct {
	domainController interfaces.IDomainController
	serverController interfaces.IServerController
}

// Fetch all domain data
func (d *Domain) Fetch(w http.ResponseWriter, r *http.Request) {
	payload, _ := d.domainController.Fetch(r.Context(), 5)

	utils.RespondwithJSON(w, http.StatusOK, payload)
}

// Create a new domain
func (d *Domain) Create(w http.ResponseWriter, r *http.Request) {
	domain := models.Domain{}
	json.NewDecoder(r.Body).Decode(&domain)

	created, err := d.domainController.Create(r.Context(), &domain)
	fmt.Println(created)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Server Error")
	} else {
		utils.RespondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
	}
}

// Update a domain by domain
func (d *Domain) Update(w http.ResponseWriter, r *http.Request) {
	domain := chi.URLParam(r, "domain")
	data := models.Domain{Domain: string(domain)}
	json.NewDecoder(r.Body).Decode(&data)
	payload, err := d.domainController.Update(r.Context(), &data)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Server Error")
	} else {
		utils.RespondwithJSON(w, http.StatusOK, payload)
	}
}

// GetByDomain returns a domain details
func (d *Domain) GetByDomain(w http.ResponseWriter, r *http.Request) {
	domain := chi.URLParam(r, "domain")
	domainPayload, err := d.domainController.GetByDomain(r.Context(), domain)

	if err != nil {
		domainPayload = new(models.Domain)
	}

	serversPayload, _ := d.serverController.FetchByDomain(r.Context(), domain)

	fmt.Println("Servers: " + string(len(serversPayload)))

	domainData := services.GetSslLabsDomainData(domain)

	// SSL Grades
	grades := []string{"F", "E", "D", "C", "B", "A", "A+"}
	savedGrade := 0

	if err == nil {
		savedGrade = utils.IndexOf(domainPayload.SslGrade, grades)
	}

	for _, v := range domainData.Endpoints {
		if utils.IndexOf(v.Grade, grades) > savedGrade {
			savedGrade = utils.IndexOf(v.Grade, grades)
		}
	}

	newGrade := grades[savedGrade]
	if err != nil {
		domainPayload.PreviousSslGrade = newGrade
		domainPayload.SslGrade = newGrade
	} else if newGrade != domainPayload.SslGrade {
		domainPayload.PreviousSslGrade = domainPayload.SslGrade
		domainPayload.SslGrade = newGrade
	}

	servers := make([]*models.Server, 0)
	for _, endpoint := range domainData.Endpoints {
		whois := services.GetWhoisIP(endpoint.IPAddress)
		server := new(models.Server)
		server.Address = endpoint.IPAddress
		server.SslGrade = endpoint.Grade
		server.Country = strings.Join(whois.Output["Country"], ",")
		server.Owner = strings.Join(whois.Output["Organization"], ",")
		servers = append(servers, server)
	}

	title, logo := services.ExtractWebData(domain)

	fmt.Println("SSL Labbs Query Status: " + domainData.Status)
	fmt.Println("Site title: " + title)
	fmt.Println("Site logo: " + logo)

	if err != nil {
		utils.RespondWithError(w, http.StatusNoContent, "Content not found")
	} else {
		utils.RespondwithJSON(w, http.StatusOK, domainPayload)
	}
}

// Delete a domain
func (d *Domain) Delete(w http.ResponseWriter, r *http.Request) {
	domain := chi.URLParam(r, "domain")
	_, err := d.domainController.Delete(r.Context(), domain)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Server Error")
	} else {
		utils.RespondwithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
	}
}
