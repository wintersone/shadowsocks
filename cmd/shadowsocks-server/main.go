package main

import (
	"fmt"
	"log"
	"net"

	"github.com/wintersone/shadowsocks"

	"github.com/phayes/freeport"
	"github.com/wintersone/shadowsocks/cmd"
)

var version = "master"

func main() {
	log.SetFlags(log.Lshortfile)

	port, err := freeport.GetFreePort()

	if err != nil {
		port = 7448
	}

	config := &cmd.Config{
		ListenAddr: fmt.Sprintf("%d", port),
		Password:   shadowsocks.RnadPassword(),
	}
	config.ReadConfig()
	config.SaveConfig()

	// server start
	lsServer, err := shadowsocks.NewLsServer(config.Password, config.ListenAddr)
	if err != nil {
		log.Fatalln(err)
	}

	log.Fatalln(lsServer.Listen(func(listenAddr net.Addr) {
		log.Println("using config: ", fmt.Sprintf(`
			listening on %s password %s
		`, listenAddr, config.Password))
		log.Printf("shadowsocks-server: %s start on %s", version, listenAddr.String())
	}))
}
