package main

import (
	"fmt"
	"log"
	"net/http"
	"server/internal/handler"
)

func main() {
	http.HandleFunc("/ws/lobby", handler.HandleLobbyWebSocket)
	http.HandleFunc("/ws/typing", handler.HandleTypingWebSocket)
	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
