package room

import (
	"SyncCinemaNet_WebGo/auth"
	info_of "SyncCinemaNet_WebGo/chat"
)

type Room struct {
	ID       int64
	Name     string
	Private  bool
	Chat     info_of.Chat
	Author   auth.User
	userlist map[string]auth.User
	Tags     map[string]string
}
