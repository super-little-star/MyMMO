using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System.IO;
using Newtonsoft.Json;

public class DataManager : Singleton<DataManager>
{
    public string DataPath = @"Assets\Resources\Data\";

    public Dictionary<int,CharacterData> Characters = new();


    public void Init()
    {
        Debug.LogFormat("DataManger:: Init");

        string json = File.ReadAllText(DataPath+"CharacterData.json");

        this.Characters = JsonConvert.DeserializeObject<Dictionary<int,CharacterData>>(json);

    }
}
