package core

import (
	"errors"
	"fmt"
	"io"
	"net"
)

const (
	BufSize = 1024
)

type SecureSocket struct {
	Cipher     *Cipher
	ListenAddr *net.TCPAddr
	RemoteAddr *net.TCPAddr
}

func (secureSocket *SecureSocket) DecodeRead(conn *net.TCPConn, bs []byte) (n int, err error) {
	n, err = conn.Read(bs)

	if err != nil {
		return
	}

	secureSocket.Cipher.decode(bs[:n])

	return
}

func (secureSocket *SecureSocket) EncodeWrite(conn *net.TCPConn, bs []byte) (n int, err error) {
	secureSocket.Cipher.encode(bs)

	return conn.Write(bs)
}

func (secureSocket *SecureSocket) EncodeCopy(dst *net.TCPConn, src *net.TCPConn) error {
	buf := make([]byte, BufSize)

	for {
		readCount, errRead := src.Read(buf)
		if errRead != nil {
			if errRead != io.EOF {
				return errRead
			} else {
				return nil
			}
		}

		if readCount > 0 {
			writeCount, errWrite := secureSocket.EncodeWrite(dst, buf[0:readCount])

			if errWrite != nil {
				return errWrite
			}

			if readCount != writeCount {
				return io.ErrShortWrite
			}
		}
	}
}

func (secureSocket *SecureSocket) DecodeCopy(dst *net.TCPConn, src *net.TCPConn) error {
	buf := make([]byte, BufSize)

	for {
		readCount, errRead := secureSocket.DecodeRead(src, buf)
		if errRead != nil {
			if errRead != io.EOF {
				return errRead
			} else {
				return nil
			}
		}

		if readCount > 0 {
			writeCount, errWrite := dst.Write(buf[0:readCount])

			if errWrite != nil {
				return errWrite
			}

			if readCount != writeCount {
				return io.ErrShortWrite
			}
		}
	}
}

func (secureSocket *SecureSocket) DialRemote() (*net.TCPConn, error) {
	remoteConn, err := net.DialTCP("tcp", nil, secureSocket.RemoteAddr)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("connect to remote server %s fail for %s", secureSocket.RemoteAddr, err))
	}

	return remoteConn, nil
}
