using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public enum CharacterClass
{
    Warrior = 0,
    Hunter = 1,
    Sorcere = 2,
}

public class CharacterData
{
    public int ID { get; set; }
    public CharacterClass Class { get; set; }

    /// <summary>
    /// 名字
    /// </summary>
    public string Name { get; set; }

    /// <summary>
    /// 资源位置
    /// </summary>
    public string Resource { get; set; }
    /// <summary>
    /// 描述
    /// </summary>
    public string Description { get; set; }

    /// <summary>
    /// 基础书读
    /// </summary>
    public int BaseSpeed { get; set; }

    /// <summary>
    /// 基础生命
    /// </summary>
    public int BaseHp { get; set; }
    /// <summary>
    /// 基础魔法值
    /// </summary>
    public int BaseMp { get; set; }

    /// <summary>
    /// 基础力量
    /// </summary>
    public int BaseSTR { get; set; }
    /// <summary>
    /// 基础智力
    /// </summary>
    public int BaseINT { get; set; }

    /// <summary>
    /// 基础体力
    /// </summary>
    public int BaseBRA { get; set; }
    /// <summary>
    /// 基础敏捷
    /// </summary>
    public int BaseDEX { get; set; }
    /// <summary>
    /// 成长力量
    /// </summary>
    public float GrowthSTR { get; set; }
    /// <summary>
    /// 成长智力
    /// </summary>
    public float GrowthINT { get;set; }
    /// <summary>
    /// 成长体力
    /// </summary>
    public float GrowthBRA { get;set; }
    /// <summary>
    /// 成长敏捷
    /// </summary>
    public float GrowthDEX { get; set;}
    /// <summary>
    /// 基础暴击率
    /// </summary>
    public float BaseCrit { get; set; }
}
