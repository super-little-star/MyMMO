package network

import (
	"encoding/binary"
	"io"
	ProtoMessage "mmo_server/ProtoMessage"
	"mmo_server/utils/mlog"
	"net"
)

type GConnection struct {
	session        *GSession
	conn           *net.TCPConn
	packageHandler *PackageHandler

	readBuf []byte

	chanIsClose  chan bool
	chanSendData chan []byte
	isClose      bool
}

func NewConnection(conn *net.TCPConn, session *GSession) *GConnection {
	c := &GConnection{
		conn:    conn,
		session: session,

		readBuf: make([]byte, MaxPackageSize),

		chanIsClose:  make(chan bool, 1),
		chanSendData: make(chan []byte),
		isClose:      false,
	}
	if err := c.conn.SetReadBuffer(MaxPackageSize); err != nil {
		mlog.Error.Printf("Connection Set Read Buffer Error : %s !!!", err)
	}
	if err := c.conn.SetWriteBuffer(MaxPackageSize); err != nil {
		mlog.Error.Printf("Connection Set Write Buffer Error : %s !!!", err)
	}
	c.packageHandler = NewPackageHandler(c)
	return c
}

// ReadMsg 读取连接上的信息
func (c *GConnection) readMsg() {
	mlog.Info.Printf("Client[%s] Read Goroutine is Running ....", c.conn.RemoteAddr())
	defer func() {
		mlog.Info.Printf("Client[%s] Read Goroutine is Close !!!!", c.conn.RemoteAddr())
		c.close()
	}()
	for {
		// 读取头部信息，获取信息长度
		head := make([]byte, 4)
		if _, err := c.conn.Read(head); err != nil {
			if err == io.EOF {
				mlog.Warning.Printf("Client[$s] connection is Close!!!", c.conn.RemoteAddr())
				return
			}
			mlog.Warning.Printf("Connection[%s] Read message error : %v\n", c.conn.RemoteAddr(), err)
			continue
		}

		msgLen := binary.LittleEndian.Uint32(head) // 字节流转换Uint32

		// 读取主要信息
		body := make([]byte, msgLen)
		if _, err := c.conn.Read(body); err != nil {
			if err == io.EOF {
				mlog.Warning.Printf("Client[$s] connection is Close!!!", c.conn.RemoteAddr())
				return
			}
			mlog.Warning.Printf("Connection[%s] Read message error : %v\n", c.conn.RemoteAddr(), err)
			continue
		}

		// 发给PackageHandler处理信息
		if err := c.packageHandler.ReceiveMsg(body, msgLen); err != nil {
			mlog.Warning.Println("package Handler Receive message error : ", err)
			continue
		}

		c.readBuf = c.readBuf[:0] // Reset Read Buf
	}
}

// WriteMsg 将chanSendData里的数据发送给客户端
func (c *GConnection) writeMsg() {
	mlog.Info.Printf("Client[%s] Write Goroutine is Running ....\n", c.conn.RemoteAddr())
	defer mlog.Info.Printf("Client[%s] Write Goroutine is Close !!!!\n", c.conn.RemoteAddr())
	defer c.close()
	for {
		select {
		case data := <-c.chanSendData: // 把需要发送的信息取出来发送出去
			if _, err := c.conn.Write(data); err != nil {
				mlog.Error.Printf("Client[%s] connection write data is error : %v\n", c.conn.RemoteAddr(), err)
				break
			}
			mlog.Info.Printf("Client[%s] connection write data is success...", c.conn.RemoteAddr())
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

	c.session.Disconnected()

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
