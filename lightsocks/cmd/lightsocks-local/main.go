package main

import (
	"fmt"
	"github.com/neocxf/go-exercises/lightsocks/cmd"
	"github.com/neocxf/go-exercises/lightsocks/core"
	"github.com/neocxf/go-exercises/lightsocks/local"
	"log"
	"net"
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

	password, err := core.ParsePassword(config.Password)
	if err != nil {
		log.Fatalln(err)
	}

	listenAddr, err := net.ResolveTCPAddr("tcp", config.ListenAddr)
	if err != nil {
		log.Fatalln(err)
	}

	remoteAddr, err := net.ResolveTCPAddr("tcp", config.RemoteAddr)
	if err != nil {
		log.Fatalln(err)
	}

	lsLocal := local.New(password, listenAddr, remoteAddr)
	log.Fatalln(lsLocal.Listen(func(listenAddr net.Addr) {
		log.Println("using config file: ", fmt.Sprintf(`local listen: %s, remote listen:%s, password:%s`, listenAddr, remoteAddr, password))
		log.Printf("lightsocks-local: %s launch success at %s\n", version, listenAddr.String())
	}))

}
