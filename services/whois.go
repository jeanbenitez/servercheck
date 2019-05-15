package services

import (
	"log"
	"net"

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

// GetWhoisIP return whois IP
func GetWhoisIP(ip string) (result *whois.Result) {
	result, err := whois.QueryIP(net.ParseIP(ip))
	if err != nil {
		log.Fatal(err)
	}
	return result
}
