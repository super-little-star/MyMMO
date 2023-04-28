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
    private Transform UIRoot;

    public void Init()
    {
        this.UIRoot = this.transform.GetChild(0);

        // Popup
        this.UIChackIn<UIWaitPopup>("UI/Popup/UIWaitPopup", true);
        this.UIChackIn<UIInfoPopup>("UI/Popup/UIInfoPopup", true);
        this.UIChackIn<UIComfirmPopup>("UI/Popup/UIComfirmPopup", true);

        // Window
        this.UIChackIn<UILogin>("UI/Window/UILogin", false);
        this.UIChackIn<UIRegister>("UI/Window/UIRegister", false);

        this.UIChackIn<UISelectCharacter>("UI/SelectCharacter/UISelectCharacter", true);
        this.UIChackIn<UICreateCharacter>("UI/SelectCharacter/UICreatCharacter", true);

        this.UIChackIn<UIBackground>("UI/UIBackground", false);
        this.UIChackIn<UILoading>("UI/UILoading", false);

    }

    /// <summary>
    /// 登记UI资源
    /// </summary>
    /// <typeparam name="T"></typeparam>
    /// <param name="resoutces">预制体位置</param>
    /// <param name="isReusable">是否可以复用</param>
    private void UIChackIn<T>(string resoutces, bool isReusable) where T : UIBase
    {
        this.UIPreforms.Add(typeof(T), new UIData() { Resources = resoutces, IsReusable = isReusable });
    }

    /// <summary>
    /// 打开UI
    /// </summary>
    /// <typeparam name="T"></typeparam>
    /// <param name="useAnimation">是否使用动画</param>
    /// <returns></returns>
    public T Open<T>(bool useAnimation = true,Transform root = null) where T : UIBase
    {
        Type type = typeof(T);
        UIData data;
        if (this.UIPreforms.TryGetValue(type, out data))
        {
            if (data.Instance != null)
            {
                data.Instance.Open(useAnimation);
            }
            else
            {
                UnityEngine.Object prefab = Resources.Load(data.Resources);
                if (prefab == null) { return default; }

                data.Instance = GameObject.Instantiate(prefab, root == null ? this.UIRoot : root).GetComponent<T>();
                data.Instance.Open(useAnimation);
            }
            return (T)data.Instance;
        }
        return default;
    }

    /// <summary>
    /// 关闭UI
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
                data.Instance.Hide();
            }
            else
            {
                data.Instance.Close();
                data.Instance = null;
            }
        }
    }

    /// <summary>
    /// 卸载UI
    /// </summary>
    /// <param name="type"></param>
    public void Kill(Type type)
    {
        UIData data;
        if(this.UIPreforms.TryGetValue(type, out data))
        {
            if (data.Instance == null) return;
            data.Instance.Close();
        }
    }

    /// <summary>
    /// 把所有打开的UI全部卸载掉
    /// </summary>
    public void KillAll()
    {
        for (int i = 0; i < UIRoot.childCount; i++)
        {
            Transform child = UIRoot.GetChild(i);
            UIBase ui = child.GetComponent<UIBase>();
            if (ui != null)  ui.Close();
        }
    }

    #region Popup
    /// <summary>
    /// 确认弹窗
    /// </summary>
    /// <param name="level">内容等级</param>
    /// <param name="title">标题</param>
    /// <param name="content">内容</param>
    /// <param name="btnComfirmText">确认按钮的文字</param>
    /// <param name="btnCancelText">取消按钮的文字</param>
    /// <returns></returns>
    public UIComfirmPopup ComfirmPopup(UIPopup.Level level, string title, string content, string btnComfirmText = "确认", string btnCancelText = "取消")
    {
        UIComfirmPopup c = Open<UIComfirmPopup>();
        if (c != default)
        {
            c.SetContent(level, title, content, btnComfirmText, btnCancelText);
        }
        return c;
    }

    /// <summary>
    /// 信息弹窗
    /// </summary>
    /// <param name="level">信息内容等级</param>
    /// <param name="content">内容</param>
    /// <returns></returns>
    public UIInfoPopup InfoPopup(UIPopup.Level level, string content)
    {
        UIInfoPopup val = Open<UIInfoPopup>();
        if (val != default)
        {
            val.Content = content;
        }
        val.SetContentColor(level);
        return val;
    }

    /// <summary>
    /// 等待弹窗
    /// </summary>
    /// <param name="content"></param>
    /// <returns></returns>
    public UIWaitPopup WaitPopup(string content)
    {
        UIWaitPopup val = Open<UIWaitPopup>(false);
        if (val != default)
        {
            val.Content = content;
        }
        return val;
    }
    #endregion

    public UILoading Loading()
    {
        return Open<UILoading>(false,this.transform);
    }

}
