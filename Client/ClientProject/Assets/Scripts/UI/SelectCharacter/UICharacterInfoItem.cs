using ProtoMessage;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

public class UICharacterInfoItem : MonoBehaviour
{
    private PCharacter info;

    public PCharacter Info
    {
        set
        {
            this.info = value;
            this.OnInfoChange();
        }
        get
        {
            return this.info;
        }
    }

    public Sprite[] ClassSprites;

    public Image Icon;

    public Text T_Name;
    public Text T_Info;

    public Button Btn_Enter;
    public Button Btn_Delect;

    private void OnInfoChange()
    {
        if (IsNull())
        {
            Debug.LogWarning("UICharacterInfoItem something is null");
            return;
        }
        this.Icon.sprite = this.ClassSprites[(int)this.info.Class];
        this.T_Name.text = this.info.Name;
        this.T_Info.text = string.Format("{0} | {1}", info.Class.ToString(), info.Level);
    }

    private bool IsNull()
    {
        return Icon == null || T_Name == null || T_Info == null || Btn_Enter == null || Btn_Delect == null;
    }
}
