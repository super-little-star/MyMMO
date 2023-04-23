using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.Events;

public class SceneManager : MonoSingleton<SceneManager>
{
    public UnityAction<float> OnLoading = null;
    public UnityAction OnSceneLoadDone = null;

    public void Init()
    {
        UIManager.Instance.Open<UILogin>();
    }

    IEnumerator Load(string sceneName)
    {
        Debug.LogFormat("Load the scene:[{0}]", sceneName);
        AsyncOperation async = UnityEngine.SceneManagement.SceneManager.LoadSceneAsync(sceneName); ;
        async.allowSceneActivation = true;
        async.completed += SceneLoadCompleted;

        while(!async.isDone)
        {
            OnLoading?.Invoke(async.progress);
            yield return null;
        }
    }

    private void SceneLoadCompleted(AsyncOperation async)
    {
        OnLoading?.Invoke(1f);

        OnSceneLoadDone?.Invoke();

        Debug.Log("Load Scene Completed...");
    }


    private void LoadCreateCharacter()
    {

    }
}
