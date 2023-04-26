using ProtoMessage;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class User : Singleton<User>
{
    private PUser info;

    public PUser Info { get { return info; } }

    public void SetUser(PUser user)
    {
        this.info = user;
    }
}
