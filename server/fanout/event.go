package fanout

const (
	// EventTypeWebP is used to signal what type of message we are sending over
	// the socket.
	EventTypeWebP = "webp"

	// EventTypeErr is used to signal there was an error encountered rendering
	// the WebP image.
	EventTypeErr = "error"
)

// WebsocketEvent is a structure used to send messages over the socket.
type WebsocketEvent struct {
	// Message is the contents of the message. This is a webp, base64 encoded.
	Message string `json:"message"`

	// Type is the type of message we are sending over the socket.
	Type string `json:"type"`
}
