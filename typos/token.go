package typos

// MakeTokenResponse -
type MakeTokenResponse struct {
	Token     string `json:"access_token"`
	TokenType string `json:"token_type"`
	Scope     string `json:"scope,omitempty"`
	Expire    int64  `json:"expires_in,omitempty"`
}

// TokenRequest -
type TokenRequest struct {
	GrantType string
	Scope     string
}

// TokenResponse -
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
	Scope       string `json:"scope,omitempty"`
}
