package network

import (
	"mmo_server/utils/mlog"
	"net"
)

type GNetTCPListener struct {
	listener *net.TCPListener
	addr     *net.TCPAddr
}

func NewListener(network string, address string) *GNetTCPListener {
	a, err := net.ResolveTCPAddr(network, address)
	if err != nil {
		mlog.Error.Fatalln("resolve tcp address is err : ", err)
		return nil
	}

	l, err2 := net.ListenTCP(network, a)
	if err2 != nil {
		mlog.Error.Fatalln("create tcp listener is err : ", err2)
		return nil
	}

	return &GNetTCPListener{
		listener: l,
		addr:     a,
	}
}

func (ntl *GNetTCPListener) AcceptConn() *net.TCPConn {
	conn, err := ntl.listener.AcceptTCP()
	if err != nil {
		mlog.Error.Println("Accept connect err : ", err)
		return nil
	}
	mlog.Info.Printf("New Client[%s] Connection ....\n", conn.RemoteAddr())
	return conn
}

func (ntl *GNetTCPListener) Close() {
	if err := ntl.listener.Close(); err != nil {
		mlog.Error.Println("TCP Listener Close is error : ", err)
		return
	}
}
