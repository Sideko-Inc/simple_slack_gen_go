package types

import (
	nullable "github.com/Sideko-Inc/simple_slack_gen_go/nullable"
)

// Message
type Message struct {
	AppId       nullable.Nullable[string]                   `json:"app_id,omitempty"`
	Attachments nullable.Nullable[[]map[string]interface{}] `json:"attachments,omitempty"`
	Blocks      nullable.Nullable[[]map[string]interface{}] `json:"blocks,omitempty"`
	BotId       nullable.Nullable[string]                   `json:"bot_id,omitempty"`
	BotProfile  nullable.Nullable[map[string]interface{}]   `json:"bot_profile,omitempty"`
	Team        nullable.Nullable[string]                   `json:"team,omitempty"`
	Text        string                                      `json:"text"`
	Ts          string                                      `json:"ts"`
	Type        string                                      `json:"type"`
	User        nullable.Nullable[string]                   `json:"user,omitempty"`
}
