package service

import (
	"encoding/json"
	"io"
	"net/http"
)

func (s SaveService[GameState_T]) R_SaveUserData(resp http.ResponseWriter, req *http.Request) {
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

	body, err := io.ReadAll(req.Body)
	if err != nil {
		resp.WriteHeader(400)
		resp.Write([]byte("No request body"))
		return
	}

	var data GameState_T
	err = json.Unmarshal(body, &data)
	if err != nil {
		resp.WriteHeader(400)
		resp.Write(append([]byte("Bad request body"+err.Error()), body...))
		return
	}

	valid = s.SetUserData(user, data)
	if !valid {
		resp.WriteHeader(403)
		resp.Write([]byte("Invalid user cookie"))
		return
	}
	resp.WriteHeader(204)
}
