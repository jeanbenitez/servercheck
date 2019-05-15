package models

// Domain model
type Domain struct {
	Domain           string `json:"domain"`
	ServersChanged   string `json:"servers_changed"`
	SslGrade         string `json:"ssl_grade"`
	PreviousSslGrade string `json:"previous_ssl_grade"`
	Logo             string `json:"logo"`
	IsDown           string `json:"is_down"`
	Servers          []struct {
		Address  string `json:"address"`
		SslGrade string `json:"ssl_grade"`
		Country  string `json:"country"`
		Owner    string `json:"owner"`
	} `json:"servers"`
}
