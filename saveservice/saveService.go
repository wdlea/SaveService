package saveservice

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"net/http"
	"os"
)

type SaveService[GameState_T IGameState, User_T IUser] struct {
	entries map[User_T]GameState_T

	cipher cipher.Block
}

type IUser interface {
	comparable

	Authenticate(ID string, key string) bool
}

func (s SaveService[GameState_T, User_T]) Listen(addr string, port string) {
	http.HandleFunc("save", s.SaveUserData)
	http.HandleFunc("login", s.AuthenticateUser)
	http.HandleFunc("new", s.NewUser)
}

func (s SaveService[GameState_T, User_T]) SaveUserData(resp http.ResponseWriter, req *http.Request) {

}
func (s SaveService[GameState_T, User_T]) AuthenticateUser(resp http.ResponseWriter, req *http.Request) {
	cookie, cookie_err := req.Cookie("user")
	if cookie_err != nil {
		s.NewUser(resp, req)
		return
	}
	decrypted := make([]byte, 1024)
	s.cipher.Decrypt(decrypted, []byte(cookie.Value))

	fmt.Println(string(decrypted))
}
func (s SaveService[GameState_T, User_T]) NewUser(resp http.ResponseWriter, req *http.Request) {

}

func (s SaveService[GameState_T, User_T]) InitCyphers() {
	key := os.Getenv("KEY")

	var err error
	s.cipher, err = aes.NewCipher([]byte(key))

	if err != nil {
		panic(err)
	}
}
