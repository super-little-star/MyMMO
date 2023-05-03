using System.Collections;
using System.Collections.Generic;
using System.Text.RegularExpressions;
using UnityEngine;
using UnityEngine.Events;
using UnityEngine.UI;

public class UICreateCharacter : UIBase
{
    //��ĸ���Ŀ�ͷ������5-16�ֽڣ���������Ӣ������
    private string CharacterNameRegular = @"^[\u4e00-\u9fa5_a-zA-Z][\u4e00-\u9fa5_-zA-Z0-9]{4,16}$";

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
        Btn_Create.onClick.AddListener(OnClickeCreate);
        Btn_Break.onClick.AddListener(OnClickBreak);
    }

    public override void Hide()
    {
        base.Hide();
        Btn_Create.onClick.RemoveAllListeners();
        Btn_Break.onClick.RemoveAllListeners();
    }

    public override void Close()
    {
        base.Close();
        Btn_Create.onClick.RemoveAllListeners();
        Btn_Break.onClick.RemoveAllListeners();
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
            UIInfoPopup pop = UIManager.Instance.InfoPopup(UIPopup.Level.Error, "�������ɫ��");
            pop.AddComfirmEvent(() => { IF_UserName.ActivateInputField(); });
            return;
        }

        if (!Regex.IsMatch(IF_UserName.text, CharacterNameRegular))
        {
            UIInfoPopup pop = UIManager.Instance.InfoPopup(UIPopup.Level.Error, "��ɫ�����ڷǷ�����");
            pop.AddComfirmEvent(() => { IF_UserName.ActivateInputField(); IF_UserName.text = string.Empty; });
            return;
        }

        UserSerice.Instance.SendCreateCharacter(currCharacterClass,IF_UserName.text);
    }

    private void OnClickBreak()
    {
        UIManager.Instance.Close(typeof(UICreateCharacter));
        UIManager.Instance.Open<UISelectCharacter>(false);
    }
}