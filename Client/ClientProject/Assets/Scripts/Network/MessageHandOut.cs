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

        public delegate void MessageEvent(object message);
        private Dictionary<string,Delegate>messageEvents = new Dictionary<string,Delegate>();


        public void HandOut(NetMessage message)
        {
            if (message == null) return;

            
            Type t = message.Response.GetType();


            PropertyInfo[] ps = t.GetProperties();
            // 通过反射遍历获取Response下的所有字段
            foreach (PropertyInfo p in ps)
            {
                
                var value = p.GetValue(message.Response);
               
                if(value == null) continue;
                p.GetType();
                // 触发对应的事件
                TriggerEvent(value);
            }
        }
        
        /// <summary>
        /// 触发信息对应的事件
        /// </summary>
        /// <typeparam name="T"></typeparam>
        /// <param name="message"></param>
        public void TriggerEvent(object message)
        {
            string key = message.GetType().Name;
            if(messageEvents.ContainsKey(key))
            {
                MessageEvent e = (MessageEvent)messageEvents[key];
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
        public void Login<T>(MessageEvent messageEvent)
        {
            string type = typeof(T).Name;
            if(!messageEvents.ContainsKey(type))
            {
                messageEvents[type] = null;
            }
            messageEvents[type] = (MessageEvent)messageEvents[type] + messageEvent; ;
        }

        /// <summary>
        /// 注销消息对应的事件
        /// </summary>
        /// <typeparam name="T">信息类型</typeparam>
        /// <param name="messageEvent">信息对应的事件</param>
        public void Logout<T>(MessageEvent messageEvent ) {
            string type = typeof(T).Name;
            if (!messageEvents.ContainsKey(type))
            {
                messageEvents[type] = null;
            }
            messageEvents[type] = (MessageEvent)messageEvents[type] - messageEvent; ;

        }
    }
}
