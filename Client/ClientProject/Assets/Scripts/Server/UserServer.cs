using Network;
using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class UserServer : Singleton<UserServer>, IDisposable
{

    public void Init()
    {

    }
    public void ConnectToServer()
    {
        Debug.Log("Connecting to server...");
        NetClient.Instance.Init("127.0.0.1", 7788);
        NetClient.Instance.ConnectToServer();
    }
    public void Dispose()
    {
        
    }
}
