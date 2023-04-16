using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

public class UIPopup : UIBase
{
    public enum Level
    {
        Normal = 0,
        Warnning = 1,
        Error = 2,
    }

    protected Color Normal = Color.white;
    protected Color Warnning = new(221, 199, 117);
    protected Color Error = new(238, 125, 73);

    /// <summary>
    /// 弹窗上的内容组件
    /// </summary>
    public Text T_Content;
    /// <summary>
    /// 弹窗内容
    /// </summary>
    public string content
    {
        get
        {
            if(T_Content != null) return T_Content.text;
            else return string.Empty;
        }
        set
        {
            if (T_Content != null) T_Content.text = value;
        }
    }

    protected override void OnStart()
    {
        base.OnStart();

    }

    public void SetContentColor(Level level = Level.Normal)
    {
        switch (level)
        {
            case Level.Warnning:
                {
                    T_Content.color = Warnning; 
                    break;
                }
            case Level.Error:
                {
                    T_Content.color = Error;
                    break;
                }
            default:
                {
                    T_Content.color = Normal;
                    break;
                }
        }
    }
}
