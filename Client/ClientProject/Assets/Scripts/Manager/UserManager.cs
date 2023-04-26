using ProtoMessage;
using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class UserManager : Singleton<UserManager>, IDisposable
{
    public void Init()
    {
        UserSerice.Instance.OnRegisterCallback += OnRegisterCallback;
        UserSerice.Instance.OnLoginCallback += OnLoginCallback;
    }
    public void Dispose()
    {
        UserSerice.Instance.OnRegisterCallback -= OnRegisterCallback;
        UserSerice.Instance.OnLoginCallback -= OnLoginCallback;
    }

    private void OnRegisterCallback(Result result,Error error)
    {
        switch (result)
        {
            case Result.Success:
                UIInfoPopup pop = UIManager.Instance.InfoPopup(UIPopup.Level.Normal, "ע��ɹ�");
                pop.AddComfirmEvent(() =>
                {
                    UIManager.Instance.Close(typeof(UIRegister));
                    UIManager.Instance.Open<UILogin>();
                });
                break;
            case Result.Failed:
                if(error == Error.UserNameExist)
                {
                    UIInfoPopup p = UIManager.Instance.InfoPopup(UIPopup.Level.Error, "�û����Ѵ���");
                    p.AddComfirmEvent(() =>
                    {
                        UIRegister r = UIManager.Instance.Open<UIRegister>();
                        r.ResetUI();
                    });
                }
                break;
        }
    }

    private void OnLoginCallback(LoginResponse response)
    {
        switch(response.Result)
        {
            case Result.Success:
                User.Instance.SetUser(response.User);
                SceneManager.Instance.LoadSelectCharacter();
                break;
            case Result.Failed:
                UIManager.Instance.InfoPopup(UIPopup.Level.Error,ProtoErr2string.Convert(response.Error)); ;
                break;
        }
    }

}
