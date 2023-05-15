using ProtoMessage;
using System.Collections;
using System.Collections.Generic;
using Unity.VisualScripting;
using UnityEngine;
using UnityEngine.Events;
using UnityEngine.UI;

public class UISelectCharacter : UIBase
{


    private List<UICharacterInfoItem> characterInfoItems = new List<UICharacterInfoItem>();

    private string CharacterInfoItemPath = "UI/SelectCharacter/UICharacterInfoItem";

    public UIAddCharacterItem AddCharacterItem;

    public Transform CharacterItemRoot;

    public UnityAction<int> OnCurrClassChange;

    private PCharacter currCharacter = null;

    public Button Btn_Enter;

    public override void Open(bool useAnimation)
    {
        base.Open(useAnimation);
        Btn_Enter.onClick.RemoveAllListeners();
        Btn_Enter.onClick.AddListener(OnClickEnter);
        ResetUI();
        User.Instance.OnUserInfoChange += ResetUI;
    }

    public override void Close()
    {
        base.Close();
        User.Instance.OnUserInfoChange -= ResetUI;
    }

    public override void Hide()
    {
        base.Hide();
        User.Instance.OnUserInfoChange -= ResetUI;
    }

    private void ResetUI()
    {
        //删除原有元素
        foreach(var c in characterInfoItems)
        {
            GameObject.Destroy(c.gameObject);
        }
        characterInfoItems.Clear();

        if (CharacterItemRoot == null) return;
        //添加新元素
        foreach(PCharacter info in User.Instance.Info.Characters)
        {
            UnityEngine.Object res = Resources.Load(CharacterInfoItemPath);
            ToggleGroup tg = CharacterItemRoot.GetComponent<ToggleGroup>();
            if(res != null)
            {
                UICharacterInfoItem ui = GameObject.Instantiate(res, CharacterItemRoot).GetComponent<UICharacterInfoItem>();
                ui.toggle.group = tg;
                ui.Info = info;
                ui.toggle.onValueChanged.AddListener((bool t) =>
                {
                    if (t)
                    {
                        this.currCharacter = info;
                        OnCurrClassChange?.Invoke((int)info.Class);
                    }
                });
                characterInfoItems.Add(ui);
            }
        }
        //把添加元素Item放到最后
        if (AddCharacterItem != null) { AddCharacterItem.transform.SetAsLastSibling(); }
    }

    private void OnClickEnter()
    {

    }
}
