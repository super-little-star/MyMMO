package Network

import (
	"mmo_server/utils/mlog"
	"net"
)

type GNetTCPListener struct {
	listener *net.TCPListener
	addr     *net.TCPAddr
}

func NewTCPListener(address string) *GNetTCPListener {
	a, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		mlog.Error.Fatalln("resolve tcp address is err : ", err)
		return nil
	}

	l, err := net.ListenTCP("tcp", a)
	if err != nil {
		mlog.Error.Fatalln("create tcp listener is err : ", err)
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
	mlog.Info.Printf("New Client[%s] Connection ", conn.RemoteAddr())
	return conn
}

func (ntl *GNetTCPListener) Close() {
	ntl.Close()
}
