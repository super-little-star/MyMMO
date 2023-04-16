using ProtoMessage;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using System.Net;
using System.Net.Sockets;
using UnityEngine;
using System;
using System.Threading;
using Unity.VisualScripting;
using UnityEngine.Events;
using UnityEditor;

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
        public bool connecting = false;
        public bool isShowConnecting = false;

        public bool Connected
        {
            get
            {
                return (this.clientSocket != default(Socket)) ? this.clientSocket.Connected : false;
            }
        }
        private int retryTimes = 0;
        private int retryTimesCount = 3;
        private Thread ConnectThread = null;
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

        #region Event 事件

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

        public void Connecting(bool isConnecting)
        {
            if(isConnecting)
            {
                UIManager.Instance.Popup<UIWaitPopup>("正在连接服务器");
            }
            else
            {
                UIManager.Instance.Close(typeof(UIWaitPopup));
            }
        }
        
        public void Update()
        {
            if (this.KeepConnect())
            {
                if (this.ReadMessage())
                {
                    if (this.WriteMessage())
                    {
                        MessageHandleCenter.Instance.MessageDelivery();// 把消息队列里的信息进行分发
                    }
                }
            }
        }

        /// <summary>
        /// 保存连接
        /// </summary>
        /// <returns></returns>
        private bool KeepConnect()
        {
            if (this.isShowConnecting!=this.connecting)
            {
                if (this.connecting) UIManager.Instance.Popup<UIWaitPopup>("正在连接服务器");
                else UIManager.Instance.Close(typeof(UIWaitPopup));
                this.isShowConnecting = this.connecting;
            }
            if (this.connecting) return false;


            if (this.address == null) return false;

            if (this.Connected) return true;

            if(this.ConnectThread != null) return false;

            ConnectToServer();
            return false;
        }

        /// <summary>
        /// 关闭连接
        /// </summary>
        /// <param name="errorCode">错误代码</param>
        public void Close(int errorCode)
        {
            Debug.LogWarning("Close Connection ,erroro code : "+errorCode.ToString()+"\n");

            if(this.clientSocket != null)
            {
                this.clientSocket.Close();
            }

            this.sendQueue.Clear();

            this.ResetAll();

            //TODO 错误处理
        }



        /// <summary>
        /// 连接到服务器
        /// </summary>
        /// <exception cref="Exception"></exception>
        private void ConnectToServer()
        {
            if (this.connecting) return;// 检查是否正在连接

            if (this.clientSocket != null) { this.clientSocket.Close(); }// 如果已经存在Socket，先断开

            if (this.address == default(IPEndPoint)) { throw new Exception("Please Init address....."); }// 服务器的地址没有初始化

            


            if (retryTimes < retryTimesCount)
            {
                // 开启一个线程连接
                this.ConnectThread = new Thread(DoConnect);
                this.ConnectThread.Start();
            }
            else
            {
                UIComfirmPopup pop = UIManager.Instance.Popup<UIComfirmPopup>(UIPopup.Level.Error, "连接服务器失败", "连接服务器失败！请检查网络并重试", "重试", "退出");
                pop.Btn_Comfirm.onClick.AddListener(() => { this.retryTimes = 0; });
                pop.Btn_Cancel.onClick.AddListener(this.Quit);
            }
        }

        private void Quit()
        {
#if UNITY_EDITOR
            EditorApplication.isPlaying = false;
#else
            Application.Quit();
#endif

        }
        private void DoConnect()
        {
            Debug.Log("NetClient Do Connect To " + this.address.ToString());
            this.connecting = true;
            // 连接到服务器
            try
            {
                this.clientSocket = new(AddressFamily.InterNetwork, SocketType.Stream, ProtocolType.Tcp)
                {
                    Blocking = true
                };

                Debug.Log(string.Format("Connect to server {0}", this.address));
                IAsyncResult result = this.clientSocket.BeginConnect(this.address, null, null);
                bool success = result.AsyncWaitHandle.WaitOne(NET_CONNECT_TIMEOUT);
                if (success)
                {
                    this.clientSocket.EndConnect(result);
                }
            }
            catch (SocketException se)
            {
                if (se.SocketErrorCode == SocketError.ConnectionRefused)
                {
                    this.Close(0);
                }
                Debug.LogErrorFormat("Socket Exception : {0},{1},{2}]{3}", se.ErrorCode, se.SocketErrorCode, se.NativeErrorCode, se.Message);
            }
            catch (Exception e)
            {
                Debug.LogError("Do Connect Exception: " + e.ToString() + "\n");
            }


            if (this.clientSocket.Connected)
            {
                this.clientSocket.Blocking = false;
                
            }
            else
            {
                this.retryTimes++;
                Debug.LogWarningFormat("Retry[{0}] To Connect to service", retryTimes);
                if (this.retryTimes >= this.retryTimesCount)
                {

                }
            }

            Debug.Log("finish do connection");
            this.connecting = false;

            this.ConnectThread = null;
        }

        /// <summary>
        /// 发送Protobuf
        /// </summary>
        /// <param name="msg">Protobuf</param>
        public void Send(NetMessage msg)
        {
            if (!running) return;

            if(!this.Connected)
            {
                //从新连接
                this.ResetBuf();
                this.ConnectToServer();
                return;
            }
            this.sendQueue.Enqueue(msg);

            this.lastSendTime = Time.time;
        }

        #region Socket
        private bool IsSocketError()
        {
            bool isError = this.clientSocket.Poll(0, SelectMode.SelectError);
            if (isError)
            {
                Debug.Log("Client Socket Poll Select Error\n");
            }
            return isError;
        }

        /// <summary>
        /// 把Socket上的信息提取出来
        /// </summary>
        /// <returns></returns>
        private bool ReadMessage()
        {
            try
            {
                if (IsSocketError()) return false;

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

        /// <summary>
        /// 把信息写到Socket上
        /// </summary>
        /// <returns></returns>
        private bool WriteMessage()
        {

            try
            {
                if (IsSocketError()) return false;// 判断socket是否有错误

                bool res = this.clientSocket.Poll(0, SelectMode.SelectWrite);
                if (res)
                {
                    if(this.sendBuffer.Position > this.sendOffset)// 判断是否有信息未发送
                    {
                        int bufsize = (int)(this.sendBuffer.Position - this.sendOffset);
                        int n = this.clientSocket.Send(this.sendBuffer.GetBuffer(), this.sendOffset,bufsize, SocketFlags.None);
                        if (n <= 0)
                        {
                            this.Close(0);
                            return false;
                        }
                        this.sendOffset += n;
                        if(this.sendOffset >=this.sendBuffer.Position)
                        {
                            this.sendOffset = 0;
                            this.sendBuffer.Position = 0;
                            this.sendQueue.Dequeue();
                        }
                    }
                    else
                    {
                        if(this.sendQueue.Count > 0)
                        {
                            NetMessage msg = this.sendQueue.Peek();
                            byte[] buf = PackageHandler.PackMessage(msg);
                            this.sendBuffer.Write(buf, 0, buf.Length);
                        }
                    }
                }
            }
            catch (Exception e)
            {
                Debug.Log("Write Message Exception : " + e.ToString() + "\n");
                this.Close(0);
                return false;
            }
            return true;
        }
        #endregion

        #region Reset 
        private void ResetBuf()
        {
            this.receiveBuffer.Position = 0;
            this.sendBuffer.Position = sendOffset = 0;
        }
        private void ResetAll()
        {
            this.ResetBuf();
            this.connecting = false;
            this.lastSendTime = 0;

        }

        private void OnDisable()
        {
            this.Close(0);
        }

        #endregion
    }
}
