using DG.Tweening;
using System.Collections;
using System.Collections.Generic;
using Unity.VisualScripting;
using UnityEngine;
using UnityEngine.UI;

public class UIWindow : UIBase
{
    public Button Btn_X;

    private void Start()
    {
        this.OnStart();
    }

    protected virtual void OnStart()
    {
        if (Btn_X != null) Btn_X.onClick.AddListener(this.Close);
    }

}
