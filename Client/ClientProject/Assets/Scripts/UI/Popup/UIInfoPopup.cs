using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.Events;
using UnityEngine.UI;

public class UIInfoPopup : UIPopup
{
    public Button Btn_Comfirm;

    private void Start()
    {
        if (Btn_Comfirm != null) Btn_Comfirm.onClick.AddListener(this.Hide);
    }

    public void AddComfirmEvent(UnityAction aciton)
    {
        if (Btn_Comfirm != null) Btn_Comfirm.onClick.AddListener(aciton);
    }

    public override void Close()
    {
        base.Close();
        if (Btn_Comfirm != null) Btn_Comfirm.onClick.RemoveAllListeners();
    }

    public override void Hide()
    {
        base.Hide();
        if (Btn_Comfirm != null) Btn_Comfirm.onClick.RemoveAllListeners();
    }

}
