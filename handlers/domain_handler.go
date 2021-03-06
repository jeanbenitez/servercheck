package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
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

// FetchDomains ...
type FetchDomains struct {
	Items []*models.Domain `json:"items"`
}

// Fetch all domain data
func (d *Domain) Fetch(w http.ResponseWriter, r *http.Request) {
	domainsPayload, err := d.domainController.Fetch(r.Context(), 100)
	if err != nil {
		fmt.Print(err.Error())
	}

	for _, domain := range domainsPayload {
		if domain != nil {
			serversPayload, err2 := d.serverController.FetchByDomain(r.Context(), domain.Domain)
			if err2 == nil {
				for _, server := range serversPayload {
					if server != nil {
						domain.Servers = append(domain.Servers, *server)
					}
				}
			}
		}
	}
	utils.RespondwithJSON(w, http.StatusOK, FetchDomains{Items: domainsPayload})
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
		domainPayload.Domain = domain
	}

	title, logo := services.ExtractWebData(domain)
	domainPayload.Logo = logo
	domainPayload.Title = title
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

	serversPayload, err2 := d.serverController.FetchByDomain(r.Context(), domain)
	d.serverController.Delete(r.Context(), domain)

	domainPayload.ServersChanged = len(domainPayload.Servers) != len(domainData.Endpoints)

	servers := make([]*models.Server, len(domainData.Endpoints))
	for _, endpoint := range domainData.Endpoints {
		whois := services.GetWhoisIP(endpoint.IPAddress)
		server := new(models.Server)
		server.Address = endpoint.IPAddress
		server.SslGrade = endpoint.Grade
		server.Country = strings.Join(whois.Output["Country"], ",")
		server.Owner = strings.Join(whois.Output["Organization"], ",")

		if !domainPayload.ServersChanged {
			var savedServer *models.Server
			if err2 == nil {
				for i, s := range serversPayload {
					if strings.Compare(s.Address, endpoint.IPAddress) == 0 {
						savedServer = serversPayload[i]
						break
					}
				}
			}

			if savedServer != nil {
				domainPayload.ServersChanged = reflect.DeepEqual(server, savedServer)
			}
		}

		servers = append(servers, server)
		d.serverController.Create(r.Context(), domain, server)
	}

	fmt.Println("SSL Labbs Query Status: " + domainData.Status)

	domainPayload.IsDown = domainData.Status == "ERROR"

	if err != nil {
		d.domainController.Create(r.Context(), domainPayload)
	} else {
		domainPayload, _ = d.domainController.Update(r.Context(), domainPayload)
	}
	for _, s := range servers {
		if s != nil {
			domainPayload.Servers = append(domainPayload.Servers, *s)
		}
	}

	utils.RespondwithJSON(w, http.StatusOK, domainPayload)
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
