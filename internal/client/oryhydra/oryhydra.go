package oryhydra

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"open-api-client/internal/logger"
	"open-api-client/internal/models"
	"open-api-client/internal/utils"
)

type IOryHydra interface {
	CreateApp(ctx context.Context, oauthClientID string, oAuthClientName string, clientSecret string, redirectUris []string, scope string, grantTypes []string)
}

type OryHydra struct {
	HttpClient *http.Client
	BaseURL    *url.URL
}

func NewOryHydra(client *http.Client, baseURL *url.URL) *OryHydra {
	return &OryHydra{HttpClient: client, BaseURL: baseURL}
}

func (o *OryHydra) CreateApp(ctx context.Context, createOauthClientReq *models.CreateOauthClientRequest) (*models.CreateOauthClientResponse, error) {
	Logger := logger.CreateFileLoggerWithCtx(ctx)

	apiUrl := utils.BuildURL(o.BaseURL, "/clients", nil)
	method := http.MethodPost
	headers := map[string]string{
		"Content-Type": "text/plain",
	}

	Logger.Infof("oryhydra/CreateApp: http request info, method: %v, url, %v, headers: %v, req_body: %v", method, apiUrl, headers, createOauthClientReq)
	statusCode, respBodyBytes, err := utils.MakeHttpRequest(ctx, o.HttpClient, method, apiUrl, createOauthClientReq, headers)
	if err != nil {
		return nil, err
	}
	Logger.Infof("oryhydra/CreateApp: http response received, status_code: %+v, response: %v.", statusCode, string(respBodyBytes)) //we cannot print the response received since it has details of the stocks the clients holds

	if !utils.Is200(statusCode) {
		return nil, errors.New("error signing up")
	}

	resp := models.CreateOauthClientResponse{}
	err = json.Unmarshal(respBodyBytes, &resp)
	if err != nil {
		Logger.Errorw("http response unmarshalling failed", "error", err, "http_response", string(respBodyBytes))
		return nil, err
	}

	return &resp, nil
}
