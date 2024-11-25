package types

import (
	nullable "github.com/Sideko-Inc/simple_slack_gen_go/nullable"
)

// Schema for successful response of chat.postMessage method
type MessageResponse struct {
	Channel          string                                    `json:"channel"`
	Message          Message                                   `json:"message"`
	Ok               bool                                      `json:"ok"`
	ResponseMetadata nullable.Nullable[map[string]interface{}] `json:"response_metadata,omitempty"`
	Ts               string                                    `json:"ts"`
	Warning          nullable.Nullable[string]                 `json:"warning,omitempty"`
}
