﻿using System.Collections;
using UnityEngine;
public abstract class MonoSingleton<T> : MonoBehaviour where T : MonoBehaviour
{
    public bool global = true;
    static T instance;

    public static T Instance
    {
        get
        {
            if (instance == null)
            {
                instance = (T)FindObjectOfType<T>();
            }
            return instance;
        }
    }

    private void Awake()
    {
        Debug.LogWarningFormat("{0}[{1}] Awake", typeof(T), this.GetInstanceID());
        if (global)
        {
            if (instance != null && instance != this.gameObject.GetComponent<T>())
            {
                Destroy(this.gameObject);
                return;
            }
            DontDestroyOnLoad(this.gameObject);
            instance = this.gameObject.GetComponent<T>();
        }
        
    }

    private void Start()
    {
        this.OnStart();
    }

    protected virtual void OnStart()
    {

    }

    
}
