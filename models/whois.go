package models

// WhoisResponse data model
type WhoisResponse struct {
	Contacts struct {
		Owner []struct {
			Name         string `json:"name"`
			Organization string `json:"organization"`
			Country      string `json:"country"`
		} `json:"owner"`
	} `json:"contacts"`
}
