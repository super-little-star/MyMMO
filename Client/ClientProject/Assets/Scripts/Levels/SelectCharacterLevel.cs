using System.Collections;
using System.Collections.Generic;
using Unity.VisualScripting;
using UnityEngine;

public class SelectCharacterLevel : MonoBehaviour
{
    public Transform cameraPos;

    public GameObject[] CharacterShow;

    public RenderTexture CameraTexture;

    // Start is called before the first frame update
    void Start()
    {
        if (cameraPos == null) return;
        CameraManager.Instance.Main.transform.position = this.cameraPos.position;
        CameraManager.Instance.Main.transform.rotation = this.cameraPos.rotation;
        CameraManager.Instance.Main.targetTexture = CameraTexture;
        CameraManager.Instance.Main.clearFlags = CameraClearFlags.SolidColor;

        UICreateCharacter create = UIManager.Instance.GetUI<UICreateCharacter>();
        UISelectCharacter select = UIManager.Instance.GetUI<UISelectCharacter>();

        create.OnCurrClassChange += OnCurrClassChange;
        select.OnCurrClassChange += OnCurrClassChange;
    }


    void OnCurrClassChange(int characterClass)
    {
        for(int i = 0; i<CharacterShow.Length; i++)
        {
            CharacterShow[i].SetActive(false);
        }

        if (characterClass == -1) return;

        CharacterShow[characterClass].SetActive(true);
    }
    
}
