package network

import (
	"mmo_server/utils/mlog"
	"net"
)

type GNetListener struct {
	listener *net.TCPListener
	addr     *net.TCPAddr
}

func NewListener(network string, address string) *GNetListener {
	a, err := net.ResolveTCPAddr(network, address)
	if err != nil {
		mlog.Error.Fatalln("resolve tcp address is err : ", err)
		return nil
	}

	l, err := net.ListenTCP(network, a)
	if err != nil {
		mlog.Error.Fatalln("create tcp listener is err : ", err)
		return nil
	}

	return &GNetListener{
		listener: l,
		addr:     a,
	}
}

func (ntl *GNetListener) AcceptConn() *net.TCPConn {
	conn, err := ntl.listener.AcceptTCP()
	if err != nil {
		mlog.Error.Println("Accept connect err : ", err)
		return nil
	}
	mlog.Info.Printf("New Client[%s] Connection ....\n", conn.RemoteAddr())
	return conn
}

func (ntl *GNetListener) Close() {
	ntl.Close()
}
