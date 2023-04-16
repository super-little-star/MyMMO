using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

public class UILogin : UIWindow
{
    public InputField IF_UserName;
    public InputField IF_Password;

    public Button Btn_Login;
    public Button Btn_Register;

    private bool IsNull
    {
        get
        {
            bool v = IF_UserName == null || IF_Password == null || Btn_Login == null || Btn_Register == null;
            if (v) Debug.LogError("Something is null!!!");
            return v;
        }
    }

    protected override void OnStart()
    {
        base.OnStart();
        if(!IsNull)
        {
            Btn_Login.onClick.AddListener(OnLoginClick);
            Btn_Register.onClick.AddListener(OnRegisterClick);
        }
    }

    private void OnLoginClick()
    {
        UIManager.Instance.Popup<UIWaitPopup>("正在登陆中");
        if(IsInputEmpty())
        {
            // TODO 有空输入处理
            return;
        }


    }

    private void OnRegisterClick()
    {
        UIManager.Instance.Close(typeof(UILogin));
        UIManager.Instance.Open<UIRegister>();
    }


    private bool IsInputEmpty()
    {
        if(IsNull) return false;
        
        return IF_Password.text == string.Empty || IF_UserName.text == string.Empty;
    }

    public void ResetUI()
    {
        IF_UserName.text = string.Empty;
        IF_Password.text = string.Empty;
    }
}
