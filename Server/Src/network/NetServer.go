package network

import (
	"mmo_server/utils/mlog"
)

type GNetServer struct {
	ServerListener *GNetListener
}

func (ns *GNetServer) Init(network string, address string) {
	ns.ServerListener = NewListener(network, address)
	mlog.Info.Printf("Start Listen success. Listen to [%s]", ns.ServerListener.addr.String())
}

func (ns *GNetServer) Start() {
	go ns.acceptConn() // 开启一个协程接受客户端的链接
}

func (ns *GNetServer) Stop() {
	ns.ServerListener.Close()
}

// 接受来自客户端的链接
func (ns *GNetServer) acceptConn() {
	for {
		conn := ns.ServerListener.AcceptConn()
		if conn == nil {
			continue
		}

	}
}
