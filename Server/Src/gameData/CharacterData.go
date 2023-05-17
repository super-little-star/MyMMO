package gameData

type CharacterClass int

const (
	Class_Warrior CharacterClass = iota
	Class_Hunter
	Class_Sorcerer
)

type CharacterData struct {
	ID        int            // 表格ID
	Class     CharacterClass // 职业
	Name      string         // 名字
	BaseSpeed int            // 基础速度

	BaseHp  int // 基础生命
	BaseMp  int // 基础蓝量
	BaseSTR int // 基础力量
	BaseINT int // 基础智力
	BaseBRA int // 基础体力
	BaseDEX int // 基础敏捷

	GrowthSTR float32 // 成长力量
	GrowthINT float32 // 成长智力
	GrowthBRA float32 // 成长体力
	GrowthDEX float32 // 成长敏捷

	BaseCrit float32 // 基础暴击率
}
