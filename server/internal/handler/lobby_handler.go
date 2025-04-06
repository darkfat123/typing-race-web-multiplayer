package handler

import (
	"log"
	"net/http"
	"server/internal/logic"
)

func HandleLobbyWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	// ส่ง room list ให้ผู้ใช้ทันที
	roomList := logic.GetRoomIdList()
	conn.WriteJSON(map[string]interface{}{
		"type":     "room_list",
		"roomList": roomList,
	})
}
