package models

// SslLabs model
type SslLabs struct {
	Host      string `json:"host"`
	Status    string `json:"status"`
	Endpoints []struct {
		IPAddress  string `json:"ipAddress"`
		ServerName string `json:"serverName"`
		Grade      string `json:"grade"`
	} `json:"endpoints"`
}
