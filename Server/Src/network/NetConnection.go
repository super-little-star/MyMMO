package network

import "net"

type GConnection struct {
	session *GSession
	conn    *net.Conn
}
