package network

import (
	"bytes"
	"encoding/binary"
	"errors"
	"google.golang.org/protobuf/proto"
	ProtoMessage "mmo_server/protocol"
	"mmo_server/utils/mlog"
)

const MaxPackageSize uint32 = 64 * 1024

type PackageHandler struct {
	stream *bytes.Buffer
	sender *GConnection
}

func NewPackageHandler(sender *GConnection) *PackageHandler {
	return &PackageHandler{
		stream: bytes.NewBuffer([]byte{}),
		sender: sender,
	}
}

// ReceiveMsg 接受信息
func (ph *PackageHandler) ReceiveMsg(data []byte) error {
	if ph.stream.Cap()+len(data) > int(MaxPackageSize) {
		return errors.New("PackageHandler : buffer overflow")
	}
	ph.stream.Write(data)
	ph.parsePackage()
	return nil
}

// 解析信息
func (ph *PackageHandler) parsePackage() error {
	for ph.stream.Len() > 0 {
		var msgLen uint32 = 0
		if err := binary.Read(ph.stream, binary.LittleEndian, &msgLen); err != nil {
			return err
		}
		_ = UnpackMessage(ph.stream, msgLen)
		// TODO 获取到Protobuf对象，往后再次进行处理
	}
	ph.stream.Reset()
	return nil
}

// PackMessage 将Protobuf转换成字节流
func PackMessage(message *ProtoMessage.NetMessage) ([]byte, error) {
	marshal, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}
	ms := bytes.NewBuffer([]byte{})
	if err := binary.Write(ms, binary.LittleEndian, uint32(len(marshal))); err != nil {
		return nil, err
	}
	if err := binary.Write(ms, binary.LittleEndian, marshal); err != nil {
		return nil, err
	}
	return ms.Bytes(), nil
}

// UnpackMessage 将字节流转换成Protobuf
func UnpackMessage(pack *bytes.Buffer, len uint32) *ProtoMessage.NetMessage {
	message := &ProtoMessage.NetMessage{}
	data := make([]byte, len)
	if err := binary.Read(pack, binary.LittleEndian, &data); err != nil {
		mlog.Error.Println("Unpack message is error : ", err)
		return nil
	}
	if err := proto.Unmarshal(data, message); err != nil {
		mlog.Error.Println("Protobuf Unmarshal is error : ", err)
		return nil
	}
	return message
}
