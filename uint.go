package buffer

import (
	"encoding/binary"
)

// ReadUint8 ...
func (buff *Buffer) ReadUint8(offset uint) uint8 {
	b := []byte(*buff)[offset]
	return uint8(b)
}

// ReadUint16BE ...
func (buff *Buffer) ReadUint16BE(offset uint) uint16 {
	b := []byte(*buff)[offset : offset+2]
	return binary.BigEndian.Uint16(b)
}

// ReadUint16LE ...
func (buff *Buffer) ReadUint16LE(offset uint) uint16 {
	b := []byte(*buff)[offset : offset+2]
	return binary.LittleEndian.Uint16(b)
}

// ReadUint32BE ...
func (buff *Buffer) ReadUint32BE(offset uint) uint32 {
	b := []byte(*buff)[offset : offset+4]
	return binary.BigEndian.Uint32(b)
}

// ReadUint32LE ...
func (buff *Buffer) ReadUint32LE(offset uint) uint32 {
	b := []byte(*buff)[offset : offset+4]
	return binary.LittleEndian.Uint32(b)
}

// ReadUint64BE ...
func (buff *Buffer) ReadUint64BE(offset uint) uint64 {
	b := []byte(*buff)[offset : offset+8]
	return binary.BigEndian.Uint64(b)
}

// ReadUint64LE ...
func (buff *Buffer) ReadUint64LE(offset uint) uint64 {
	b := []byte(*buff)[offset : offset+8]
	return binary.LittleEndian.Uint64(b)
}
