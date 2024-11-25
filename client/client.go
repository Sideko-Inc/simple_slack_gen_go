package client

import (
	sdkcore "github.com/Sideko-Inc/simple_slack_gen_go/core"
	chat "github.com/Sideko-Inc/simple_slack_gen_go/resources/chat"
	conversations "github.com/Sideko-Inc/simple_slack_gen_go/resources/conversations"
)

type Client struct {
	coreClient    *sdkcore.CoreClient
	Conversations *conversations.Client
	Chat          *chat.Client
}

// Instantiate a new API client
func NewClient(builders ...func(*sdkcore.CoreClient)) *Client {
	baseURL := `https://slack.com/api`
	coreClient := sdkcore.NewCoreClient(baseURL)
	for _, b := range builders {
		b(coreClient)
	}

	client := Client{
		coreClient:    coreClient,
		Conversations: conversations.NewClient(coreClient),
		Chat:          chat.NewClient(coreClient),
	}

	return &client
}
