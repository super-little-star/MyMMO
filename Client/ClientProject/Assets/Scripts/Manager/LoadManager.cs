using Network;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class LoadManager : MonoBehaviour
{
    IEnumerator Start()
    {
        Debug.Log("Connecting to server...");
        NetClient.Instance.Init("127.0.0.1", 7788);
        UserSerice.Instance.Init();

        yield return null;
    }
    
}
