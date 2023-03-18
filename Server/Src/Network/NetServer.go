package Network

import (
	"mmo_server/utils/mlog"
)

type GNetServer struct {
	ServerListener *GNetTCPListener
}

func (ns *GNetServer) Init(address string) {
	ns.ServerListener = NewTCPListener(address)
	mlog.Info.Printf("Start Listen success. Listen to [%s]", ns.ServerListener.addr.String())
}

func (ns *GNetServer) Start() {
	//接受连接过来的客户端
	go ns.acceptConn()
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

func (ns *GNetServer) Stop() {
	ns.ServerListener.Close()
}
