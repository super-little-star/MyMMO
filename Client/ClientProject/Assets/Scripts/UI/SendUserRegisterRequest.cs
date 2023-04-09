using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class SendUserRegisterRequest : MonoBehaviour
{
    // Start is called before the first frame update
    public void SendRequest()
    {
        string name = "wzs";
        string psw = "112233";
        UserSerice.Instance.SendUserRegister(name, psw);
    }
}
