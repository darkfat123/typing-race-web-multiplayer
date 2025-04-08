package logic

import (
	"log"
	"math/rand"
	"server/internal/model"
	"server/pkg/texts"
	"time"

	"github.com/gorilla/websocket"
)

var (
	rooms      = make(map[string]*model.Room)
	RoomIdList = make(map[string][]string)
)

func init() {
	// RoomIdList = map[string][]string{
	// 	"room1": {"user1", "user2", "user3", "user4", "user5", "user6", "user7", "user8", "user9"},
	// 	"room2": {"user4", "user5"},
	// 	"room3": {"user6", "user7", "user8", "user9"},
	// }
}

// GetOrCreateRoom retrieves an existing room or creates a new one if it doesn't exist.
func GetOrCreateRoom(roomIDInput string, language string) *model.Room {
	// ถ้า roomIDInput มีอยู่แล้ว ก็ return ห้องเดิม
	if room, exists := rooms[roomIDInput]; exists {
		return room
	}

	selectedText := getRandomText(language)
	randomID := RandomRoomId() // สุ่ม ID ก่อน

	room := &model.Room{
		ID:       randomID, // ใช้ ID ที่สุ่มมา
		Language: language,
		Players:  make(map[*websocket.Conn]*model.Player),
		Text:     selectedText,
	}
	rooms[randomID] = room // เก็บใน map โดยใช้ random ID

	return room
}

// getRandomText selects a random text based on the provided language.
func getRandomText(language string) string {
	if language == "th" {
		return texts.ThaiTexts[rand.Intn(len(texts.ThaiTexts))]
	}
	return texts.EngTexts[rand.Intn(len(texts.EngTexts))]
}

// UpdateUserList updates the user list of the room and broadcasts the new list.
func UpdateUserList(room *model.Room) {
	room.Mutex.Lock()
	usernames := getUsernamesFromRoom(room)
	room.Mutex.Unlock()

	// Update user list in roomIdList
	RoomIdList[room.ID] = usernames

	// Broadcast updated user list
	Broadcast(room, map[string]interface{}{
		"type":  "update_users",
		"users": usernames,
	})
}

// getUsernamesFromRoom returns a list of usernames from the room.
func getUsernamesFromRoom(room *model.Room) []string {
	var usernames []string
	for _, player := range room.Players {
		usernames = append(usernames, player.Username)
	}
	return usernames
}

// UpdateReadyStatus updates the ready status of players and broadcasts the ready users.
func UpdateReadyStatus(room *model.Room) {
	room.Mutex.Lock()
	readyUsers := getReadyUsersFromRoom(room)
	room.Mutex.Unlock()

	Broadcast(room, map[string]interface{}{
		"type":  "update_ready",
		"users": readyUsers,
	})
}

// getReadyUsersFromRoom returns a list of players who are ready in the room.
func getReadyUsersFromRoom(room *model.Room) []string {
	var readyUsers []string
	for _, player := range room.Players {
		if player.Ready {
			readyUsers = append(readyUsers, player.Username)
		}
	}
	return readyUsers
}

// Broadcast sends a message to all players in the room.
func Broadcast(room *model.Room, message interface{}) {
	room.Mutex.Lock()
	defer room.Mutex.Unlock()

	for conn := range room.Players {
		if err := conn.WriteJSON(message); err != nil {
			log.Println("Error sending broadcast message:", err)
		}
	}
}

// IsAllPlayersReady checks if all players are ready and locks the room if they are.
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

// CleanupPlayer removes a player from the room and deletes the room if empty.
func CleanupPlayer(room *model.Room, conn *websocket.Conn, roomID string) {
	room.Mutex.Lock()
	delete(room.Players, conn)
	room.Mutex.Unlock()

	// Remove empty rooms
	if len(room.Players) == 0 {
		log.Printf("Room %s is empty. Deleting room", roomID)
		delete(rooms, roomID)
	}

	UpdateUserList(room)
}
