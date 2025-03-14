package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

type Player struct {
	Conn      *websocket.Conn
	Username  string
	StartTime time.Time
	WordCount int
	Finished  bool
}

type Room struct {
	Players map[*websocket.Conn]*Player
	Text    string
	Mutex   sync.Mutex
}

var rooms = make(map[string]*Room)
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
	player := &Player{Conn: conn, Username: username, StartTime: time.Now(), Finished: false}

	room.Mutex.Lock()
	room.Players[conn] = player
	room.Mutex.Unlock()

	if err := conn.WriteJSON(map[string]string{"text": room.Text}); err != nil {
		log.Println("Error sending text:", err)
	}

	for {
		var message map[string]string
		if err := conn.ReadJSON(&message); err != nil {
			log.Println("Error reading message:", err)
			break
		}

		if text, ok := message["text"]; ok {
			if strings.TrimSpace(text) == room.Text && !player.Finished {
				player.WordCount = len(strings.Fields(text))
				player.Finished = true
				wpm := calculateWPM(player)

				room.Mutex.Lock()
				for c := range room.Players {
					if err := c.WriteJSON(map[string]interface{}{
						"username": player.Username,
						"wpm":      wpm,
						"status":   "finished",
					}); err != nil {
						log.Println("Error sending WPM:", err)
					}
				}
				room.Mutex.Unlock()

				log.Printf("Player %s has finished typing. Sending 'finished' status.", player.Username)
			}
		}
	}

	room.Mutex.Lock()
	delete(room.Players, conn)
	room.Mutex.Unlock()
}

func calculateWPM(player *Player) float64 {
	elapsedMinutes := time.Since(player.StartTime).Minutes()
	if elapsedMinutes == 0 {
		return 0
	}
	return float64(player.WordCount) / elapsedMinutes
}

func getOrCreateRoom(roomID string) *Room {
	if _, exists := rooms[roomID]; !exists {
		rooms[roomID] = &Room{
			Players: make(map[*websocket.Conn]*Player),
			Text:    texts[rand.Intn(len(texts))],
		}
	}
	return rooms[roomID]
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
