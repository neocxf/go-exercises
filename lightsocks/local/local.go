package local

import (
	"github.com/neocxf/go-exercises/lightsocks/core"
	"log"
	"net"
)

type LsLocal struct {
	*core.SecureSocket
}

func New(password *core.Password, listenAddr, remoteAddr *net.TCPAddr) *LsLocal {
	return &LsLocal{
		SecureSocket: &core.SecureSocket{
			Cipher:     core.NewCipher(password),
			ListenAddr: listenAddr,
			RemoteAddr: remoteAddr,
		},
	}
}

func (local *LsLocal) Listen(didListen func(listenAddr net.Addr)) error {
	listener, err := net.ListenTCP("tcp", local.ListenAddr)

	if err != nil {
		return err
	}

	defer listener.Close()

	if didListen != nil {
		didListen(listener.Addr())
	}

	for {
		userConn, err := listener.AcceptTCP()

		if err != nil {
			log.Println(err)
			continue
		}

		// when userConn closed, just remove all the data no matter whether we have the un-send data
		userConn.SetLinger(0)

		go local.handleConn(userConn)
	}

	return nil
}

func (local *LsLocal) handleConn(userConn *net.TCPConn) {
	defer userConn.Close()

	proxyServer, err := local.DialRemote()

	if err != nil {
		log.Println(err)
		return
	}

	defer proxyServer.Close()

	proxyServer.SetLinger(0)

	go func() {
		err := local.DecodeCopy(userConn, proxyServer)

		if err != nil { // while coping, the network may timeout in such case, just quit
			userConn.Close()
			proxyServer.Close()
		}

	}()

	local.EncodeCopy(proxyServer, userConn)
}
