package models

type HTTPRequest struct{}

type HTTPResponse struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
	Status  string `json:"status"`
	Error   bool   `json:"error"`
}

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ValidateTokenRequest struct{}

type GenerateRayTokenRequest struct{}

type GetOMSTokenRequest struct{}

type PutOMSTokenRequest struct{}

type GenerateTokenRequest struct{}

type ExchangeTokenRequest struct{}

type RevokeOMSTokenRequest struct{}

type RevokeTokenRequest struct{}

type RevokeEncryptedTokenRequest struct{}

type RevokeClientTokenRequest struct{}

type RevokeSessionsRequest struct{}
