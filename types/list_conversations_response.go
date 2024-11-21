package types

// ListConversationsResponse
type ListConversationsResponse struct {
	Channels []Channel `json:"channels"`
	Ok       bool      `json:"ok"`
}
