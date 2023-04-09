package network

import (
	"bytes"
	"encoding/binary"
	"errors"
	"google.golang.org/protobuf/proto"
	ProtoMessage "mmo_server/protocol"
	"mmo_server/utils/mlog"
)

const MaxPackageSize int = 64 * 1024

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
func (ph *PackageHandler) ReceiveMsg(data []byte, msgLen uint32) error {
	//检查数据长度是否超出规定长度
	if ph.stream.Cap()+len(data)+4 > MaxPackageSize {
		return errors.New("PackageHandler : buffer overflow")
	}
	//写到缓存里
	ph.stream.Write(data)
	//将缓存中的数据进行解析
	if err := ph.parsePackage(msgLen); err != nil {
		return err
	}
	return nil
}

// 解析Client发送过来的包
func (ph *PackageHandler) parsePackage(msgLen uint32) error {

	//把Protobuf字节流的的部分进行解包，转化成protobuf对象
	buf := UnpackMessage(ph.stream, msgLen)

	// 把转换到的Protobuf对象传给消息处理中心处理
	Instance().MessageHandleCenter.AcceptMessage(ph.sender, buf)

	//包体读取完毕，重置缓存
	ph.stream.Reset()
	return nil
}

// PackMessage 将Protobuf转换成字节流
func PackMessage(message *ProtoMessage.NetMessage) []byte {
	//把Protobuf转化成字节流
	marshal, err := proto.Marshal(message)
	if err != nil {
		mlog.Error.Println("Protobuf Marshal data is error : ", err)
		return nil
	}
	ms := bytes.NewBuffer([]byte{})
	//把字节长度写在数据头部
	if err := binary.Write(ms, binary.LittleEndian, uint32(len(marshal))); err != nil {
		mlog.Error.Println("binary write data is error : ", err)
		return nil
	}
	//把protobuf字节流部分写到后面
	if err := binary.Write(ms, binary.LittleEndian, marshal); err != nil {
		mlog.Error.Println("binary write data is error : ", err)
		return nil
	}
	return ms.Bytes()
}

// UnpackMessage 将字节流转换成Protobuf
func UnpackMessage(pack *bytes.Buffer, len uint32) *ProtoMessage.NetMessage {
	message := &ProtoMessage.NetMessage{}
	data := make([]byte, len)
	//将protobuf字节流的部分读出来，放到data里
	if err := binary.Read(pack, binary.LittleEndian, &data); err != nil {
		mlog.Error.Println("Unpack message is error : ", err)
		return nil
	}
	//在对protobuf字节流转换成protobuf对象
	if err := proto.Unmarshal(data, message); err != nil {
		mlog.Error.Println("Protobuf Unmarshal is error : ", err)
		return nil
	}
	return message
}
