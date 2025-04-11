package model

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Player struct {
	Conn      *websocket.Conn
	Username  string
	StartTime time.Time
	WordCount int
	Finished  bool
	Ready     bool
}

type Room struct {
	ID           string
	Language     string
	Players      map[*websocket.Conn]*Player
	Text         string
	Mutex        sync.Mutex
	Locked       bool
	Limit        int
	RestartVotes map[string]bool
}
