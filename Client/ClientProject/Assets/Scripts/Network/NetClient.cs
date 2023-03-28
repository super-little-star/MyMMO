using ProtoMessage;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using System.Net;
using System.Net.Sockets;
using UnityEngine;
using System;

namespace Network
{
    public class NetClient : MonoSingleton<NetClient>
    {
        public NetClient() { }

        PackageHandler packageHandler = new PackageHandler();

        #region const
        /// <summary>
        /// 能大接受的包体大小
        /// </summary>
        public const int DEF_RECV_BUFFER_SIZE = 64 * 1024;
        /// <summary>
        /// 连接超时时间，毫秒
        /// </summary>
        public const int NET_CONNECT_TIMEOUT = 10000;
        #endregion

        #region 状态属性
        /// <summary>
        /// 脚本是否在运行
        /// </summary>
        private bool running = false;
        /// <summary>
        /// 是否正在连接
        /// </summary>
        private bool connecting = false; 
        public bool Connected
        {
            get
            {
                return (this.clientSocket != default(Socket)) ? this.clientSocket.Connected : false;
            }
        }
        #endregion

        #region Socket 实体
        private IPEndPoint address;
        private Socket clientSocket;
        private MemoryStream sendBuffer = new MemoryStream();
        private MemoryStream receiveBuffer = new MemoryStream(DEF_RECV_BUFFER_SIZE);
        private Queue<NetMessage> sendQueue = new Queue<NetMessage>();
        private int sendOffset = 0;
        private float lastSendTime = 0;
        #endregion

        /// <summary>
        /// 初始化客户端连接
        /// </summary>
        /// <param name="serverIP">服务器IP</param>
        /// <param name="port">端口</param>
        public void Init(string serverIP,int port)
        {
            this.address = new IPEndPoint(IPAddress.Parse(serverIP), port);
        }
        
        protected override void OnStart()
        {
            this.running= true;
        }

        private void Update()
        {
            if (!running) return;

            if (this.KeepConnect())
            {

            }
        }

        /// <summary>
        /// 保存连接
        /// </summary>
        /// <returns></returns>
        private bool KeepConnect()
        {
            if (this.connecting) return false;

            if (this.address == null) return false;

            if (this.Connected) return true;

            ConnectToServer();
            return false;
        }

        public void Close(int errorCode)
        {
            Debug.LogWarning("Close Connection ,erroro code : "+errorCode.ToString()+"\n");
            this.connecting = false;
            if(this.clientSocket != null)
            {
                this.clientSocket.Close();
            }

            this.sendQueue.Clear();

            this.receiveBuffer.Position = 0;
            this.sendBuffer.Position = sendOffset = 0;

            //TODO 错误处理
        }

        /// <summary>
        /// 连接到服务器
        /// </summary>
        /// <exception cref="Exception"></exception>
        public void ConnectToServer()
        {
            if (this.connecting) return;// 检查是否正在连接

            if (this.clientSocket != null) { this.clientSocket.Close(); }// 如果已经存在Socket，先断开

            if (this.address == default(IPEndPoint)) { throw new Exception("Please Init address....."); }// 服务器的地址没有初始化

            this.connecting = true;

            Debug.Log("NetClient Do Connect To " + this.address.ToString());

            // 连接到服务器
            try
            {
                this.clientSocket = new Socket(AddressFamily.InterNetwork, SocketType.Stream, ProtocolType.Tcp);
                this.clientSocket.Blocking = false;

                Debug.Log(string.Format("Connect to server {0}\n", this.address));
                IAsyncResult result = this.clientSocket.BeginConnect(this.address, null, null);
                bool success = result.AsyncWaitHandle.WaitOne(NET_CONNECT_TIMEOUT);
                if (success)
                {
                    this.clientSocket.EndConnect(result);
                }
            }
            catch (Exception e)
            {
                Debug.Log("Do Connect Exception: " + e.ToString() + "\n");
            }


            if (this.clientSocket.Connected)
            {
                this.clientSocket.Blocking = false;
            }
            this.connecting = false;
        }

        private bool IsSockeError()
        {
            bool isError = this.clientSocket.Poll(0, SelectMode.SelectError);
            if (isError)
            {
                Debug.Log("Client Socket Poll Select Error\n");
            }
            return isError;
        }

        private bool ReadMessage()
        {
            try
            {
                if (IsSockeError()) return false;

                bool res = this.clientSocket.Poll(0, SelectMode.SelectRead);
                if (res)
                {
                    int n = this.clientSocket.Receive(this.receiveBuffer.GetBuffer(), 0, this.receiveBuffer.Capacity, SocketFlags.None);
                    if (n <= 0)
                    {
                        this.Close(0);
                        return false;
                    }
                    this.packageHandler.ReceiveMsg(this.receiveBuffer.GetBuffer());
                }
            }
            catch(Exception e)
            {
                Debug.Log("Read Message Exception : " + e.ToString() + "\n");
                this.Close(0);
                return false;
            }
            return true;
            
        }
    }
}
