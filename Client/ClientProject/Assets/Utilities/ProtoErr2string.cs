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
            case ProtoMessage.Error.UserNameExist:return "用户名已存在";
            case ProtoMessage.Error.UserNotExist:return "用户不存在";
            case ProtoMessage.Error.PasswordNotMatch:return "密码不正确";
            case ProtoMessage.Error.UserIsOnline:return "用户正在游戏中";
            default:
                return string.Empty;
        }
    }
}
