using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class CameraManager : MonoSingleton<CameraManager>
{
    private Camera main;

    public Camera Main
    {
        get
        {
            return main;
        }
    }

    private void Start()
    {
        this.main = Camera.main;
    }

    public void SetMainPosition(Vector3 pos)
    {
        if (this.main == null) return;

        this.main.transform.position = pos;
    }

}
