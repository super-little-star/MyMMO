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
    public UnityAction<LoginResponse> OnLoginCallback;
    public UnityAction<CreateCharacterResponse> OnCreateCharacterCallback;
    public UnityAction<DeleteCharacterResponse> OnDeleteCharacterCallback;

    public void Init()
    {
        UserManager.Instance.Init();
        MessageHandOut.Instance.Login<RegisterResponse>(OnUserRegister);
        MessageHandOut.Instance.Login<LoginResponse>(OnUserLogin);
        MessageHandOut.Instance.Login<CreateCharacterResponse>(OnCreateCharacter);
    }

    public void Dispose()
    {
        MessageHandOut.Instance.Logout<RegisterResponse>(OnUserRegister);
        MessageHandOut.Instance.Logout<LoginResponse>(OnUserLogin);
        MessageHandOut.Instance.Logout<CreateCharacterResponse>(OnCreateCharacter);
    }

    #region Register
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

        NetSerice.Instance.Send(msg);
    }

    private void OnUserRegister(object message)
    {
        RegisterResponse response = (RegisterResponse)message;
        Debug.LogFormat("OnUserRegister:: Result[{0}],Message[{1}]", response.Result, response.Error);

        OnRegisterCallback?.Invoke(response.Result, response.Error);
    }
    #endregion


    #region Login
    public void SendUserLogin(string userName,string password)
    {
        Debug.LogFormat("SendUserLogin:: UserName[{0}],Password[{1}]",userName,password);

        NetMessage msg = new()
        {
            Request = new()
            {
                Login = new()
                {
                    userName = userName,
                    Passward = password
                }
            }
        };
        
        NetSerice.Instance.Send(msg);
    }

    private void OnUserLogin(object message)
    {
        LoginResponse response = (LoginResponse)message;
        Debug.LogFormat("OnUserLogin:: Result[{0}],Message[{1}]",response.Result, response.Error);

        this.OnLoginCallback?.Invoke(response);
    }
    #endregion

    #region CreateCharacter
    public void SendCreateCharacter(int characterClass , string name)
    {
        Debug.LogFormat("SendCreateCharacter:: Class[{0}] , Name[{1}]",characterClass,name);
        NetMessage msg = new()
        {
            Request = new()
            {
                CreateCharacter = new()
                {
                    Name = name,
                    characterClass = (ProtoMessage.CharacterClass)characterClass
                }
            }
        };

        NetSerice.Instance.Send(msg);
    }

    private void OnCreateCharacter(object message)
    {
        CreateCharacterResponse response = (CreateCharacterResponse)message;
        Debug.LogFormat("OnCreateCharacter:: Result[{0}] , Error[{1}], Character[{2}]", response.Result, response.Error, response.Characters);

        this.OnCreateCharacterCallback?.Invoke(response);
    }
    #endregion

    #region DelectCharacter
    public void SendDeleteCharacter(int characterId)
    {
        Debug.LogFormat("SendDelectCharacter:: CharacterId[{0}]",characterId);
        NetMessage msg = new()
        {
            Request = new()
            {
                DeleteCharacter = new()
                {
                    characterId = characterId
                }
            }
        };

        NetSerice.Instance.Send(msg);
    }

    public void OnDeleteCharacter(object message)
    {
        DeleteCharacterResponse response = (DeleteCharacterResponse)message;
        Debug.LogFormat("OnDeleteCharacter:: Result[{0}] Error[{1}] Characters[{2}]",response.Result,response.Error, response.Characters);

        this.OnDeleteCharacterCallback?.Invoke(response);
    }

    #endregion
}
