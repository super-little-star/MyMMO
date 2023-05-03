using ProtoMessage;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.Events;

public class User : Singleton<User>
{
    private PUser info;

    public PUser Info { get { return info; } }

    public UnityAction OnUserInfoChange;

    public void SetUser(PUser user)
    {
        this.info = user;
        this.OnUserInfoChange?.Invoke();
    }

    public void SetCharacters(List<PCharacter> characters)
    {
        this.info.Characters.Clear();
        this.info.Characters.AddRange(characters);
        this.OnUserInfoChange?.Invoke();
    }
}
