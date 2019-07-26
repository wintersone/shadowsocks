package shadowsocks

type cipher struct {
	encodePassword *password
	decodePassword *password
}

func (cipher *cipher) encode(bs []byte) {
	for i, v := range bs {
		bs[i] = cipher.encodePassword[v]
	}
}

func (cipher *cipher) decode(bs []byte) {
	for i, v := range bs {
		bs[i] = cipher.decodePassword[v]
	}
}

func newCipher(encodePassword *password) *cipher {
	decodePassword := &password{}

	for i, v := range encodePassword {
		encodePassword[i] = v
		decodePassword[v] = byte(i)
	}

	return &cipher{
		encodePassword: encodePassword,
		decodePassword: decodePassword,
	}
}
