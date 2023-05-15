using DG.Tweening;
using System.Collections;
using System.Collections.Generic;
using Unity.VisualScripting;
using UnityEngine;
using UnityEngine.UI;

public class UIWindow : UIBase
{
    public Button Btn_X;

    public override void Open(bool useAnimation)
    {
        base.Open(useAnimation);
        if (Btn_X != null) Btn_X.onClick.AddListener(this.Close);
    }

    public override void Close()
    {
        base.Close();
        if (Btn_X != null) Btn_X.onClick.RemoveAllListeners();
    }

    public override void Hide()
    {
        base.Hide();
        if (Btn_X != null) Btn_X.onClick.RemoveAllListeners();
    }
}
