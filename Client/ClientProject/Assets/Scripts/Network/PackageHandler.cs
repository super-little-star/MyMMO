using System.IO;
using System;
using ProtoMessage;
namespace Network
{
    public class PackageHandler
    {
        private MemoryStream stream = new MemoryStream(64 * 1024);
        private int readOffset = 0;


        public void ReceiveMsg(byte[] data)
        {
            if (stream.Position + data.Length > stream.Capacity)
            {
                throw new Exception("Package Handler buffer overflow");
            }
            stream.Write(data, 0, data.Length);

            ParsePackage();
        }

        private void ParsePackage()
        {
            while (this.readOffset < stream.Position)
            {
                int dataSize = BitConverter.ToInt32(stream.GetBuffer(), this.readOffset);
                if (dataSize + readOffset + 4 <= stream.Position)// 包有效检查
                {
                    NetMessage message = UnpackMessage(stream.GetBuffer(), this.readOffset + 4, dataSize);
                    if (message == null)
                    {
                        throw new Exception("PackageHander ParsePackage is error!!!");
                    }

                    MessageHandleCenter.Instance.AcceptMessage(message); ;

                    this.readOffset += (dataSize + 4);
                }
            }

            this.readOffset = 0;
            stream.Position = 0;
            stream.SetLength(0);
        }

        /// <summary>
        /// 将Protobuf转换成字节流
        /// </summary>
        /// <param name="message">Protobuf</param>
        /// <returns></returns>
        public static byte[] PackMessage(NetMessage message)
        {
            byte[] data = null;

            using (MemoryStream msg = new MemoryStream())// 创建临时流来处理，用完马上释放
            {
                ProtoBuf.Serializer.Serialize(msg, message);
                data = new byte[msg.Length + 4];
                Buffer.BlockCopy(BitConverter.GetBytes(msg.Length), 0, data, 0, 4);
                Buffer.BlockCopy(msg.GetBuffer(), 0, data, 4, (int)msg.Length);
            }
            return data;
        }

        public static NetMessage UnpackMessage(byte[] data, int offset, int len)
        {
            NetMessage message = null;
            using (MemoryStream msg = new MemoryStream(data, offset, len))
            {
                message = ProtoBuf.Serializer.Deserialize<NetMessage>(msg);
            }
            return message;
        }
    }

}
