package service

import (
	"net/http"

	"github.com/wdlea/SaveSystem/set"
)

// A save service is the functionality of this code
// it holds data in a GameState_T type which is the
// generic argument
// NOTE: SaveService.Init() must be called before using this
type SaveService[GameState_T IGameState] struct {
	entries map[uint64]GameState_T
	users   set.Set[User]

	currentID uint64
}

// The interface which all GameStates derive from
type IGameState interface{}

// Hosts the SaveService on a particular address
func (s SaveService[GameState_T]) Listen(addr string) {
	http.HandleFunc("/save", s.SaveUserData)
	http.HandleFunc("/new", s.newUser)

	http.ListenAndServe(addr, nil)
}

// Initializes the SaveService
// Note: this must be called
func (s *SaveService[GameState_T]) Init() {
	s.users = set.MakeSet[User](1024)
}

// Returns the users data, if
// 1. the user has valid ID and Key
// 2. the user has saved anything
// othewise valid will be false and data nil
func (s *SaveService[GameState_T]) GetUserData(user User) (valid bool, data GameState_T) {
	valid = s.users.Has(user)
	if valid {
		data = s.entries[user.ID]
	}
	return
}

// Sets the users data if the user has valid Key and ID
func (s *SaveService[GameState_T]) SetUserData(user User, data GameState_T) (valid bool) {
	valid = s.users.Has(user)
	if valid {
		s.entries[user.ID] = data
	}
	return
}
