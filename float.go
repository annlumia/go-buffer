package buffer

import (
	"encoding/binary"
	"math"
)

// ReadFloat32BE ...
func (buff *Buffer) ReadFloat32BE(offset uint) float32 {
	b := []byte(*buff)[offset : offset+4]
	u := binary.BigEndian.Uint32(b)
	return math.Float32frombits(u)
}

// ReadFloat32LE ...
func (buff *Buffer) ReadFloat32LE(offset uint) float32 {
	b := []byte(*buff)[offset : offset+4]
	u := binary.LittleEndian.Uint32(b)
	return math.Float32frombits(u)
}

// ReadFloat64BE ...
func (buff *Buffer) ReadFloat64BE(offset uint) float64 {
	b := []byte(*buff)[offset : offset+8]
	u := binary.BigEndian.Uint64(b)
	return math.Float64frombits(u)
}

// ReadFloat64LE ...
func (buff *Buffer) ReadFloat64LE(offset uint) float64 {
	b := []byte(*buff)[offset : offset+8]
	u := binary.LittleEndian.Uint64(b)
	return math.Float64frombits(u)
}
