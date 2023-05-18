package gameData

import (
	"encoding/json"
	"os"
	"reflect"
)

const path = "gameData/"

var DataManager *GDataManager

type GDataManager struct {
	Characters map[int]*CharacterData
	Maps       map[int]*MapData
}

func Init() error {

	DataManager = &GDataManager{}

	DataManager.Characters = make(map[int]*CharacterData)
	j, err := os.ReadFile(path + "CharacterData.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(j, &DataManager.Characters)
	if err != nil {
		return err
	}

	DataManager.Maps, err = UnmarshalJson[MapData]()
	if err != nil {
		return err
	}

	return nil
}

func UnmarshalJson[T any]() (map[int]*T, error) {
	var t T
	key := reflect.TypeOf(t).Name()
	m := make(map[int]*T)
	j, err := os.ReadFile(path + key + ".json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(j, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
