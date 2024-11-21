package conversations

import (
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

// Lists channels in the workspace.
//
// GET /conversations.list
func (c *Client) List(request ListRequest, reqModifiers ...RequestModifier) (types.ListConversationsResponse, error) {
	// URL formatting
	joined, err := url.JoinPath(c.coreClient.BaseURL, "/conversations.list")
	if err != nil {
		return types.ListConversationsResponse{}, err
	}
	url, err := url.Parse(joined)
	if err != nil {
		return types.ListConversationsResponse{}, err
	}

	// Query params
	params := url.Query()
	cursor, err := request.Cursor.Value()
	if err == nil {
		sdkcore.AddQueryParam(params, "cursor", cursor, false)
	}
	excludeArchived, err := request.ExcludeArchived.Value()
	if err == nil {
		sdkcore.AddQueryParam(params, "exclude_archived", excludeArchived, false)
	}
	limit, err := request.Limit.Value()
	if err == nil {
		sdkcore.AddQueryParam(params, "limit", limit, false)
	}
	teamId, err := request.TeamId.Value()
	if err == nil {
		sdkcore.AddQueryParam(params, "team_id", teamId, false)
	}
	q_types, err := request.Types.Value()
	if err == nil {
		sdkcore.AddQueryParam(params, "types", q_types, false)
	}
	url.RawQuery = params.Encode()

	// Init request
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return types.ListConversationsResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")

	// Add auth
	c.coreClient.AddAuth([]string{"auth"}, req)

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.ListConversationsResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.ListConversationsResponse{}, err
	}
	defer resp.Body.Close()

	// Handle response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.ListConversationsResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.ListConversationsResponse{}, sdkcore.NewApiError(*req, *resp, body)
	}
	var bodyData types.ListConversationsResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.ListConversationsResponse{}, err
	}
	return bodyData, nil

}
