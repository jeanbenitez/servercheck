package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	domain "github.com/jeanbenitez/servercheck/controllers/domain"
	interfaces "github.com/jeanbenitez/servercheck/interfaces"
	models "github.com/jeanbenitez/servercheck/models"
)

// NewDomainHandler ...
func NewDomainHandler(db *sql.DB) *Domain {
	return &Domain{
		controller: domain.NewSQLDomain(db),
	}
}

// Domain ...
type Domain struct {
	controller interfaces.IDomainController
}

// Fetch all domain data
func (d *Domain) Fetch(w http.ResponseWriter, r *http.Request) {
	payload, _ := d.controller.Fetch(r.Context(), 5)

	respondwithJSON(w, http.StatusOK, payload)
}

// Create a new domain
func (d *Domain) Create(w http.ResponseWriter, r *http.Request) {
	domain := models.Domain{}
	json.NewDecoder(r.Body).Decode(&domain)

	created, err := d.controller.Create(r.Context(), &domain)
	fmt.Println(created)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
}

// Update a domain by domain
func (d *Domain) Update(w http.ResponseWriter, r *http.Request) {
	domain := chi.URLParam(r, "domain")
	data := models.Domain{Domain: string(domain)}
	json.NewDecoder(r.Body).Decode(&data)
	payload, err := d.controller.Update(r.Context(), &data)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// GetByDomain returns a domain details
func (d *Domain) GetByDomain(w http.ResponseWriter, r *http.Request) {
	domain := chi.URLParam(r, "domain")
	payload, err := d.controller.GetByDomain(r.Context(), domain)

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// Delete a domain
func (d *Domain) Delete(w http.ResponseWriter, r *http.Request) {
	domain := chi.URLParam(r, "domain")
	_, err := d.controller.Delete(r.Context(), domain)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
