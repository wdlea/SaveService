package service

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"math/big"
)

type User struct {
	ID  uint64 `json:"id"`
	Key uint64 `json:"key"`
}

const MAX_KEY_SIZE = 0xFFFFFFFFFFFFFFFF

var maxKeySize *big.Int = big.NewInt(0).SetUint64(MAX_KEY_SIZE)

func (s *SaveService[GameState_T]) MakeUser() (user User, err error) {
	key, err := rand.Int(rand.Reader, maxKeySize)

	s.currentID += 1
	user = User{
		ID:  s.currentID,
		Key: key.Uint64(),
	}
	return
}
func (s *SaveService[GameState_T]) EncryptUser(u User) (encrypted []byte) {
	encrypted = u.EncodeUser()

	s.cipher.Encrypt(encrypted, encrypted)
	return
}
func (s *SaveService[GameState_T]) DecryptUser(encrypted []byte) (valid bool, user User) {
	decrypted := make([]byte, len(encrypted)+1)

	defer func() {
		if err := recover(); err != nil {
			println(err)
		}
	}()

	s.cipher.Decrypt(encrypted, decrypted)

	user, err := DecodeUser(decrypted)
	valid = err == nil

	return
}

func (u User) EncodeUser() (encoded []byte) {
	encoded = binary.LittleEndian.AppendUint64(encoded, u.ID)
	encoded = binary.LittleEndian.AppendUint64(encoded, u.Key)
	println("Encoded:", encoded)
	return
}
func DecodeUser(encoded []byte) (u User, err error) {
	if len(encoded) != 16 {
		err = errors.New("encoded user not 16 bytes in length")
	}

	u = User{
		ID:  binary.LittleEndian.Uint64(encoded[0:8]),
		Key: binary.LittleEndian.Uint64(encoded[8:16]),
	}

	return
}

func (u User) Hash(size uint64) uint64 {
	return u.ID % size
}
