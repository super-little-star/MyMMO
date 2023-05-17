package gameData

type MapType int

const (
	MapType_City = iota
	MapType_Field
)

type MapData struct {
	ID   int32
	Name string
	Type MapType
}
