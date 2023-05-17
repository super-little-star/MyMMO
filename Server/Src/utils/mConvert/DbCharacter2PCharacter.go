package mConvert

import (
	"mmo_server/DB/DbObject"
	"mmo_server/ProtoMessage"
)

// DbCharacters2PCharacters
//
//	@Description: 将数据库Character切片转换成Protobuf的Character切片
//	@param dbCharacter
//	@return []*ProtoMessage.PCharacter
func DbCharacters2PCharacters(dbCharacter []*DbObject.DbCharacter) []*ProtoMessage.PCharacter {
	var pCharacters []*ProtoMessage.PCharacter
	for _, c := range dbCharacter {
		pc := DbCharacter2PCharacter(c)
		pCharacters = append(pCharacters, pc)
	}

	return pCharacters
}

// DbCharacter2PCharacter
//
//	@Description:将DB的Character对象转换成Protobuf的Character对象
//	@param dbCharacter
//	@return *ProtoMessage.PCharacter
func DbCharacter2PCharacter(dbCharacter *DbObject.DbCharacter) *ProtoMessage.PCharacter {
	if dbCharacter == nil {
		return nil
	}
	return &ProtoMessage.PCharacter{
		Id:    dbCharacter.ID,
		Name:  dbCharacter.Name,
		Class: ProtoMessage.CharacterClass(dbCharacter.Class),
		Level: dbCharacter.Level,
	}
}
