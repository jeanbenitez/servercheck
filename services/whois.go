package services

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jeanbenitez/servercheck/models"
)

// GetWhois return whois domain
func GetWhois(domain string) (whoisData models.WhoisResponse) {
	username := "420205657"
	passwd := "KQB2Yr3uAJlvp4XzwBrDwg"
	url := "https://jsonwhoisapi.com/api/v1/whois?identifier=" + domain

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(username, passwd)
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(response.Body)
	err2 := decoder.Decode(&whoisData)
	if err2 != nil {
		log.Fatal(err2)
	}
	return whoisData
}
