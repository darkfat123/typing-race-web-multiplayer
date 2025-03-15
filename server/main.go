package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"server/models"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var rooms = make(map[string]*models.Room)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return r.Header.Get("Origin") == os.Getenv("ALLOWED_ORIGIN")
	},
}

var texts = []string{
	"Despite what your teacher may have told you, there is a wrong way to wield a lasso. They got there early, and they got really good seats. The crowd yells and screams for more memes.",
	"Her life in the confines of the house became her new normal. The door had been painted and fixed. The fish twisted and turned to be free of the net. The memory we used to share is no longer coherent.",
	"The lyrics of the song sounded like fingernails on a chalkboard. The small white buoys marked the location of the divers. The beauty of the view is one of the reasons she moved to the countryside. The stranger officiates the meal.",
	"His ultimate dream fantasy consisted of being content and sleeping eight hours in a row. The sudden rainstorm washed crocodiles into the ocean. The sky is clear; the stars are twinkling. The waves were crashing on the shore; it was a lovely sight.",
	"The light in his life was never extinguished. The stranger officiates the meal. The lyrics of the song sounded like fingernails on a chalkboard. The small white buoys marked the location of the divers.",
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	var username, roomID string
	var data map[string]string
	if err := conn.ReadJSON(&data); err != nil {
		log.Println("Error reading username & roomID:", err)
		return
	}
	username = data["username"]
	roomID = data["roomID"]

	room := getOrCreateRoom(roomID)
	player := &models.Player{Conn: conn, Username: username, StartTime: time.Now(), Finished: false, Ready: false}

	room.Mutex.Lock()
	room.Players[conn] = player
	room.Mutex.Unlock()

	// Send text to player
	if err := conn.WriteJSON(map[string]string{"text": room.Text}); err != nil {
		log.Println("Error sending text:", err)
	}

	log.Printf("Player %s has joined room %s", player.Username, roomID)

	updateUserList(room)
	updateReadyStatus(room)

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
			updateReadyStatus(room)

			if isAllPlayersReady(room) {
				log.Printf("All players in room %s are ready. Starting the game!", roomID)
				broadcast(room, map[string]string{"type": "start_game"})
			}
		}

		if text, ok := message["text"]; ok {
			if strings.TrimSpace(text) == room.Text && !player.Finished {
				player.WordCount = len(strings.Fields(text))
				player.Finished = true
				wpm := calculateWPM(player)

				log.Printf("Player %s in room %s has finished typing with WPM: %.2f", player.Username, roomID, wpm)

				broadcast(room, map[string]interface{}{
					"type":     "finished",
					"username": player.Username,
					"wpm":      wpm,
				})
			}
		}
	}

	// Cleanup when player leaves
	room.Mutex.Lock()
	delete(room.Players, conn)
	room.Mutex.Unlock()
	updateUserList(room)

	// Remove empty rooms
	if len(room.Players) == 0 {
		log.Printf("Room %s is empty. Deleting room", roomID)
		delete(rooms, roomID)
	}
}

func isAllPlayersReady(room *models.Room) bool {
	room.Mutex.Lock()
	defer room.Mutex.Unlock()

	if len(room.Players) == 0 {
		return false
	}

	// Check if all players are ready
	for _, player := range room.Players {
		if !player.Ready {
			return false
		}
	}

	// All players are ready, lock the room
	room.Locked = true
	return true
}

func calculateWPM(player *models.Player) float64 {
	elapsedMinutes := time.Since(player.StartTime).Minutes()
	if elapsedMinutes == 0 {
		return 0
	}
	return float64(player.WordCount) / elapsedMinutes
}

func getOrCreateRoom(roomID string) *models.Room {
	if _, exists := rooms[roomID]; !exists {
		rooms[roomID] = &models.Room{
			Players: make(map[*websocket.Conn]*models.Player),
			Text:    texts[rand.Intn(len(texts))],
		}
	}
	return rooms[roomID]
}

func updateUserList(room *models.Room) {
	room.Mutex.Lock()
	var usernames []string
	for _, player := range room.Players {
		usernames = append(usernames, player.Username)
	}
	room.Mutex.Unlock()

	broadcast(room, map[string]interface{}{
		"type":  "update_users",
		"users": usernames,
	})
}

func updateReadyStatus(room *models.Room) {
	room.Mutex.Lock()
	var readyUsers []string
	for _, player := range room.Players {
		if player.Ready {
			readyUsers = append(readyUsers, player.Username)
		}
	}
	room.Mutex.Unlock()

	broadcast(room, map[string]interface{}{
		"type":  "update_ready",
		"users": readyUsers,
	})
}

func broadcast(room *models.Room, message interface{}) {
	room.Mutex.Lock()
	defer room.Mutex.Unlock()

	for conn := range room.Players {
		if err := conn.WriteJSON(message); err != nil {
			log.Println("Error sending broadcast message:", err)
		}
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	http.HandleFunc("/ws", handleWebSocket)
	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
