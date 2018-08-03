package main

import (
	"fmt"
	"github.com/neocxf/go-exercises/lightsocks/cmd"
	"github.com/neocxf/go-exercises/lightsocks/core"
	"github.com/neocxf/go-exercises/lightsocks/server"
	"github.com/phayes/freeport"
	"log"
	"net"
)

var version = "master"

func main() {
	log.SetFlags(log.Lshortfile)

	port, err := freeport.GetFreePort()
	if err != nil {
		port = 7448
	}

	config := &cmd.Config{
		ListenAddr: fmt.Sprintf(":%d", port),
		Password:   core.RandPassword().String(),
	}

	config.ReadConfig()
	config.SaveConfig()

	password, err := core.ParsePassword(config.Password)
	if err != nil {
		log.Fatalln(err)
	}

	listenAddr, err := net.ResolveTCPAddr("tcp", config.ListenAddr)
	if err != nil {
		log.Fatalln(err)
	}

	lsServer := server.New(password, listenAddr)
	log.Fatalln(lsServer.Listen(func(listenAddr net.Addr) {
		log.Println("using config file:", fmt.Sprintf(`local listen: %s, password: %s`, listenAddr, password))
		log.Printf("lightsocks-server: %s launch success at %s\n", version, listenAddr.String())

	}))

}
