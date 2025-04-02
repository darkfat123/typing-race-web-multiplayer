package models

import (
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
