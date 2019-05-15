package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/jeanbenitez/servercheck/controllers"
	"github.com/jeanbenitez/servercheck/interfaces"
	"github.com/jeanbenitez/servercheck/models"
	"github.com/jeanbenitez/servercheck/utils"
)

// NewDomainHandler ...
func NewDomainHandler(db *sql.DB) *Domain {
	return &Domain{
		controller: controllers.NewSQLDomain(db),
	}
}

// Domain ...
type Domain struct {
	controller interfaces.IDomainController
}

// Fetch all domain data
func (d *Domain) Fetch(w http.ResponseWriter, r *http.Request) {
	payload, _ := d.controller.Fetch(r.Context(), 5)

	utils.RespondwithJSON(w, http.StatusOK, payload)
}

// Create a new domain
func (d *Domain) Create(w http.ResponseWriter, r *http.Request) {
	domain := models.Domain{}
	json.NewDecoder(r.Body).Decode(&domain)

	created, err := d.controller.Create(r.Context(), &domain)
	fmt.Println(created)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	utils.RespondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
}

// Update a domain by domain
func (d *Domain) Update(w http.ResponseWriter, r *http.Request) {
	domain := chi.URLParam(r, "domain")
	data := models.Domain{Domain: string(domain)}
	json.NewDecoder(r.Body).Decode(&data)
	payload, err := d.controller.Update(r.Context(), &data)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	utils.RespondwithJSON(w, http.StatusOK, payload)
}

// GetByDomain returns a domain details
func (d *Domain) GetByDomain(w http.ResponseWriter, r *http.Request) {
	domain := chi.URLParam(r, "domain")
	payload, err := d.controller.GetByDomain(r.Context(), domain)

	if err != nil {
		utils.RespondWithError(w, http.StatusNoContent, "Content not found")
	}

	utils.RespondwithJSON(w, http.StatusOK, payload)
}

// Delete a domain
func (d *Domain) Delete(w http.ResponseWriter, r *http.Request) {
	domain := chi.URLParam(r, "domain")
	_, err := d.controller.Delete(r.Context(), domain)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	utils.RespondwithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
}
