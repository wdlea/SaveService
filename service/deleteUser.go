package service

import "net/http"

func (s *SaveService[GameState_T]) r_DeleteUser(resp http.ResponseWriter, req *http.Request) {
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

	if s.users.Pop(user) {
		delete(s.entries, user.ID)
	}

	resp.Write([]byte("Sucessfully deleted user."))
}
