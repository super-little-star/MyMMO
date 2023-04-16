using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System;
using Unity.VisualScripting;

public class UIManager : MonoSingleton<UIManager>
{

    class UIData
    {
        public string Resources;
        public bool IsReusable;
        public UIBase Instance;
    }

    private Dictionary<Type, UIData> UIPreforms = new Dictionary<Type, UIData>();


    public void Init()
    {
        this.UIChackIn<UIWaitPopup>("UI/Popup/UIWaitPopup", true);
        this.UIChackIn<UIInfoPopup>("UI/Popup/UIInfoPopup", true);
        this.UIChackIn<UIComfirmPopup>("UI/Popup/UIComfirmPopup", true);

        this.UIChackIn<UILogin>("UI/Window/UILogin", false);
        this.UIChackIn<UIRegister>("UI/Window/UIRegister", false);
    }

    /// <summary>
    /// �Ǽ�UI��Դ
    /// </summary>
    /// <typeparam name="T"></typeparam>
    /// <param name="resoutces">Ԥ����λ��</param>
    /// <param name="isReusable">�Ƿ���Ը���</param>
    private void UIChackIn<T>(string resoutces, bool isReusable) where T : UIBase
    {
        this.UIPreforms.Add(typeof(T), new UIData() { Resources = resoutces, IsReusable = isReusable });
    }

    /// <summary>
    /// ��UI
    /// </summary>
    /// <typeparam name="T"></typeparam>
    /// <returns></returns>
    public T Open<T>() where T : UIBase
    {
        Type type = typeof(T);
        UIData data;
        if (this.UIPreforms.TryGetValue(type, out data))
        {
            if (data.Instance != null)
            {
                data.Instance.Open();
            }
            else
            {
                UnityEngine.Object prefab = Resources.Load(data.Resources);
                if (prefab == null) { return default; }

                data.Instance = GameObject.Instantiate(prefab,this.transform).GetComponent<T>();
                data.Instance.Open();
            }
            return (T)data.Instance;
        }
        return default;
    }

    /// <summary>
    /// �ر�UI
    /// </summary>
    /// <param name="type"></param>
    public void Close(Type type) 
    {
        UIData data;
        if(this.UIPreforms.TryGetValue(type,out data))
        {
            if (data.Instance == null) return;
            if(data.IsReusable)
            {
                data.Instance.gameObject.SetActive(false);
            }
            else
            {
                GameObject.Destroy(data.Instance.gameObject);
                data.Instance = null;
            }
        }
    }

    /// <summary>
    /// ȷ�ϵ���
    /// </summary>
    /// <typeparam name="T"></typeparam>
    /// <param name="level">���ݵȼ�</param>
    /// <param name="title">����</param>
    /// <param name="content">����</param>
    /// <param name="btnComfirmText">ȷ�ϰ�ť������</param>
    /// <param name="btnCancelText">ȡ����ť������</param>
    /// <returns></returns>
    public T Popup<T>(UIPopup.Level level, string title, string content, string btnComfirmText = "ȷ��", string btnCancelText = "ȡ��") where T : UIComfirmPopup
    {
        T val = Open<T>();
        if(val != default)
        {
            val.SetContent(level, title, content, btnComfirmText, btnCancelText);
        }
        return val;
    }

    /// <summary>
    /// ��Ϣ����
    /// </summary>
    /// <typeparam name="T"></typeparam>
    /// <param name="level">��Ϣ���ݵȼ�</param>
    /// <param name="content">����</param>
    /// <returns></returns>
    public T Popup<T>(UIPopup.Level level,string content) where T : UIInfoPopup
    {
        T val = Open<T>();
        if(val != default)
        {
            val.content = content;
        }
        val.SetContentColor(level);
        return val;
    }

    /// <summary>
    /// �ȴ�����
    /// </summary>
    /// <typeparam name="T"></typeparam>
    /// <param name="content"></param>
    /// <returns></returns>
    public T Popup<T>(string content) where T : UIWaitPopup
    {
        T val =  Open<T>();
        if (val != default)
        {
            val.content = content;
        }
        return val;
    }
}
