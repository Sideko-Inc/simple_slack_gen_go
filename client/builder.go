package client

import (
	http "net/http"

	sdkcore "github.com/Sideko-Inc/simple_slack_gen_go/core"
)

// Provide your own http.Client to be used for all requests
func WithHTTPClient(httpClient *http.Client) func(*sdkcore.CoreClient) {
	return func(c *sdkcore.CoreClient) {
		c.HttpClient = httpClient
	}
}

// Provide non-default baseURL for all requests
func WithBaseURL(baseURL string) func(*sdkcore.CoreClient) {
	return func(c *sdkcore.CoreClient) {
		c.BaseURL = baseURL
	}
}

type RequestModifier = func(req *http.Request) error

// Provide modifiers to be applied to all client requests
func WithModifiers(modifiers ...RequestModifier) func(*sdkcore.CoreClient) {
	return func(c *sdkcore.CoreClient) {
		c.Modifiers = append(c.Modifiers, modifiers...)
	}
}

func WithAuth(val string) func(*sdkcore.CoreClient) {
	return func(c *sdkcore.CoreClient) {
		c.Auth["auth"] = func(request *http.Request) {
			request.Header.Add("Authorization", "Bearer "+val)
		}
	}
}
