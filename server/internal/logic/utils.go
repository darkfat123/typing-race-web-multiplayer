package logic

import (
	"log"
	"math/rand"
	"server/internal/model"
	"server/pkg/texts"
	"strconv"
	"time"
)

func CalculateWPM(player *model.Player) float64 {
	elapsedSeconds := time.Since(player.StartTime).Seconds() - 3
	elapsedMinutes := elapsedSeconds / 60

	if elapsedMinutes == 0 {
		return 0
	}
	log.Printf("Elapsed time for %s in minutes: %f", player.Username, elapsedMinutes)
	log.Printf("Word count for %s: %d", player.Username, player.WordCount)

	return float64(player.WordCount) / elapsedMinutes
}

func RandomRoomId() string {
	id := strconv.Itoa(rand.Intn(10000))
	return id
}

// getRandomText selects a random text based on the provided language.
func GetRandomText(language string) string {
	if language == "th" {
		return texts.ThaiTexts[rand.Intn(len(texts.ThaiTexts))]
	}
	return texts.EngTexts[rand.Intn(len(texts.EngTexts))]
}
