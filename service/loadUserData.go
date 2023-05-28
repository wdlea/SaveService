package service

import (
	"encoding/json"
	"net/http"
)

func (s SaveService[GameState_T]) r_LoadUserData(resp http.ResponseWriter, req *http.Request) {
	userCookie, err := req.Cookie("user")
	if err != nil || userCookie == nil {
		resp.WriteHeader(403)
		return
	}
	valid, user := s.UserFromCookie(userCookie.Value)
	if !valid {
		resp.WriteHeader(403)
		return
	}

	valid, data := s.GetUserData(user)
	if !valid {
		resp.WriteHeader(403)
		return
	}

	encoded, err := json.Marshal(data)
	if err != nil {
		resp.WriteHeader(500)
		return
	}

	resp.Write(encoded)
}
