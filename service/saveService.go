package service

import (
	"SaveSystem/set"
	"crypto/aes"
	"crypto/cipher"
	"net/http"
	"os"
)

type SaveService[GameState_T IGameState] struct {
	entries map[uint64]GameState_T
	users   set.Set[User]

	currentID uint64

	cipher cipher.Block
}

type IGameState interface{}

func (s SaveService[GameState_T]) Listen(addr string) {
	http.HandleFunc("/save", s.SaveUserData)
	http.HandleFunc("/new", s.NewUser)

	http.ListenAndServe(addr, nil)
}

func (s *SaveService[GameState_T]) Init() {
	key := os.Getenv("KEY")
	println("KEY: ", key)

	var err error
	s.cipher, err = aes.NewCipher([]byte(key))

	s.users = set.MakeSet[User](1024)

	if err != nil {
		panic(err)
	}
}

func (s *SaveService[GameState_T]) GetUserData(user User) (valid bool, data GameState_T) {
	valid = s.users.Has(user)
	if valid {
		data = s.entries[user.ID]
	}
	return
}

func (s *SaveService[GameState_T]) SetUserData(user User, data GameState_T) (valid bool) {
	valid = s.users.Has(user)
	if valid {
		s.entries[user.ID] = data
	}
	return
}
