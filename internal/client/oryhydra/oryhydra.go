package oryhydra

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"open-api/internal/logger"
	"open-api/internal/models"
	"open-api/internal/utils"
	"strings"
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

func (o *OryHydra) CreateApp(ctx context.Context, oAuthClientID string, oAuthClientName string, clientSecret string, redirectUris []string, scopes []string, grantTypes []string) (*models.HTTPResponse, error) {
	Logger := logger.CreateFileLoggerWithCtx(ctx)

	apiUrl := utils.BuildURL(o.BaseURL, "/clients", nil)
	method := http.MethodPost
	headers := map[string]string{
		"Content-Type": "text/plain",
	}

	reqBody := map[string]interface{}{
		"client_id":     oAuthClientID,
		"client_secret": clientSecret,
		"client_name":   oAuthClientName,
		"scopes":        strings.Join(scopes, " "),
		"grant_types":   grantTypes,
		"redirect_uris": redirectUris,
	}

	Logger.Infof("oryhydra/CreateApp: http request info, method: %v, url, %v, headers: %v, req_body: %v", method, apiUrl, headers, reqBody)
	statusCode, respBodyBytes, err := utils.MakeHttpRequest(ctx, o.HttpClient, method, apiUrl, reqBody, headers)
	if err != nil {
		return nil, err
	}
	Logger.Infof("oryhydra/CreateApp: http response received, status_code: %+v, response: %v.", statusCode, string(respBodyBytes)) //we cannot print the response received since it has details of the stocks the clients holds

	if !utils.Is200(statusCode) {
		return nil, errors.New("error signing up")
	}

	resp := models.HTTPResponse{}
	err = json.Unmarshal(respBodyBytes, &resp)
	if err != nil {
		Logger.Errorw("http response unmarshalling failed", "error", err, "http_response", string(respBodyBytes))
		return nil, err
	}

	return &resp, nil
}
