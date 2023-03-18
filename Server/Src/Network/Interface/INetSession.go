package Network

type INetSession interface {
	GetResponse() []byte
	Disconnected()
}
