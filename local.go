package shadowsocks

import (
	"log"
	"net"
)

type LsLocal struct {
	Cipher     *cipher
	ListenAddr *net.TCPAddr
	RemoteAddr *net.TCPAddr
}

func NewLsLocal(password string, listenAddr, remoteAddr string) (*LsLocal, error) {
	bsPassword, err := parsePassword(password)
	if err != nil {
		return nil, err
	}

	structListenAddr, err := net.ResolveTCPAddr("TCP", listenAddr)
	if err != nil {
		return nil, err
	}

	structRemoteAddr, err := net.ResolveTCPAddr("tcp", remoteAddr)
	if err != nil {
		return nil, err
	}

	return &LsLocal{
		Cipher:     newCipher(bsPassword),
		ListenAddr: structListenAddr,
		RemoteAddr: structRemoteAddr,
	}, nil
}

func (local *LsLocal) Listen(didListen func(listenAddr net.Addr)) error {
	return ListenSecureTCP(local.ListenAddr, local.Cipher, local.handleConn, didListen)
}

func (local *LsLocal) handleConn(userConn *SecureTCPConn) {
	defer userConn.Close()

	proxyServer, err := DialTCPSecure(local.RemoteAddr, local.Cipher)
	if err != nil {
		log.Println(err)
		return
	}

	defer proxyServer.Close()

	go func() {
		err := proxyServer.DecodeCopy(userConn)
		if err != nil {
			userConn.Close()
			proxyServer.Close()
		}
	}()

	userConn.EncodeCopy(proxyServer)
}
