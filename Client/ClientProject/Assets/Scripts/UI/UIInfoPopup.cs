using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

public class UIInfoPopup : UIPopup
{
    public Button Btn_Comfirm;

    protected override void OnStart()
    {
        base.OnStart();
        if (Btn_Comfirm != null) Btn_Comfirm.onClick.AddListener(OnComfirmClick);
    }

    public override void Close()
    {
        base.Close();
        if (Btn_Comfirm != null) Btn_Comfirm.onClick.RemoveAllListeners();
    }

    private void OnComfirmClick()
    {
        UIManager.Instance.Close(typeof(UIInfoPopup));
    }
}
