using DG.Tweening;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public abstract class UIBase : MonoBehaviour
{
    
    public virtual void Open()
    {
        if (transform.gameObject.activeSelf == true) return;
        this.transform.gameObject.SetActive(true);
        this.transform.localScale = new Vector3(0.6f, 0.6f, 0.6f);
        this.transform.DOScale(1f, 0.2f).SetEase(Ease.OutElastic);
    }
    public virtual void Close()
    {
        GameObject.Destroy(this.gameObject);
    }

    public virtual void Hide()
    {
        this.gameObject.SetActive(false);
    }
}
