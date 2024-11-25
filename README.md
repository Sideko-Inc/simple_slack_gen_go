
# Slack Go SDK

## Overview
One way to interact with the Slack platform is its HTTP RPC-based Web API, a collection of methods requiring OAuth 2.0-based user, bot, or workspace tokens blessed with related OAuth scopes.

### Example Client Initialization

```go
import (
	sdk "github.com/Sideko-Inc/simple_slack_gen_go/client"
	os "os"
)

client := sdk.NewClient(sdk.WithAuth(os.Getenv("API_TOKEN")))
```

## Module Documentation and Snippets

### [chat](resources/chat/README.md)

* [post_message](resources/chat/README.md#post_message) - 

### [conversations](resources/conversations/README.md)

* [list](resources/conversations/README.md#list) - 

<!-- MODULE DOCS END -->
