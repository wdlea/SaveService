package service

import (
	"net/http"
)

func (s SaveService[GameState_T]) NewUser(resp http.ResponseWriter, req *http.Request) {
	var cookie http.Cookie
	cookie.Name = "user"

	user, err := s.MakeUser()
	if err != nil {
		resp.WriteHeader(500)
		return
	}

	encoded := s.EncryptUser(user)

	cookie.Value = string(encoded)

	http.SetCookie(resp, &cookie)

	resp.Write([]byte("balls in ur mouth"))
	resp.WriteHeader(0)
}
