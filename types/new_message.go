package types

import (
	nullable "github.com/Sideko-Inc/simple_slack_gen_go/nullable"
)

// NewMessage
type NewMessage struct {
	// Pass true to post the message as the authed user, instead of as a bot. Defaults to false. See [authorship](#authorship) below.
	AsUser nullable.Nullable[string] `json:"as_user,omitempty"`
	// A JSON-based array of structured attachments, presented as a URL-encoded string.
	Attachments nullable.Nullable[string] `json:"attachments,omitempty"`
	// A JSON-based array of structured blocks, presented as a URL-encoded string.
	Blocks nullable.Nullable[string] `json:"blocks,omitempty"`
	// Channel, private group, or IM channel to send message to. Can be an encoded ID, or a name. See [below](#channels) for more details.
	Channel string `json:"channel"`
	// Emoji to use as the icon for this message. Overrides `icon_url`. Must be used in conjunction with `as_user` set to `false`, otherwise ignored. See [authorship](#authorship) below.
	IconEmoji nullable.Nullable[string] `json:"icon_emoji,omitempty"`
	// URL to an image to use as the icon for this message. Must be used in conjunction with `as_user` set to false, otherwise ignored. See [authorship](#authorship) below.
	IconUrl nullable.Nullable[string] `json:"icon_url,omitempty"`
	// Find and link channel names and usernames.
	LinkNames nullable.Nullable[bool] `json:"link_names,omitempty"`
	// Disable Slack markup parsing by setting to `false`. Enabled by default.
	Mrkdwn nullable.Nullable[bool] `json:"mrkdwn,omitempty"`
	// Change how messages are treated. Defaults to `none`. See [below](#formatting).
	Parse nullable.Nullable[string] `json:"parse,omitempty"`
	// Used in conjunction with `thread_ts` and indicates whether reply should be made visible to everyone in the channel or conversation. Defaults to `false`.
	ReplyBroadcast nullable.Nullable[bool] `json:"reply_broadcast,omitempty"`
	// How this field works and whether it is required depends on other fields you use in your API call. [See below](#text_usage) for more detail.
	Text string `json:"text"`
	// Provide another message's `ts` value to make this message a reply. Avoid using a reply's `ts` value; use its parent instead.
	ThreadTs nullable.Nullable[string] `json:"thread_ts,omitempty"`
	// Pass true to enable unfurling of primarily text-based content.
	UnfurlLinks nullable.Nullable[bool] `json:"unfurl_links,omitempty"`
	// Pass false to disable unfurling of media content.
	UnfurlMedia nullable.Nullable[bool] `json:"unfurl_media,omitempty"`
	// Set your bot's user name. Must be used in conjunction with `as_user` set to false, otherwise ignored. See [authorship](#authorship) below.
	Username nullable.Nullable[string] `json:"username,omitempty"`
}
