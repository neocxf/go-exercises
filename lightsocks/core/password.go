package core

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

const PasswordLength int = 256

var ErrInvalidPassword = errors.New("Illegal password")

type Password [PasswordLength]byte

func init() {
	log.Println("init func in lightsocks called")
	rand.Seed(time.Now().Unix())
}

func (password *Password) String() string {
	return base64.StdEncoding.EncodeToString(password[:])
}

func ParsePassword(passwordString string) (*Password, error) {
	bs, err := base64.StdEncoding.DecodeString(strings.TrimSpace(passwordString))
	if err != nil || len(bs) != PasswordLength {
		return nil, ErrInvalidPassword
	}

	password := Password{}
	copy(password[:], bs)
	bs = nil
	return &password, nil
}

func RandPassword() *Password {
	intArr := rand.Perm(PasswordLength)
	password := &Password{}

	for i, v := range intArr {
		password[i] = byte(v)
		if i == v {
			log.Println(fmt.Sprintf("something bad happens, %v with %v \n", i, v))
			return RandPassword()
		}
	}

	return password
}
