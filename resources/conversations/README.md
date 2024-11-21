
### list <a name="list"></a>


Lists channels in the workspace.

**API Endpoint**: `GET /conversations.list`

#### Example Snippet

```go
import (
	sdk "github.com/Sideko-Inc/simple_slack_gen_go/client"
	os "os"
	conversations "github.com/Sideko-Inc/simple_slack_gen_go/resources/conversations"
)

client := sdk.NewClient(sdk.WithAuth(os.Getenv("API_TOKEN")))

res, err := client.Conversations.List(conversations.ListRequest {  })
```
