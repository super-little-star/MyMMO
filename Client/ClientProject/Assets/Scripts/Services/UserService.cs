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
    public UnityAction<Result, ProtoMessage.Error> OnRegisterCallback;

    public void Init()
    {
        UserManager.Instance.Init();
        MessageHandOut.Instance.Login<RegisterResponse>(OnUserRegister);
    }

    public void Dispose()
    {
        MessageHandOut.Instance.Logout<RegisterResponse>(OnUserRegister);
    }


    public void SendUserRegister(string userName,string password)
    {
        Debug.LogFormat("SendUserRegister:: UserName[{0}],Password[{1}]",userName,password);

        NetMessage msg = new()
        {
            Request = new()
        };
        msg.Request.Register = new()
        {
            userName = userName,
            Passward = password
        };

        NetWork.Instance.Send(msg);
    }

    private void OnUserRegister(object message)
    {
        RegisterResponse response = (RegisterResponse)message;
        Debug.LogFormat("OnUserRegister:: Result[{0}],Message[{1}]", response.Result, response.Error);

        OnRegisterCallback?.Invoke(response.Result, response.Error);
    }

    public void SendUserLogin(string userName,string password)
    {
        Debug.LogFormat("SendUserLogin:: UserName[{0}],Password[{1}]",userName,password);

        NetMessage msg = new()
        {
            Request = new()
            {
                
            }
        };
    }
}
