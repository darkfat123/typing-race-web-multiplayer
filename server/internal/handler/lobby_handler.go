package handler

import (
	"log"
	"net/http"
	"server/internal/logic"
	"server/internal/model"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

var lobbyClients = make(map[*websocket.Conn]bool)
var lobbyMutex = sync.Mutex{}

func BroadcastRoomListToLobby() {
	lobbyMutex.Lock()
	defer lobbyMutex.Unlock()

	validRoomList := FilterUnlockedRooms()

	if len(validRoomList) == 0 {
		log.Println("No rooms found.")
	} else {
		for roomID, data := range validRoomList {
			users := data["users"].([]string)
			language := data["language"].(string)
			limit := data["limit"]
			log.Printf("🏠 Room %s (max: %d) → [%s] (%s) ", roomID, limit, strings.Join(users, ", "), language)
		}

	}

	message := map[string]interface{}{
		"type":     "room_list",
		"roomList": validRoomList,
	}

	for conn := range lobbyClients {
		if err := conn.WriteJSON(message); err != nil {
			log.Println("Error broadcasting to client:", err)
			conn.Close()
			delete(lobbyClients, conn)
		}
	}
}

func HandleLobbyWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	lobbyMutex.Lock()
	lobbyClients[conn] = true
	lobbyMutex.Unlock()

	BroadcastRoomListToLobby()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}

	lobbyMutex.Lock()
	delete(lobbyClients, conn)
	lobbyMutex.Unlock()
}

func FilterUnlockedRooms() map[string]map[string]interface{} {
	result := make(map[string]map[string]interface{})

	for roomID, room := range logic.Rooms {
		if !room.Locked {
			if users, exists := logic.RoomIdList[roomID]; exists && len(users) > 0 {
				result[roomID] = map[string]interface{}{
					"users":    users,
					"language": room.Language,
					"limit":    room.Limit,
				}
			}
		}
	}

	return result
}

func JoinRoom(roomID string, username string) {
	logic.RoomIdList[roomID] = append(logic.RoomIdList[roomID], username)
	BroadcastRoomListToLobby()
}

func IsRoomLocked(roomList map[string]*model.Room) {
	for id, room := range roomList {
		if room.Locked {
			delete(logic.RoomIdList, id)
		}
	}
	BroadcastRoomListToLobby()
}

func RemoveUserFromRoom(roomID string, username string) {
	users := logic.RoomIdList[roomID]
	newUsers := []string{}

	// กรองชื่อที่ไม่ใช่ username นี้
	for _, user := range users {
		if user != username {
			newUsers = append(newUsers, user)
		}
	}

	if len(newUsers) == 0 {
		delete(logic.RoomIdList, roomID)
		log.Printf("Room %s deleted because it's empty", roomID)
	} else {
		logic.RoomIdList[roomID] = newUsers
	}

	BroadcastRoomListToLobby()
}
