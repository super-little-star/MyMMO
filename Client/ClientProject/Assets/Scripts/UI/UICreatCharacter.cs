using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.Events;
using UnityEngine.UI;

public class UICreatCharacter : UIBase
{
    public InputField IF_UserName;

    public ToggleGroup TG_Class;

    public RawImage RI_Characterimg;

    public Button Btn_Creat;
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

    public void Init()
    {
        this.CurrCharacterClass = 0;
    }

    private void OnClassChange()
    {
        CharacterData data = DataManager.Instance.Characters[this.CurrCharacterClass];

        if(T_CurrClassName !=null ) { T_CurrClassName.text = data.Name; }
        if(T_CurrClassDes !=null) { T_CurrClassDes.text = data.Description; }   


        OnCurrClassChange?.Invoke(this.CurrCharacterClass);
    }
}
