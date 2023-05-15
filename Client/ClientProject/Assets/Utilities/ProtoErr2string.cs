using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

public class ProtoErr2string
{
    public static string Convert(ProtoMessage.Error err)
    {
        switch (err)
        {
            case ProtoMessage.Error.RegisterUserNameExist:return "用户名已存在";
            case ProtoMessage.Error.LoginUserNotExist:return "用户不存在";
            case ProtoMessage.Error.LoginPasswordNotMatch:return "密码不正确";
            case ProtoMessage.Error.LoginUserIsOnline:return "用户正在游戏中";
            case ProtoMessage.Error.CreateCharacterNameExist:return "角色名已存在";
            default:
                return string.Empty;
        }
    }
}
