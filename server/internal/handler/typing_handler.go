package handler

import (
	"log"
	"net/http"
	"server/internal/logic"
	"server/internal/model"
	"strconv"
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

	var username, roomIDInput, language string
	var data map[string]string
	if err := conn.ReadJSON(&data); err != nil {
		log.Println("Error reading username & roomID:", err)
		return
	}
	username = data["username"]
	roomIDInput = data["roomID"]
	language = data["language"]

	room := logic.GetOrCreateRoom(roomIDInput, language)

	if room.Limit > 0 && len(room.Players) >= room.Limit {
		log.Printf("Room is %s full. Rejecting %s.", room.ID, username)
		conn.WriteJSON(map[string]string{"error": "Room is full"})
		conn.Close()
		return
	}

	if limitStr, ok := data["limit"]; ok && room.Limit == 0 {
		if limit, err := strconv.Atoi(limitStr); err == nil {
			room.Limit = limit
		}
	}

	player := &model.Player{Conn: conn, Username: username, StartTime: time.Time{}, Finished: false, Ready: false}
	JoinRoom(room.ID, player.Username)

	room.Mutex.Lock()
	room.Players[conn] = player
	room.Mutex.Unlock()

	if err := conn.WriteJSON(map[string]string{"text": room.Text}); err != nil {
		log.Println("Error sending text:", err)
	}

	log.Printf("Player %s has joined room %s (Max players: %d, Now: %d)", player.Username, room.ID, room.Limit, len(room.Players))

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

		if message["type"] == "vote_restart" {
			room.RestartVotes[player.Username] = true

			if float64(len(room.RestartVotes))/float64(len(room.Players)) > 0.6 {
				room.RestartVotes = make(map[string]bool)
				newText := logic.GetRandomText(room.Language)
				room.Text = newText

				logic.Broadcast(room, map[string]interface{}{
					"type": "restart_game",
					"text": newText,
				})
				player.StartTime = time.Now()
				player.Finished = false

			} else {
				logic.Broadcast(room, map[string]interface{}{
					"type":  "update_votes",
					"votes": len(room.RestartVotes),
					"total": len(room.Players),
				})
			}
		}

		if message["type"] == "close" {
			log.Printf("%s is leaving from Room No. %s, closing connection", player.Username, room.ID)
			conn.Close()
			break
		}

		if status, ok := message["status"]; ok && (status == "ready" || status == "not_ready") {
			room.Mutex.Lock()
			if status == "ready" {
				player.Ready = true
			} else {
				player.Ready = false
			}
			room.Mutex.Unlock()

			log.Printf("Player %s in room %s is %s", player.Username, room.ID, status)
			logic.UpdateReadyStatus(room)

			if logic.IsAllPlayersReady(room) {
				log.Printf("All players in room %s are ready. Starting the game!", room.ID)
				logic.Broadcast(room, map[string]string{"type": "start_game"})
			}
		}

		if text, ok := message["text"]; ok {
			player.WordCount = len(strings.Fields(text))
			wpm := logic.CalculateWPM(player)
			if strings.TrimSpace(text) == room.Text && !player.Finished {
				player.Finished = true

				log.Printf("Player %s in room %s has finished typing with WPM: %.2f", player.Username, room.ID, wpm)

				logic.Broadcast(room, map[string]interface{}{
					"type":     "finished",
					"username": player.Username,
					"wpm":      wpm,
				})
			} else {
				logic.Broadcast(room, map[string]interface{}{
					"type":     "calculate_wpm",
					"username": player.Username,
					"wpm":      wpm,
				})
			}
		}

	}

	RemoveUserFromRoom(room.ID, player.Username)
	logic.CleanupPlayer(room, conn, room.ID)
}
