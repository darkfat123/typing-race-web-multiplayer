package handler

import (
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		allowedOrigin := os.Getenv("ALLOWED_ORIGIN")

		return origin == allowedOrigin
	},
}
