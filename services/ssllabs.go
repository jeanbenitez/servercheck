package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jeanbenitez/servercheck/models"
)

// GetSslLabsDomainData return domain data
func GetSslLabsDomainData(domain string) (domainData models.SslLabsDomainDataResponse) {
	response, err := http.Get("https://api.ssllabs.com/api/v3/analyze?host=" + domain)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	decoder := json.NewDecoder(response.Body)
	err2 := decoder.Decode(&domainData)
	if err2 != nil {
		panic(err)
	}
	return domainData
}
