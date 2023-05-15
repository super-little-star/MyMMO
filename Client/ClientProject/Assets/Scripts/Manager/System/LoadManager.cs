using DG.Tweening;
using Network;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class LoadManager : MonoBehaviour
{
    IEnumerator Start()
    {
        Logger.Init();
        DOTween.Init();
        UIManager.Instance.Init();
        SceneManager.Instance.Init();

        DataManager.Instance.Init();
        
        Debug.Log("Connecting to server...");
        NetSerice.Instance.Init("127.0.0.1", 7788);
        UserSerice.Instance.Init();
        


        yield return null;
    }
    
}
