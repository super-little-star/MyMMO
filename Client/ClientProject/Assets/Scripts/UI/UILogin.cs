using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;

public class UILogin : MonoBehaviour
{
    public InputField IF_UserName;
    public InputField IF_Password;

    public Button Btn_Login;
    public Button Btn_Register;

    private void Start()
    {
        
    }

    public void ResetUI()
    {
        IF_UserName.text = string.Empty;
        IF_Password.text = string.Empty;
    }
}
