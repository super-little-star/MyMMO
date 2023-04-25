using DG.Tweening;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.Events;

public class SceneManager : MonoSingleton<SceneManager>
{

    public void Init()
    {
        UIManager.Instance.Open<UIBackground>(false);
        UIManager.Instance.Open<UILogin>();
    }

    IEnumerator Load(string sceneName)
    {
        Debug.LogFormat("Load the scene:[{0}]", sceneName);

        UILoading ui = UIManager.Instance.Loading();

        AsyncOperation async = UnityEngine.SceneManagement.SceneManager.LoadSceneAsync(sceneName); ;
        async.allowSceneActivation = true;

        while(!async.isDone)
        {
            ui.SetSliderValue(async.progress);
            yield return null;
        }
        yield return new WaitForSeconds(1f);

        ui.S_Progress.DOValue(1f, 0.3f).OnComplete(() => { UIManager.Instance.Close(typeof(UILoading)); });

    }


}
