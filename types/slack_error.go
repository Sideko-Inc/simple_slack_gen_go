package types

// SlackError
type SlackError struct {
	// error type
	Error string `json:"error"`
	Ok    bool   `json:"ok"`
}
