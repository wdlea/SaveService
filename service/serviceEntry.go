package service

import "sync"

// A service entry holds the state and a mutex to stop race conditions
type ServiceEntry[GameState_T IGameState] struct {
	state GameState_T
	mu    *sync.Mutex

	changed bool
}

// Creates a new service entry
func MakeServiceEntry[GameState_T IGameState](fresh bool) *ServiceEntry[GameState_T] {
	return &ServiceEntry[GameState_T]{
		mu:      &sync.Mutex{},
		changed: fresh,
	}
}
