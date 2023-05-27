package service

import (
	"testing"
)

func TestUserEncyptDecrypt(t *testing.T) {
	s := SaveService[bool]{}
	s.Init()

	u := User{
		ID: 69,
		Key: 123456789,
	}
	enc := s.encryptUser(u)
	valid, dec := s.decryptUser(enc)

	if !valid{
		t.Fatalf("User is not valid when should be valid")
	}
	if dec != u{
		t.Fatalf("User and decrypted do not match")
	}
}