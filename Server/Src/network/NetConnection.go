package network

import (
	"io"
	ProtoMessage "mmo_server/protocol"
	"mmo_server/utils/mlog"
	"net"
)

type GConnection struct {
	session        *GSession
	conn           *net.TCPConn
	packageHandler *PackageHandler
	chanIsClose    chan bool
	chanSendData   chan []byte
	isClose        bool
}

func NewConnection(conn *net.TCPConn, session *GSession) *GConnection {
	return &GConnection{
		conn:         conn,
		session:      session,
		chanIsClose:  make(chan bool, 1),
		chanSendData: make(chan []byte),
		isClose:      false,
	}
}

// ReadMsg 读取连接上的信息
func (c *GConnection) readMsg() {
	mlog.Info.Printf("Client[%s] Read Goroutine is Running ....\n", c.conn.RemoteAddr())
	defer func() {
		mlog.Info.Printf("Client[%s] is Disconnected!!!!\n", c.conn.RemoteAddr())
		c.close()
	}()
	c.packageHandler = NewPackageHandler(c)
	for {
		var buf []byte
		if _, err := io.ReadFull(c.conn, buf); err != nil {
			mlog.Warning.Printf("Connection[%s] Read message error : %v\n", c.conn.RemoteAddr(), err)
			break
		}
		if err := c.packageHandler.ReceiveMsg(buf); err != nil {
			mlog.Warning.Println("package Handler Receive message error : ", err)
			continue
		}
	}
}

// WriteMsg 将chanSendData里的数据发送给客户端
func (c *GConnection) writeMsg() {
	mlog.Info.Printf("Client[%s] Write Goroutine is Running ....\n", c.conn.RemoteAddr())
	defer mlog.Info.Printf("Client[%s] Write Goroutine is Close !!!!\n", c.conn.RemoteAddr())

	for {
		select {
		case data := <-c.chanSendData: // 把需要发送的信息取出来发送出去
			if _, err := c.conn.Write(data); err != nil {
				mlog.Error.Printf("Client[%s] connection write data is error : %v\n", c.conn.RemoteAddr(), err)
				continue
			}
		case <-c.chanIsClose:
			return
		}
	}
}

// Close 关闭连接
func (c *GConnection) close() {
	if c.isClose {
		return
	}
	if err := c.conn.Close(); err != nil {
		mlog.Error.Println("conn close err : ", err)
	}
	c.isClose = true
	c.chanIsClose <- true

	close(c.chanIsClose)
	close(c.chanSendData)
}

// SendMsg 把Protobuf转化成字节流byte塞到chanSendData里等待发送
func (c *GConnection) SendMsg(msg *ProtoMessage.NetMessage) {
	if msg == nil {
		return
	}
	if c.isClose {
		mlog.Warning.Printf("Client[%s] connection is closed , can't send message!!!!\n", c.conn.RemoteAddr())
		return
	}
	data := PackMessage(msg)
	if data != nil {
		c.chanSendData <- data
	}
}

// Session 返回连接上的会话
func (c *GConnection) Session() *GSession {
	if c.session != nil {
		return c.session
	}
	return nil
}
