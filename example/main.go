package main

import (
	"github.com/wdlea/SaveSystem/service"
)

func main() {
	var s service.SaveService[GameState]
	s.Init()

	s.Listen("127.0.0.1:8080")
}

type GameState struct {
	SomeBool bool `json:"some_bool"`
	// SomeInt  int                `json:"some_int"`
	// AStruct  struct{ a uint64 } `json:"a_struct"`
}
