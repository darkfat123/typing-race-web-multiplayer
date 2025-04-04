package logic

import (
	"log"
	"math/rand"
	"server/internal/model"
	"server/pkg/texts"
	"time"

	"github.com/gorilla/websocket"
)

var rooms = make(map[string]*model.Room)

func GetOrCreateRoom(roomID string, language string) *model.Room {
	if _, exists := rooms[roomID]; !exists {
		var selectedText string
		if language == "th" {
			selectedText = texts.ThaiTexts[rand.Intn(len(texts.ThaiTexts))]
		} else {
			selectedText = texts.EngTexts[rand.Intn(len(texts.EngTexts))]
		}

		rooms[roomID] = &model.Room{
			Players: make(map[*websocket.Conn]*model.Player),
			Text:    selectedText,
		}
	}
	return rooms[roomID]
}

func UpdateUserList(room *model.Room) {
	room.Mutex.Lock()
	var usernames []string
	for _, player := range room.Players {
		usernames = append(usernames, player.Username)
	}
	room.Mutex.Unlock()

	Broadcast(room, map[string]interface{}{
		"type":  "update_users",
		"users": usernames,
	})
}

func UpdateReadyStatus(room *model.Room) {
	room.Mutex.Lock()
	var readyUsers []string
	for _, player := range room.Players {
		if player.Ready {
			readyUsers = append(readyUsers, player.Username)
		}
	}
	room.Mutex.Unlock()

	Broadcast(room, map[string]interface{}{
		"type":  "update_ready",
		"users": readyUsers,
	})
}

func Broadcast(room *model.Room, message interface{}) {
	room.Mutex.Lock()
	defer room.Mutex.Unlock()

	for conn := range room.Players {
		if err := conn.WriteJSON(message); err != nil {
			log.Println("Error sending broadcast message:", err)
		}
	}
}

func IsAllPlayersReady(room *model.Room) bool {
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

	// All players are ready
	room.Locked = true
	now := time.Now()

	// Set StartTime for each player
	for _, player := range room.Players {
		player.StartTime = now
	}

	return true
}

func CleanupPlayer(room *model.Room, conn *websocket.Conn, roomID string) {
	room.Mutex.Lock()
	delete(room.Players, conn)
	room.Mutex.Unlock()
	UpdateUserList(room)

	// Remove empty rooms
	if len(room.Players) == 0 {
		log.Printf("Room %s is empty. Deleting room", roomID)
		delete(rooms, roomID)
	}
}
