using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.Events;
using UnityEngine.UI;

public class UIComfirmPopup : UIPopup
{
    public Text T_Title;

    public Button Btn_Comfirm;
    public Button Btn_Cancel;

    public override void Open()
    {
        base.Open();
        if (Btn_Comfirm != null) Btn_Comfirm.onClick.AddListener(this.Hide);
        if (Btn_Cancel != null) Btn_Cancel.onClick.AddListener(this.Hide);
    }

    /// <summary>
    /// …Ë÷√µØ¥∞µƒƒ⁄»›
    /// </summary>
    /// <param name="level"></param>
    /// <param name="title"></param>
    /// <param name="content"></param>
    /// <param name="btnComfirmText"></param>
    /// <param name="btnCancelText"></param>
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

    public void AddComfirmEvent(UnityAction action)
    {
        if(Btn_Comfirm != null)
        {
            Btn_Comfirm.onClick.AddListener(action);
        }
    }

    public void AddCancelEvent(UnityAction action)
    {
        if(action != null)
        {
            Btn_Cancel.onClick.AddListener(action);
        }
    }

    public override void Close()
    {
        base.Close();
        if(Btn_Comfirm!=null) Btn_Comfirm.onClick.RemoveAllListeners();
        if(Btn_Cancel!=null) Btn_Cancel.onClick.RemoveAllListeners();
    }

    public override void Hide()
    {
        base.Hide();
        if (Btn_Comfirm != null) Btn_Comfirm.onClick.RemoveAllListeners();
        if (Btn_Cancel != null) Btn_Cancel.onClick.RemoveAllListeners();
    }

}
