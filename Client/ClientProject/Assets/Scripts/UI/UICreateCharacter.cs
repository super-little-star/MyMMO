using System.Collections;
using System.Collections.Generic;
using System.Text.RegularExpressions;
using UnityEngine;
using UnityEngine.Events;
using UnityEngine.UI;

public class UICreateCharacter : UIBase
{
    //字母中文开头，允许5-16字节，允许中文英文数字
    public string CharacterNameRegular = @"^[\u4e00-\u9fa5_a-zA-Z][\u4e00-\u9fa5_-zA-Z0-9]{4,16}$";

    public InputField IF_UserName;

    public ToggleGroup TG_Class;

    public RawImage RI_Characterimg;

    public Button Btn_Create;
    public Button Btn_Break;

    public Text T_CurrClassName;
    public Text T_CurrClassDes;

    public UnityAction<int> OnCurrClassChange;

    private int currCharacterClass;
    public int CurrCharacterClass
    {
        get { return currCharacterClass; }
        set
        {
            this.currCharacterClass = value;
            OnClassChange();
        }
    }

    public override void Open(bool useAnimation)
    {
        base.Open(useAnimation);
        this.CurrCharacterClass = 0;
    }

    private void OnClassChange()
    {
        CharacterData data = DataManager.Instance.Characters[this.CurrCharacterClass];

        if(T_CurrClassName !=null ) { T_CurrClassName.text = data.Name; }
        if(T_CurrClassDes !=null) { T_CurrClassDes.text = data.Description; }   


        OnCurrClassChange?.Invoke(this.CurrCharacterClass);
    }

    private void OnClickeCreate()
    {
        if (IF_UserName == null) return;

        if(IF_UserName.text == string.Empty)
        {
            UIInfoPopup pop = UIManager.Instance.InfoPopup(UIPopup.Level.Error, "请输入角色名");
            pop.AddComfirmEvent(() => { IF_UserName.ActivateInputField(); });
            return;
        }

        if (!Regex.IsMatch(IF_UserName.text, CharacterNameRegular))
        {
            UIInfoPopup pop = UIManager.Instance.InfoPopup(UIPopup.Level.Error, "角色名存在非法输入");
            pop.AddComfirmEvent(() => { IF_UserName.ActivateInputField(); IF_UserName.text = string.Empty; });
            return;
        }

        UserSerice.Instance.SendCreateCharacter(currCharacterClass,IF_UserName.text);
    }
}
