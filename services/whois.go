package services

import (
	"log"

	"github.com/xellio/whois"
)

// GetWhois return whois domain
func GetWhois(domain string) (result *whois.Result) {
	result, err := whois.QueryHost(domain)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
