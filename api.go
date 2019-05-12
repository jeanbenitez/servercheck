package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// SslLabsResponse data model
type SslLabsResponse struct {
	Host      string `json:"host"`
	Status    string `json:"status"`
	Endpoints []struct {
		IPAddress  string `json:"ipAddress"`
		ServerName string `json:"serverName"`
		Grade      string `json:"grade"`
	} `json:"endpoints"`
}

func getDomainDataInSllLabs(domain string) (domainData SslLabsResponse) {
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

// WhoisResponse data model
type WhoisResponse struct {
	Changed  string `json:"changed"`
	Contacts struct {
		Owner []struct {
			Handle       interface{} `json:"handle"`
			Type         interface{} `json:"type"`
			Name         string      `json:"name"`
			Organization string      `json:"organization"`
			Country      string      `json:"country"`
		} `json:"owner"`
	} `json:"contacts"`
}

func getWhoisDomain(domain string) (whoisData WhoisResponse) {
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
