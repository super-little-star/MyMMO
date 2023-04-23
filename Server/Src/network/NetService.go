package network

import (
	"mmo_server/utils/mlog"
)

type GNetService struct {
	isRunning      bool
	ServerListener *GNetTCPListener
}

// Init 初始化网络服务
func (ns *GNetService) Init(network string, address string) {
	ns.isRunning = false
	ns.ServerListener = NewListener(network, address)
	mlog.Info.Printf("Start Listen success. Listen to [%s]", ns.ServerListener.addr.String())
}

// Start 开启网络服务
func (ns *GNetService) Start() {
	ns.isRunning = true
	MessageHandleCenter().Start(10) //开启10个Goroutine处理消息
	go ns.acceptConn()              // 开启一个协程接受客户端的链接
}

// Stop 关闭网络服务
func (ns *GNetService) Stop() {
	ns.isRunning = false
	ns.ServerListener.Close()
}

// 接受来自客户端的链接
func (ns *GNetService) acceptConn() {
	for ns.isRunning {
		conn := ns.ServerListener.AcceptConn()
		if conn == nil {
			continue
		}
		session := NewSession()
		connection := NewConnection(conn, session)
		mlog.Info.Printf("Client[%s] is Connected....\n", connection.conn.RemoteAddr())

		go connection.readMsg()
		go connection.writeMsg()
	}
}
