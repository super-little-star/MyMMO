package Model

type DbUser struct {
	UID          int64
	UserName     string
	Password     string
	Characters   []*DbCharacter
	RegisterTime int64
}
