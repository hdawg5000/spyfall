package main

import (
	"fmt"
	"math/rand"
	"time"

	"./player"
)

var gameIDs [10]int

func main() {
	rand.Seed(time.Now().UnixNano())
	id := generateGameID()
	gameIDs[0] = id
	fmt.Printf("game id is %v", gameIDs[0])

	player := player.New(1234, 4567)
	fmt.Printf("\nplayer id %v", player.GetID())
}

func generateGameID() int {
	return rand.Intn(99999)
}
