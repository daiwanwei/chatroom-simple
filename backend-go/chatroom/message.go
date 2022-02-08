package chatroom

import "time"

type Message struct {
	Name    string    `json:"name"`
	Message string    `json:"message"`
	When    time.Time `json:"when"`
}
