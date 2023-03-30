using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class LoadManager : MonoBehaviour
{
    IEnumerator Start()
    {
        UserServer.Instance.Init();
        UserServer.Instance.ConnectToServer();

        yield return null;
    }
}
