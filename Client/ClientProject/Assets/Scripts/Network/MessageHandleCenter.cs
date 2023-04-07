using ProtoMessage;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;


namespace Network
{
    public class MessageHandleCenter : Singleton<MessageHandleCenter>,IDisposable
    {
        public MessageHandleCenter() { }

        private Queue<NetMessage> messageQueue = new Queue<NetMessage>();


        public void MessageDelivery()
        {
            if (messageQueue.Count == 0) return;

            while (messageQueue.Count > 0)
            {
                NetMessage msg = messageQueue.Dequeue();


                if(msg.Response!= null)
                {
                    MessageHandOut.Instance.HandOut(msg);
                }
            }
        }

        public void AcceptMessage(NetMessage message)
        {
            if(message == null) return;
            this.messageQueue.Enqueue(message);
        }

        public void Dispose()
        {
            messageQueue.Clear();
        }
    }
}

