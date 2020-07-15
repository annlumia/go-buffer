package buffer

import "encoding/hex"

// Buffer ...
type Buffer []byte

// FromBytes create buffer type from slice of byte
func FromBytes(bytes []byte) Buffer {
	return bytes
}

// FromStringHex create buffer from string hex
func FromStringHex(str string) (Buffer, error) {
	data, err := hex.DecodeString(str)
	return data, err
}

// FromString create buffer from string
func FromString(str string) Buffer {
	return []byte(str)
}

// Bytes ...
func (b *Buffer) Bytes() []byte {
	return []byte(*b)
}
