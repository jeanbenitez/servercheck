package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	// for test purpose
	body, err := ioutil.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)

	err2 := json.Unmarshal(body, &domainData)
	if err2 != nil {
		panic(err)
	}
	return domainData
}
