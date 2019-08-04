package chat

import (
	"SyncCinemaNet_WebGo/auth"
	"SyncCinemaNet_WebGo/room"
)

type Chat struct {
	ID   int64
	Room room.Room
}

type Message struct {
	ID      int
	Chat_ID Chat
	Time    string
	Author  auth.User
	Text    string
}
