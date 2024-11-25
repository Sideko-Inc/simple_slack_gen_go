package test_conversations_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/Sideko-Inc/simple_slack_gen_go/client"
	conversations "github.com/Sideko-Inc/simple_slack_gen_go/resources/conversations"
)

func TestList200GeneratedSuccess(t *testing.T) {
	// Success case test
	client := sdk.NewClient(sdk.WithAuth("API_TOKEN"), sdk.WithBaseURL("https://api.sideko-staging.dev/v1/mock/demo/simple-slack/0.2.0"))
	res, err := client.Conversations.List(conversations.ListRequest{})

	if err != nil {
		t.Fatalf("TestList200GeneratedSuccess - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
