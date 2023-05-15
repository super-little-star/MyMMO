using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

public class UIAddCharacterItem : MonoBehaviour
{
    public Button Btn_Add;

    private void Start()
    {
        if(Btn_Add != null)
        {
            Btn_Add.onClick.AddListener(() =>
            {
                UIManager.Instance.Close(typeof(UISelectCharacter));
                UIManager.Instance.Open<UICreateCharacter>(false);
            });
        }
    }
}
