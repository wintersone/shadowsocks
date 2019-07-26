package shadowsocks

import (
	"encoding/base64"
	"errors"
	"math/rand"
	"strings"
	"time"
)

const passwordLength = 256

type password [passwordLength]byte

func init() {
	rand.Seed(time.Now().Unix())
}

func (password *password) String() string {
	return base64.StdEncoding.EncodeToString(password[:])
}

func parsePassword(passwordString string) (*password, error) {
	bs, err := base64.StdEncoding.DecodeString(strings.TrimSpace(passwordString))
	if err != nil || len(bs) != passwordLength {
		return nil, errors.New("Illigal Password")
	}
	password := password{}
	copy(password[:], bs)
	bs = nil
	return &password, nil
}

func RnadPassword() string {
	intArr := rand.Perm(passwordLength)
	password := &password{}

	for i, v := range intArr {
		password[i] = byte(v)
		if i == v {
			return RnadPassword()
		}
	}
	return password.String()
}
