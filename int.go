package buffer

import (
	"encoding/binary"
)

// ReadInt8 ...
func (buff *Buffer) ReadInt8(offset uint) int8 {
	b := []byte(*buff)[offset]
	return int8(b)
}

// ReadInt16BE ...
func (buff *Buffer) ReadInt16BE(offset uint) int16 {
	b := []byte(*buff)[offset : offset+2]
	return int16(binary.BigEndian.Uint16(b))
}

// ReadInt16LE ...
func (buff *Buffer) ReadInt16LE(offset uint) int16 {
	b := []byte(*buff)[offset : offset+2]
	return int16(binary.LittleEndian.Uint16(b))
}

// ReadInt32BE ...
func (buff *Buffer) ReadInt32BE(offset uint) int32 {
	b := []byte(*buff)[offset : offset+4]
	return int32(binary.BigEndian.Uint32(b))
}

// ReadInt32LE ...
func (buff *Buffer) ReadInt32LE(offset uint) int32 {
	b := []byte(*buff)[offset : offset+4]
	return int32(binary.LittleEndian.Uint32(b))
}

// ReadInt64BE ...
func (buff *Buffer) ReadInt64BE(offset uint) int64 {
	b := []byte(*buff)[offset : offset+8]
	return int64(binary.BigEndian.Uint64(b))
}

// ReadInt64LE ...
func (buff *Buffer) ReadInt64LE(offset uint) int64 {
	b := []byte(*buff)[offset : offset+8]
	return int64(binary.LittleEndian.Uint64(b))
}
