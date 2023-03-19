package network

import (
	"io"
	"mmo_server/utils/mlog"
	"net"
)

type GConnection struct {
	session        *GSession
	conn           *net.TCPConn
	packageHandler *PackageHandler
}

func NewConnection(conn *net.TCPConn, session *GSession) *GConnection {
	return &GConnection{
		conn:    conn,
		session: session,
	}
}

// ReadMsg 读取连接上的信息
func (c *GConnection) ReadMsg() {
	defer func() {
		mlog.Info.Printf("Client[%s] is Disconnected!!!!")
		c.Close()
	}()
	c.packageHandler = NewPackageHandler(c)
	for {
		var buf []byte
		if _, err := io.ReadFull(c.conn, buf); err != nil {
			mlog.Warning.Printf("Connection[%s] Read message error : %v", c.conn.RemoteAddr(), err)
			break
		}
		if err := c.packageHandler.ReceiveMsg(buf); err != nil {
			mlog.Warning.Println("package Handler Receive message error : ", err)
			continue
		}
	}

}

// SendMsg 发送信息
func (c *GConnection) SendMsg() {

}

// Close 关闭连接
func (c *GConnection) Close() {
	if err := c.conn.Close(); err != nil {
		mlog.Error.Println("conn close err : ", err)
	}
}
