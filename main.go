package main

import (
	"log"
	"net/http"
)

// type Stone string

// type Player struct {
// 	ID         int    `json:"id"`
// 	Name       string `json:"name"`
// 	Points     int    `json:"points"`
// 	Rubles     int    `json:"rubles"`
// 	Cards      []Card `json:"cards"`
// 	Hand       []Card `json:"hand"`
// 	StartStone Stone  `json:"startStone"`
// }

// type Card struct {
// 	ID          int    `json:"id"`
// 	Stone       Stone  `json:"stone"`
// 	Name        string `json:"name"`
// 	Value       int    `json:"value"`
// 	ResaleValue int    `json:"resaleValue"`
// 	PointReward int    `json:"pointReward"`
// 	RubleReward int    `json:"rubleReward"`
// }

// type Action struct {
// 	ID       int `json:"id"`
// 	PlayerID int `json:"playerId"`
// }

// type BuyAction struct {
// 	*Action
// 	CardID int `json:"cardId"`
// }

// type Game struct {
// 	ID      int               `json:"id"`
// 	OwnerID int               `json:"ownerId"`
// 	Players []Player          `json:"players"`
// 	Cards   map[string][]Card `json:"cards"`
// 	Phase   Stone             `json:"phase"`
// }

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	http.HandleFunc("/", helloWorldHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
