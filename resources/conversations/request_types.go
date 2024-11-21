package conversations

import (
	nullable "github.com/Sideko-Inc/simple_slack_gen_go/nullable"
)

// ListRequest
type ListRequest struct {
	Cursor          nullable.Nullable[string] `json:"cursor,omitempty"`
	ExcludeArchived nullable.Nullable[bool]   `json:"excludeArchived,omitempty"`
	Limit           nullable.Nullable[int]    `json:"limit,omitempty"`
	TeamId          nullable.Nullable[string] `json:"teamId,omitempty"`
	Types           nullable.Nullable[string] `json:"types,omitempty"`
}
