package manager

import (
	"errors"
	"mmo_server/DB/DbObject"
	"mmo_server/object"
)

var (
	ErrCharacterExist      = errors.New("character mgr:: character is exist")
	ErrCharacterIsNotFound = errors.New("character mgr:: character is not found")
)

var CharacterManager ICharacterManager

type ICharacterManager interface {
	Add(db *DbObject.DbCharacter) (*object.Character, error)
	Remove(characterId int32)
	Get(characterId int32) (*object.Character, error)
}

type GCharacterManager struct {
	Characters map[int32]*object.Character
}

// InitCharacterManager
//
//	@Description: 初始化CharacterManager
func InitCharacterManager() {
	CharacterManager = &GCharacterManager{
		Characters: make(map[int32]*object.Character),
	}
}

// Add
//
//	@Description: 添加角色
//	@receiver g
//	@param db
//	@return error
func (g *GCharacterManager) Add(db *DbObject.DbCharacter) (*object.Character, error) {
	if _, ok := g.Characters[db.ID]; ok {
		return nil, ErrCharacterExist
	}

	character := object.NewCharacter(db)

	g.Characters[character.Db.ID] = character
	return character, nil
}

// Remove
//
//	@Description: 移出角色
//	@receiver g
//	@param characterId
func (g *GCharacterManager) Remove(characterId int32) {
	delete(g.Characters, characterId)
}

// Get
//
//	@Description: 查询角色
//	@receiver g
//	@param characterId
//	@return *object.Character
//	@return error
func (g *GCharacterManager) Get(characterId int32) (*object.Character, error) {
	if res, ok := g.Characters[characterId]; ok {
		return res, nil
	} else {
		return nil, ErrCharacterIsNotFound
	}
}
