using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class UIManager : Singleton<UIManager>
{

    class UIObj
    {
        public string Resources;
        public bool IsReusable;
        public GameObject Instance;
    }
}
