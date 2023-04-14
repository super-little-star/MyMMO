
using DG.Tweening;
using System.Collections;
using System.Collections.Generic;
using System.ServiceModel;
using System.Text.RegularExpressions;
using UnityEngine;
using UnityEngine.UI;


public class UIRegister : MonoBehaviour
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
    private void Start()
    {
        if(Btn_Register!=null) Btn_Register.onClick.AddListener(OnRegisterClick);
    }

    

    /// <summary>
    /// 注册按钮按下触发事件
    /// </summary>
    private void OnRegisterClick()
    {
        if (IsNull) return;

        this.Show();
        if (IsInputEmpty())
        {
            // TODO 弹窗说明
            return;
        }


        if(!T_Agree.isOn)
        {
            // TODO 弹窗说明
            return;
        }
        if(!Regex.IsMatch(IF_UserName.text, UserNameRegular))
        {
            // TODO 弹窗说明
            return;
        }
        if(!Regex.IsMatch(IF_Password.text, PasswordRegular))
        {
            // TODO 弹窗说明
            return;
        }

        if(IF_Password.text != IF_Comfirm.text)
        {
            // TODO 弹窗说明
            return;
        }

        UserSerice.Instance.SendUserRegister(IF_UserName.text,IF_Password.text);

       

    }

    public void Hide()
    {
        this.transform.gameObject.SetActive(false);
    }
    public void Show()
    {
        this.transform.gameObject.SetActive(true);
        if (Btn_Register != null) Btn_Register.interactable = false;
        transform.localScale = new Vector3(0.6f, 0.6f, 0.6f);
        transform.DOScale(1f, 0.3f).SetEase(Ease.OutElastic).OnComplete(() => { if (Btn_Register != null) Btn_Register.interactable = true; });
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
