package core

import (
	"math"
	"mmo_server/ProtoMessage"
)

type Vector3Int struct {
	X int32
	Y int32
	Z int32
}

func NewVector3Int(x int32, y int32, z int32) Vector3Int {
	return Vector3Int{
		X: x,
		Y: y,
		Z: z,
	}
}

func ProtoConvertVectorInt(pv *ProtoMessage.PVector3) Vector3Int {
	return Vector3Int{
		X: pv.X,
		Y: pv.Y,
		Z: pv.Z,
	}
}

func (v Vector3Int) Minus(b Vector3Int) Vector3Int {
	return Vector3Int{
		X: v.X - b.X,
		Y: v.Y - b.Y,
		Z: v.Z - b.Z,
	}
}

func (v Vector3Int) Magnitude() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z)))
}

func (v Vector3Int) Distance(b Vector3Int) float32 {
	return v.Minus(b).Magnitude()
}

func (v Vector3Int) ConvertProto() *ProtoMessage.PVector3 {
	return &ProtoMessage.PVector3{
		X: v.X,
		Y: v.Y,
		Z: v.Z,
	}
}
