package chat

import (
	types "github.com/Sideko-Inc/simple_slack_gen_go/types"
)

// PostMessageRequest
type PostMessageRequest struct {
	Data types.NewMessage `json:"data"`
}
