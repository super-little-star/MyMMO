
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

    //��ĸ��ͷ������5-16�ֽڣ�������ĸ����
    public static  string UserNameRegular = @"^[a-zA-Z][a-zA-Z0-9]{4,16}$";
    //���������Сд��ĸ�����ֵ���ϣ�����ʹ�������ַ���������8-16֮��
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

    protected override void OnStart()
    {
        base.OnStart();
        if (Btn_Register != null) Btn_Register.onClick.AddListener(OnRegisterClick);
    }

    

    /// <summary>
    /// ע�ᰴť���´����¼�
    /// </summary>
    private void OnRegisterClick()
    {
        if (IsNull) return;


        if (IsInputEmpty())
        {
            
            return;
        }


        if(!T_Agree.isOn)
        {
            // TODO ����˵��
            return;
        }
        if(!Regex.IsMatch(IF_UserName.text, UserNameRegular))
        {
            // TODO ����˵��
            return;
        }
        if(!Regex.IsMatch(IF_Password.text, PasswordRegular))
        {
            // TODO ����˵��
            return;
        }

        if(IF_Password.text != IF_Comfirm.text)
        {
            // TODO ����˵��
            return;
        }

        UserSerice.Instance.SendUserRegister(IF_UserName.text,IF_Password.text);

       

    }

    private void OnRegisterCallback(ProtoMessage.Result result,ProtoMessage.Error error)
    {
        if(result == ProtoMessage.Result.Success)
        {
            UIInfoPopup pop = UIManager.Instance.InfoPopup(UIPopup.Level.Normal, "ע��ɹ�");
            pop.AddComfirmEvent(this.Close);
        }
        else
        {
            UIManager.Instance.InfoPopup(UIPopup.Level.Error, "ע��ʧ��");
        }
    }

    public override void Close()
    {
        base.Close();
        UIManager.Instance.Open<UILogin>();
    }

    /// <summary>
    /// ��������Ƿ�Ϊ��
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
