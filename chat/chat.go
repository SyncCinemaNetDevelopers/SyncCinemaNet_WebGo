package chat

import (
	"SyncCinemaNet_WebGo/auth"
	"time"
)

type Chat struct {
	ID int64
}

type Message struct {
	ID      int
	Chat_ID Chat
	Time    time.Time
	Author  auth.User
	Text    string
}
