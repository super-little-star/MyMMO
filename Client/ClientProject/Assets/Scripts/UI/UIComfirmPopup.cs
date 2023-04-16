using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

public class UIComfirmPopup : UIPopup
{
    public Text T_Title;

    public Button Btn_Comfirm;
    public Button Btn_Cancel;

    

    protected override void OnStart()
    {
        base.OnStart();
        if (Btn_Comfirm != null) Btn_Comfirm.onClick.AddListener(OnComfirmClick);
        if(Btn_Cancel!=null)Btn_Cancel.onClick.AddListener(OnCancelClick);
    }

    public void SetContent(Level level,string title,string content,string btnComfirmText,string btnCancelText)
    {
        switch (level)
        {
            case Level.Error:T_Title.color = Error;break;
            default:T_Title.color = Normal; break;
        }
        T_Title.text = title;
        T_Content.text = content;

        Btn_Comfirm.GetComponentInChildren<Text>().text = btnComfirmText;
        Btn_Cancel.GetComponentInChildren<Text>().text = btnCancelText;
    }

    public override void Close()
    {
        base.Close();
        if(Btn_Comfirm!=null) Btn_Comfirm.onClick.RemoveAllListeners();
        if(Btn_Cancel!=null) Btn_Cancel.onClick.RemoveAllListeners();
    }

    private void OnComfirmClick()
    {
        UIManager.Instance.Close(typeof(UIComfirmPopup));
    }

    private void OnCancelClick()
    {
        UIManager.Instance.Close(typeof(UIComfirmPopup));
    }
}
