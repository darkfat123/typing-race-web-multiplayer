package logic

import (
	"server/internal/model"
	"time"
)

func CalculateWPM(player *model.Player) float64 {
	elapsedMinutes := time.Since(player.StartTime).Minutes()
	if elapsedMinutes == 0 {
		return 0
	}
	return float64(player.WordCount) / elapsedMinutes
}
