using DG.Tweening;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class UIWaitPopup : UIPopup
{

    public override void Open()
    {
        this.gameObject.SetActive(true);
    }

}
