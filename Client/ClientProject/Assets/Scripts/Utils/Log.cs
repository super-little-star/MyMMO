using log4net;
using System.Collections;
using System.Collections.Generic;
using System.IO;
using UnityEngine;

public static class Logger 
{
    private static ILog log = LogManager.GetLogger("UnityLogger");

    public static void Init()
    {
        Application.logMessageReceived += OnLogMessageReceived;

        Debug.Log("log4net Init ....");
    }

    public static void OnLogMessageReceived(string condition, string stackTrace, LogType type)
    {
        switch (type)
        {
            case LogType.Error:
                log.ErrorFormat("{0}\r\n{1}", condition, stackTrace.Replace("\n", "\r\n"));
                break;
            case LogType.Assert:
                log.DebugFormat("{0}\r\n{1}", condition, stackTrace.Replace("\n", "\r\n"));
                break;
            case LogType.Exception:
                log.FatalFormat("{0}\r\n{1}", condition, stackTrace.Replace("\n", "\r\n"));
                break;
            case LogType.Warning:
                log.WarnFormat("{0}\r\n{1}", condition, stackTrace.Replace("\n", "\r\n"));
                break;
            default:
                log.Info(condition);
                break;
        }
    }
}
