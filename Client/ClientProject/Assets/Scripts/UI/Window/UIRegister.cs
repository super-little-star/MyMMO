
using DG.Tweening;
using System.Collections;
using System.Collections.Generic;
using System.ServiceModel;
using System.Text.RegularExpressions;
using UnityEngine;
using UnityEngine.UI;


public class UIRegister : UIWindow
{
    public InputField IF_UserName;
    public InputField IF_Password;
    public InputField IF_Comfirm;

    public Toggle T_Agree;

    public Button Btn_Register;

    //字母开头，允许5-16字节，允许字母数字
    public static  string UserNameRegular = @"^[a-zA-Z][a-zA-Z0-9]{4,16}$";
    //必须包含大小写字母和数字的组合，不能使用特殊字符，长度在8-16之间
    public static string PasswordRegular = @"^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{8,16}$";

    private bool IsNull
    {
        get
        {
            bool val = IF_UserName == null || IF_Password == null || IF_Comfirm == null || T_Agree == null || Btn_Register == null;

            if (val) Debug.LogWarning("Something is null!!");
            return val;
        }
    }

    public override void Open()
    {
        base.Open();
        if (Btn_Register != null) Btn_Register.onClick.AddListener(OnRegisterClick);
    }

    

    /// <summary>
    /// 注册按钮按下触发事件
    /// </summary>
    private void OnRegisterClick()
    {
        if (IsNull) return;


        if (IsInputEmpty())
        {
            UIManager.Instance.InfoPopup(UIPopup.Level.Warnning, "输入不能为空");
            return;
        }


        if(!T_Agree.isOn)
        {
            UIManager.Instance.InfoPopup(UIPopup.Level.Warnning, "请勾选同意用户协定");
            return;
        }
        if(!Regex.IsMatch(IF_UserName.text, UserNameRegular))
        {
            UIManager.Instance.InfoPopup(UIPopup.Level.Warnning, "用户名存在非法输入");
            return;
        }
        if(!Regex.IsMatch(IF_Password.text, PasswordRegular))
        {
            UIManager.Instance.InfoPopup(UIPopup.Level.Warnning, "密码存在非法输入");
            return;
        }

        if(IF_Password.text != IF_Comfirm.text)
        {
            UIManager.Instance.InfoPopup(UIPopup.Level.Warnning, "两次密码输入不一致");
            return;
        }

        UserSerice.Instance.SendUserRegister(IF_UserName.text,IF_Password.text);

    }


    public override void Close()
    {
        base.Close();
        if (Btn_Register != null) Btn_Register.onClick.RemoveAllListeners();
        UIManager.Instance.Open<UILogin>();
    }

    public override void Hide()
    {
        base.Hide();
        if (Btn_Register != null) Btn_Register.onClick.RemoveAllListeners();
    }

    /// <summary>
    /// 检查输入是否为空
    /// </summary>
    /// <returns></returns>
    private bool IsInputEmpty()
    {
        if(IsNull) return false;
        return IF_UserName.text == string.Empty || IF_Password.text == string.Empty || IF_Comfirm.text == string.Empty;
    }



    public void ResetUI()
    {
        if(IsNull) return;
        IF_UserName.text = string.Empty;
        IF_Comfirm.text = string.Empty;
        IF_Password.text = string.Empty;
    }

}
