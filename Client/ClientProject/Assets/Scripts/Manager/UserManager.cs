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
        UserSerice.Instance.OnCreateCharacterCallback += OnCreateCharacterCallback;
    }
    public void Dispose()
    {
        UserSerice.Instance.OnRegisterCallback -= OnRegisterCallback;
        UserSerice.Instance.OnLoginCallback -= OnLoginCallback;
        UserSerice.Instance.OnCreateCharacterCallback -= OnCreateCharacterCallback;
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
                if(error == Error.RegisterUserNameExist)
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

    private void OnCreateCharacterCallback(CreateCharacterResponse response)
    {
        switch(response.Result)
        {
            case Result.Success:
                User.Instance.SetCharacters(response.Characters);
                UIManager.Instance.InfoPopup(UIPopup.Level.Normal, "������ɫ�ɹ�").AddComfirmEvent(() =>
                {
                    UIManager.Instance.Open<UISelectCharacter>(false);
                    UIManager.Instance.Close(typeof(UICreateCharacter));
                });
                break;
            case Result.Failed:
                UIManager.Instance.InfoPopup(UIPopup.Level.Error, ProtoErr2string.Convert(response.Error));
                break;
        }
    }
}
