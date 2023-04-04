using ProtoMessage;
using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.Reflection;
using System.Text;
using System.Threading.Tasks;
using static Network.MessageHandOut;

namespace Network
{
    public class MessageHandOut : Singleton<MessageHandOut>
    {
        public MessageHandOut() { }

        public delegate void MessageEvent<T>(T message);
        private Dictionary<string,Delegate>messageEvents = new Dictionary<string,Delegate>();


        public void HandOut(NetMessage message)
        {
            if (message == null) return;

            
            Type type = message.Response.GetType();
            // 通过反射遍历获取Response下的所有字段
            foreach (FieldInfo f in type.GetFields())
            {
                var value = f.GetValue(message.Response);
                if(value == null) continue;
                // 触发对应的事件
                TriggerEvent(value);
            }
        }
        
        /// <summary>
        /// 触发信息对应的事件
        /// </summary>
        /// <typeparam name="T"></typeparam>
        /// <param name="message"></param>
        public void TriggerEvent<T>(T message)
        {
            string key = message.GetType().Name;
            if(messageEvents.ContainsKey(key))
            {
                MessageEvent<T> e = (MessageEvent<T>)messageEvents[key];
                if(e != null)
                {
                    try
                    {
                        e(message);
                    }
                    catch(Exception ex)
                    {
                        Debug.Print("Message Event exception : ",ex.Message);
                    }
                    
                }
                else
                {
                    Debug.Print("Message Event is Null : " + key);
                }
            }
        }

        /// <summary>
        /// 注册对应信息的事件
        /// </summary>
        /// <typeparam name="T">信息类型</typeparam>
        /// <param name="e">对应的事件</param>
        public void Login<T>(MessageEvent<T> messageEvent)
        {
            string type = typeof(T).Name;
            if(!messageEvents.ContainsKey(type))
            {
                messageEvents[type] = null;
            }
            messageEvents[type] = (MessageEvent<T>)messageEvents[type] + messageEvent; ;
        }

        /// <summary>
        /// 注销消息对应的事件
        /// </summary>
        /// <typeparam name="T">信息类型</typeparam>
        /// <param name="messageEvent">信息对应的事件</param>
        public void Logout<T>(MessageEvent<T> messageEvent ) {
            string type = typeof(T).Name;
            if (!messageEvents.ContainsKey(type))
            {
                messageEvents[type] = null;
            }
            messageEvents[type] = (MessageEvent<T>)messageEvents[type] - messageEvent; ;

        }
    }
}
