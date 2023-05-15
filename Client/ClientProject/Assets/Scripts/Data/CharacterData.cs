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
    /// ����
    /// </summary>
    public string Name { get; set; }

    /// <summary>
    /// ��Դλ��
    /// </summary>
    public string Resource { get; set; }
    /// <summary>
    /// ����
    /// </summary>
    public string Description { get; set; }

    /// <summary>
    /// �������
    /// </summary>
    public int BaseSpeed { get; set; }

    /// <summary>
    /// ��������
    /// </summary>
    public int BaseHp { get; set; }
    /// <summary>
    /// ����ħ��ֵ
    /// </summary>
    public int BaseMp { get; set; }

    /// <summary>
    /// ��������
    /// </summary>
    public int BaseSTR { get; set; }
    /// <summary>
    /// ��������
    /// </summary>
    public int BaseINT { get; set; }

    /// <summary>
    /// ��������
    /// </summary>
    public int BaseBRA { get; set; }
    /// <summary>
    /// ��������
    /// </summary>
    public int BaseDEX { get; set; }
    /// <summary>
    /// �ɳ�����
    /// </summary>
    public float GrowthSTR { get; set; }
    /// <summary>
    /// �ɳ�����
    /// </summary>
    public float GrowthINT { get;set; }
    /// <summary>
    /// �ɳ�����
    /// </summary>
    public float GrowthBRA { get;set; }
    /// <summary>
    /// �ɳ�����
    /// </summary>
    public float GrowthDEX { get; set;}
    /// <summary>
    /// ����������
    /// </summary>
    public float BaseCrit { get; set; }
}
