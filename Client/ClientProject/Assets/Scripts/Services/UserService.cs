using Network;
using ProtoMessage;
using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.Events;

public class UserSerice : Singleton<UserSerice>, IDisposable
{
    public UnityAction<Nresult, string> OnRegisterCallback;

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

        NetMessage msg = new NetMessage();
        msg.Request = new  NetMessageRequest();
        msg.Request.userRegister = new NUserRegisterRequest();
        msg.Request.userRegister.userName = userName;
        msg.Request.userRegister.Passward = password;

        NetClient.Instance.Send(msg);
    }

    private void OnUserRegister(NUserRegisterResponse response)
    {
        Debug.LogFormat("OnUserRegister:: Result[{0}],Message[{1}]", response.Result, response.Errormsg);

        if(OnRegisterCallback != null) { OnRegisterCallback.Invoke(response.Result,response.Errormsg); }
    }
}
