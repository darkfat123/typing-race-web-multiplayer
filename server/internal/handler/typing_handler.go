package handler

import (
	"log"
	"net/http"
	"server/internal/logic"
	"server/internal/model"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

func HandleTypingWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	var username, roomID, language string
	var data map[string]string
	if err := conn.ReadJSON(&data); err != nil {
		log.Println("Error reading username & roomID:", err)
		return
	}
	username = data["username"]
	roomID = data["roomID"]
	language = data["language"]

	room := logic.GetOrCreateRoom(roomID, language)
	player := &model.Player{Conn: conn, Username: username, StartTime: time.Time{}, Finished: false, Ready: false}

	room.Mutex.Lock()
	room.Players[conn] = player
	room.Mutex.Unlock()

	// Send text to player
	if err := conn.WriteJSON(map[string]string{"text": room.Text}); err != nil {
		log.Println("Error sending text:", err)
	}

	log.Printf("Player %s has joined room %s", player.Username, roomID)

	logic.UpdateUserList(room)
	logic.UpdateReadyStatus(room)

	for {
		var message map[string]string
		err := conn.ReadJSON(&message)
		if err != nil {
			if websocket.IsCloseError(err) {
				log.Println("WebSocket closed gracefully:", err)
			} else {
				log.Println("Error reading message:", err)
			}
			break
		}

		if message["type"] == "close" {
			log.Printf("%s is leaving from Room No. %s, closing connection", player.Username, roomID)
			conn.Close()
			break
		}

		if status, ok := message["status"]; ok && status == "ready" {
			room.Mutex.Lock()
			player.Ready = true
			room.Mutex.Unlock()

			log.Printf("Player %s in room %s is ready", player.Username, roomID)
			logic.UpdateReadyStatus(room)

			if logic.IsAllPlayersReady(room) {
				log.Printf("All players in room %s are ready. Starting the game!", roomID)
				logic.Broadcast(room, map[string]string{"type": "start_game"})
			}
		}

		if text, ok := message["text"]; ok {
			if strings.TrimSpace(text) == room.Text && !player.Finished {
				player.WordCount = len(strings.Fields(text))
				player.Finished = true
				wpm := logic.CalculateWPM(player)

				log.Printf("Player %s in room %s has finished typing with WPM: %.2f", player.Username, roomID, wpm)

				logic.Broadcast(room, map[string]interface{}{
					"type":     "finished",
					"username": player.Username,
					"wpm":      wpm,
				})
			}
		}
	}

	// Cleanup when player leaves
	logic.CleanupPlayer(room, conn, roomID)
}
