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

    public override void Open()
    {
        base.Open();
        this.ResetUI();
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

        if(IsInputEmpty())
        {
            UIManager.Instance.InfoPopup(UIPopup.Level.Error, "用户名&密码不能为空");
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
