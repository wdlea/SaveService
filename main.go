package main

import (
	"github.com/wdlea/SaveSystem/service"

	"github.com/joho/godotenv"
)

func main() {
	var s service.SaveService[GameState]
	godotenv.Load()
	s.Init()

	s.Listen("127.0.0.1:8080")
}

type GameState struct {
	Balls bool `json:"balls"`
}
