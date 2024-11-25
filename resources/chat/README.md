
### post_message <a name="post_message"></a>


Sends a message to a channel.

**API Endpoint**: `POST /chat.postMessage`

#### Example Snippet

```go
import (
	sdk "github.com/Sideko-Inc/simple_slack_gen_go/client"
	os "os"
	chat "github.com/Sideko-Inc/simple_slack_gen_go/resources/chat"
	types "github.com/Sideko-Inc/simple_slack_gen_go/types"
	nullable "github.com/Sideko-Inc/simple_slack_gen_go/nullable"
)

client := sdk.NewClient(sdk.WithAuth(os.Getenv("API_TOKEN")))

res, err := client.Chat.PostMessage(chat.PostMessageRequest { Data: types.NewMessage { AsUser: nullable.NewValue("string"), Attachments: nullable.NewValue("string"), Blocks: nullable.NewValue("string"), Channel: "channel_id", IconEmoji: nullable.NewValue("string"), IconUrl: nullable.NewValue("string"), LinkNames: nullable.NewValue(true), Mrkdwn: nullable.NewValue(true), Parse: nullable.NewValue("string"), ReplyBroadcast: nullable.NewValue(true), Text: "Hello World!", ThreadTs: nullable.NewValue("string"), UnfurlLinks: nullable.NewValue(true), UnfurlMedia: nullable.NewValue(true), Username: nullable.NewValue("string") } })
```
