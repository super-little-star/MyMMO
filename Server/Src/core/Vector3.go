package core

var (
	Zero    = NewVector3(0, 0, 0)
	One     = NewVector3(1, 1, 1)
	Up      = NewVector3(0, 1, 0)
	Down    = NewVector3(0, -1, 0)
	Left    = NewVector3(-1, 0, 0)
	Right   = NewVector3(1, 0, 0)
	Forward = NewVector3(0, 0, 1)
	Back    = NewVector3(0, 0, -1)
)

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

// NewVector3
//
//	@Description: 创建一个三维向量
//	@param x
//	@param y
//	@param z
//	@return Vector3
func NewVector3(x float32, y float32, z float32) Vector3 {
	return Vector3{
		X: x,
		Y: y,
		Z: z,
	}
}

func (v Vector3) Set(nX float32, nY float32, nZ float32) {
	v.X = nX
	v.Y = nY
	v.Z = nZ
}

func (v Vector3) Equals(other Vector3) bool {
	return v.X == other.X && v.Y == other.Y && v.Z == other.Z
}
