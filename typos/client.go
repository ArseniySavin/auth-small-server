package typos

// ClientRequest -
type ClientRequest struct {
	ClientId     string             `json:"client_id"`
	ClientName   string             `json:"client_name"`
	ClientSecret ClientSecretBase64 `json:"client_secret"`
	GrantTypes   []string           `json:"grant_types"`
	Scope        ClientScope        `json:"scope"`
	Data         ClientData         `json:"data"`
}

type ClientData string
type ClientSecretBase64 string
type ClientScope []string
