package logic

import (
	"math/rand"
	"server/internal/model"
	"strconv"
	"time"
)

func CalculateWPM(player *model.Player) float64 {
	elapsedMinutes := time.Since(player.StartTime).Minutes()
	if elapsedMinutes == 0 {
		return 0
	}
	return float64(player.WordCount) / elapsedMinutes
}

func RandomRoomId() string {
	id := strconv.Itoa(rand.Intn(10000))
	return id
}
