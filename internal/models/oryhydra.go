package models

import "time"

type CreateOauthClientRequest struct {
	ClientId     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	ClientName   string   `json:"client_name"`
	Scopes       string   `json:"scopes"`
	GrantTypes   []string `json:"grant_types"`
	RedirectUris []string `json:"redirect_uris"`
}

type CreateOauthClientResponse struct {
	ClientId                  string        `json:"client_id"`
	ClientName                string        `json:"client_name"`
	ClientSecret              string        `json:"client_secret"`
	RedirectUris              []string      `json:"redirect_uris"`
	GrantTypes                []string      `json:"grant_types"`
	ResponseTypes             interface{}   `json:"response_types"`
	Scope                     string        `json:"scope"`
	Audience                  []interface{} `json:"audience"`
	Owner                     string        `json:"owner"`
	PolicyUri                 string        `json:"policy_uri"`
	AllowedCorsOrigins        []interface{} `json:"allowed_cors_origins"`
	TosUri                    string        `json:"tos_uri"`
	ClientUri                 string        `json:"client_uri"`
	LogoUri                   string        `json:"logo_uri"`
	Contacts                  interface{}   `json:"contacts"`
	ClientSecretExpiresAt     int           `json:"client_secret_expires_at"`
	SubjectType               string        `json:"subject_type"`
	Jwks                      struct{}      `json:"jwks"`
	TokenEndpointAuthMethod   string        `json:"token_endpoint_auth_method"`
	UserinfoSignedResponseAlg string        `json:"userinfo_signed_response_alg"`
	CreatedAt                 time.Time     `json:"created_at"`
	UpdatedAt                 time.Time     `json:"updated_at"`
	Metadata                  struct{}      `json:"metadata"`
	RegistrationAccessToken   string        `json:"registration_access_token"`
	RegistrationClientUri     string        `json:"registration_client_uri"`
}
