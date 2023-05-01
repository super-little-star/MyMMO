using ProtoMessage;
using System.Collections;
using System.Collections.Generic;
using Unity.VisualScripting;
using UnityEngine;
using UnityEngine.Events;

public class UISelectCharacter : UIBase
{
    private List<PCharacter> charactersInfo;

    public List<PCharacter> CharactersInfo
    {
        get { return charactersInfo; }
        set
        {
            charactersInfo = value;
            OnCharactersInfoChange();
        }
    }

    private string CharacterInfoItemPath = "UI/SelectCharacter/UICharacterInfoItem";

    public UIAddCharacterItem AddCharacterItem;

    public Transform CharacterItemRoot;

    public UnityAction<int> OnCurrClassChange;



    private void OnCharactersInfoChange()
    {
        if(AddCharacterItem != null) { AddCharacterItem.transform.SetAsLastSibling(); }
        foreach(PCharacter info in CharactersInfo)
        {
            UnityEngine.Object res = Resources.Load(CharacterInfoItemPath);
            if(res != null)
            {
                UICharacterInfoItem ui = GameObject.Instantiate(res, CharacterItemRoot).GetComponent<UICharacterInfoItem>();
                // TODO
            }
        }
    }
}
