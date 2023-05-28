package service

import (
	"net/http"
)

func (s SaveService[GameState_T]) r_NewUser(resp http.ResponseWriter, req *http.Request) {
	var cookie http.Cookie
	cookie.Name = "user"

	user, err := s.MakeUser()
	if err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte("Eror in account creation"))
		return
	}

	cookie.Value = s.UserToCookie(user)

	http.SetCookie(resp, &cookie)

	resp.Write([]byte("balls in ur mouth"))
}
