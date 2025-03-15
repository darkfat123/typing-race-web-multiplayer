package models

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Room struct {
	Players map[*websocket.Conn]*Player
	Text    string
	Mutex   sync.Mutex
	Locked  bool
}
