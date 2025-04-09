package handler

import (
	"log"
	"net/http"
	"server/internal/logic"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

var lobbyClients = make(map[*websocket.Conn]bool)
var lobbyMutex = sync.Mutex{}

func broadcastRoomListToLobby() {
	lobbyMutex.Lock()
	defer lobbyMutex.Unlock()

	// กรองห้องที่มีผู้ใช้อยู่
	validRoomList := make(map[string][]string)
	log.Println("Test: ", logic.RoomIdList)
	for roomID, users := range logic.RoomIdList {
		if len(users) > 0 {
			validRoomList[roomID] = users
		} else {
			// ลบห้องที่ไม่มีผู้ใช้ทันที
			delete(logic.RoomIdList, roomID)
		}
	}

	// ตรวจสอบว่า validRoomList ว่างหรือไม่
	if len(validRoomList) == 0 {
		log.Println("No rooms found.")
	} else {
		// แสดงข้อมูลห้องที่มีผู้ใช้
		for roomID, users := range validRoomList {
			log.Printf("🏠 Room %s → [%s]", roomID, strings.Join(users, ", "))
		}
	}

	message := map[string]interface{}{
		"type":     "room_list",
		"roomList": validRoomList,
	}

	// broadcast ข้อมูลไปยังทุกคนใน lobby
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

	roomList := logic.RoomIdList
	conn.WriteJSON(map[string]interface{}{
		"type":     "room_list",
		"roomList": roomList,
	})

	// รออ่าน (บล็อคไว้ เพื่อไม่ให้ฟังก์ชันจบเร็ว)
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}

	// ลบ client ออกเมื่อหลุด
	lobbyMutex.Lock()
	delete(lobbyClients, conn)
	lobbyMutex.Unlock()
}

func JoinRoom(roomID string, username string) {
	logic.RoomIdList[roomID] = append(logic.RoomIdList[roomID], username)
	broadcastRoomListToLobby()
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
		// ถ้าไม่มีผู้ใช้เหลืออยู่ในห้อง → ลบห้องออกจาก map
		delete(logic.RoomIdList, roomID)
		log.Printf("Room %s deleted because it's empty", roomID)
	} else {
		// ถ้ามีผู้ใช้เหลืออยู่ในห้อง → อัปเดตรายชื่อ
		logic.RoomIdList[roomID] = newUsers
	}

	// อัปเดตข้อมูลไปยังผู้ใช้ทั้งหมดใน lobby
	broadcastRoomListToLobby()
}
