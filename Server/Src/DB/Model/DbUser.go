package Model

type DbUser struct {
	UID            int64
	UserName       string
	Password       string
	CharacterCount int8
	Characters     []DbCharacter
	RegisterTime   int64
}
