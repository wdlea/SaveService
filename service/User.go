package service

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"math/big"
)

type User struct {
	ID  uint64 `json:"id"`
	Key uint64 `json:"key"`
}

const MAX_KEY_SIZE = 0xFFFFFFFFFFFFFFFF

var maxKeySize *big.Int = big.NewInt(0).SetUint64(MAX_KEY_SIZE)


//makes a new user with the next available ID
func (s *SaveService[GameState_T]) MakeUser() (user User, err error) {
	key, err := rand.Int(rand.Reader, maxKeySize)

	s.currentID += 1
	user = User{
		ID:  s.currentID,
		Key: key.Uint64(),
	}
	return
}

func (u User) Hash(size uint64) uint64 {
	return u.ID % size
}

func (s SaveService[GameState_T]) UserToCookie(u User) (value string) {
	encrypted := s.encryptUser(u)
	value = base64.URLEncoding.EncodeToString(encrypted)
	return
}
func (s SaveService[GameState_T]) UserFromCookie(value string) (valid bool, u User) {
	encrypted, err := base64.URLEncoding.DecodeString(value)
	if err != nil {
		valid = false
		return
	}
	valid, u = s.decryptUser(encrypted)
	return
}

func (s SaveService[GameState_T]) encryptUser(u User) (ciphertext []byte) {
	return u.packUser()
}
func (s SaveService[GameState_T]) decryptUser(ciphertext []byte) (valid bool, u User) {
	return unpackUser(ciphertext)
}

func (u User) packUser() (packed []byte) {
	packed = binary.BigEndian.AppendUint64(packed, u.ID)
	packed = binary.BigEndian.AppendUint64(packed, u.Key)
	return
}
func unpackUser(packed []byte) (valid bool, u User) {
	if len(packed) != 16 {
		valid = false
		return
	}

	id_enc, key_enc := packed[:8], packed[8:]
	u.ID, u.Key = binary.BigEndian.Uint64(id_enc), binary.BigEndian.Uint64(key_enc)
	valid = true
	return
}
