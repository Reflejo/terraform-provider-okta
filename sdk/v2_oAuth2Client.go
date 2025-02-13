package sdk

type OAuth2Client struct {
	Links      interface{} `json:"_links,omitempty"`
	ClientId   string      `json:"client_id,omitempty"`
	ClientName string      `json:"client_name,omitempty"`
	ClientUri  string      `json:"client_uri,omitempty"`
	LogoUri    string      `json:"logo_uri,omitempty"`
}
