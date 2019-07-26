package main

import (
	"fmt"
	"log"
	"net"

	"github.com/wintersone/shadowsocks"

	"github.com/wintersone/shadowsocks/cmd"
)

const (
	DefaultListenAddr = ":7448"
)

var version = "master"

func main() {
	log.SetFlags(log.Lshortfile)

	config := &cmd.Config{
		ListenAddr: DefaultListenAddr,
	}
	config.ReadConfig()
	config.SaveConfig()

	lsLocal, err := shadowsocks.NewLsLocal(config.Password, config.ListenAddr, config.RemoteAddr)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatalln(lsLocal.Listen(func(listenAddr net.Addr) {
		log.Println("using config: ", fmt.Sprintf(`
			local listening port: %s
			remote server address: %s
			password: %s
			`, listenAddr, config.RemoteAddr, config.Password))
		log.Printf("shadowsocks-local:%s start, listening on port %s", version, listenAddr.String())
	}))

}
