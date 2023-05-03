using ProtoMessage;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.Events;
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

    public Button Btn_Delect;

    public Toggle toggle;

    private void OnInfoChange()
    {
        if (IsNull()) return;

        this.Icon.sprite = this.ClassSprites[(int)this.info.Class];
        this.T_Name.text = this.info.Name;
        this.T_Info.text = string.Format("{0} | Lv.{1}", DataManager.Instance.Characters[(int)this.info.Class].Name, info.Level);

        // TODO 添加按钮事件
    }

    private bool IsNull()
    {
        bool isNull = Icon == null || T_Name == null || T_Info == null || Btn_Delect == null;
        if (isNull)
        {
            Debug.LogWarning("UICharacterInfoItem something is null");
        }
        return isNull;
    }


}
