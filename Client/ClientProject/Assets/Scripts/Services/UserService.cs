using Network;
using ProtoMessage;
using System;
using System.Collections;
using System.Collections.Generic;
using UnityEditor.PackageManager;
using UnityEngine;
using UnityEngine.Events;

public class UserSerice : Singleton<UserSerice>, IDisposable
{
    public UnityAction<Result,ProtoMessage.Error> OnRegisterCallback;

    public void Init()
    {
        MessageHandOut.Instance.Login<NUserRegisterResponse>(OnUserRegister);
    }

    public void Dispose()
    {
        MessageHandOut.Instance.Logout<NUserRegisterResponse>(OnUserRegister);
    }


    public void SendUserRegister(string userName,string password)
    {
        Debug.LogFormat("SendUserRegister:: UserName[{0}],Password[{1}]",userName,password);

        NetMessage msg = new()
        {
            Request = new()
        };
        msg.Request.userRegister = new()
        {
            userName = userName,
            Passward = password
        };

        NetClient.Instance.Send(msg);
    }

    private void OnUserRegister(object message)
    {
        NUserRegisterResponse response = (NUserRegisterResponse)message; 
        Debug.LogFormat("OnUserRegister:: Result[{0}],Message[{1}]", response.Result, response.Error);

        OnRegisterCallback?.Invoke(response.Result, response.Error);
    }
}
