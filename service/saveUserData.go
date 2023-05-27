package service

import (
	"encoding/json"
	"io"
	"net/http"
)

func (s SaveService[GameState_T]) SaveUserData(resp http.ResponseWriter, req *http.Request) {
	userCookie, err := req.Cookie("user")
	if err != nil || userCookie == nil {
		resp.WriteHeader(403)
		return
	}
	valid, user := s.DecryptUser([]byte(userCookie.Value))
	if !valid {
		resp.WriteHeader(403)
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		resp.WriteHeader(400)
		return
	}

	var data GameState_T
	err = json.Unmarshal(body, &data)
	if err != nil {
		resp.WriteHeader(400)
		return
	}

	valid = s.SetUserData(user, data)
	if !valid {
		resp.WriteHeader(403)
		return
	}
	resp.WriteHeader(204)
}
