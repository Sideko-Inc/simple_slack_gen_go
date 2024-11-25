package test_chat_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/Sideko-Inc/simple_slack_gen_go/client"
	nullable "github.com/Sideko-Inc/simple_slack_gen_go/nullable"
	chat "github.com/Sideko-Inc/simple_slack_gen_go/resources/chat"
	types "github.com/Sideko-Inc/simple_slack_gen_go/types"
)

func TestPostMessage200SuccessDefault(t *testing.T) {
	// Success test for Default body
	client := sdk.NewClient(sdk.WithAuth("API_TOKEN"), sdk.WithBaseURL("https://api.sideko-staging.dev/v1/mock/demo/simple-slack/0.2.0"))
	res, err := client.Chat.PostMessage(chat.PostMessageRequest{Data: types.NewMessage{AsUser: nullable.NewValue("string"), Attachments: nullable.NewValue("string"), Blocks: nullable.NewValue("string"), Channel: "channel_id", IconEmoji: nullable.NewValue("string"), IconUrl: nullable.NewValue("string"), LinkNames: nullable.NewValue(true), Mrkdwn: nullable.NewValue(true), Parse: nullable.NewValue("string"), ReplyBroadcast: nullable.NewValue(true), Text: "Hello World!", ThreadTs: nullable.NewValue("string"), UnfurlLinks: nullable.NewValue(true), UnfurlMedia: nullable.NewValue(true), Username: nullable.NewValue("string")}})

	if err != nil {
		t.Fatalf("TestPostMessage200SuccessDefault - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
