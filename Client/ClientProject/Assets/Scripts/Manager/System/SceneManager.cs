using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.Events;

public class SceneManager : MonoSingleton<SceneManager>
{
    public UnityAction OnSceneLoadDone = null;

    public void Init()
    {
        UIManager.Instance.Open<UILogin>();
    }
}
