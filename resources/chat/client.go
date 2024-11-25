package chat

import (
	bytes "bytes"
	json "encoding/json"
	io "io"
	http "net/http"
	url "net/url"

	sdkcore "github.com/Sideko-Inc/simple_slack_gen_go/core"
	types "github.com/Sideko-Inc/simple_slack_gen_go/types"
)

type Client struct {
	coreClient *sdkcore.CoreClient
}
type RequestModifier = func(req *http.Request) error

// Instantiate a new resource client
func NewClient(coreClient *sdkcore.CoreClient) *Client {
	client := Client{
		coreClient: coreClient,
	}

	return &client
}

// Sends a message to a channel.
//
// POST /chat.postMessage
func (c *Client) PostMessage(request PostMessageRequest, reqModifiers ...RequestModifier) (types.MessageResponse, error) {
	// URL formatting
	joined, err := url.JoinPath(c.coreClient.BaseURL, "/chat.postMessage")
	if err != nil {
		return types.MessageResponse{}, err
	}
	url, err := url.Parse(joined)
	if err != nil {
		return types.MessageResponse{}, err
	}

	// Prep body
	reqBody, err := json.Marshal(request.Data)
	if err != nil {
		return types.MessageResponse{}, err
	}
	reqBodyBuf := bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", url.String(), reqBodyBuf)
	if err != nil {
		return types.MessageResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	c.coreClient.AddAuth([]string{"auth"}, req)

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.MessageResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.MessageResponse{}, err
	}
	defer resp.Body.Close()

	// Handle response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.MessageResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.MessageResponse{}, sdkcore.NewApiError(*req, *resp, body)
	}
	var bodyData types.MessageResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.MessageResponse{}, err
	}
	return bodyData, nil

}
