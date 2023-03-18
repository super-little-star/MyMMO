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

            Connect();
            return false;
        }

        public void Connect()
        {
            if (this.connecting) return;

            if (this.clientSocket != null) { this.clientSocket.Close(); }

            if (this.address == default(IPEndPoint)) { throw new Exception("Please Init address....."); }

            Debug.Log("DoConnect");
            this.connecting = true;

            this.DoConnect();
        }

        private void DoConnect()
        {
            Debug.Log("NetClient Do Connect On " + this.address.ToString());

            try 
            {
                if (this.clientSocket != null) { this.clientSocket.Close(); }

                this.clientSocket = new Socket(AddressFamily.InterNetwork, SocketType.Stream, ProtocolType.Tcp);
                this.clientSocket.Blocking = true;

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
                Debug.Log("Do Connect Exception: " + e.ToString());
            }


            if (this.clientSocket.Connected)
            {
                this.clientSocket.Blocking = false;
            }
            this.connecting = false;
        }
    }
}
