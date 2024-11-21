package types

// Channel
type Channel struct {
	Created   int    `json:"created"`
	Id        string `json:"id"`
	IsChannel bool   `json:"is_channel"`
	IsGroup   bool   `json:"is_group"`
	IsPrivate bool   `json:"is_private"`
	Name      string `json:"name"`
}
