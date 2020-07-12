package buffer

// ReadBool ...
func (buff *Buffer) ReadBool(byteOffset int, bitNumber uint8) bool {
	return (*buff)[byteOffset]>>bitNumber&0x1 == 1
}

// WriteBool ...
func (buff *Buffer) WriteBool(byteOffset int, bitNumber uint8, val bool) {
	if val {
		(*buff)[byteOffset] = SetBitInByte((*buff)[byteOffset], bitNumber)
	} else {
		(*buff)[byteOffset] = ClearBitInByte((*buff)[byteOffset], bitNumber)
	}
}

// SetBitInByte ...
func SetBitInByte(b byte, bitNumber uint8) byte {
	b |= (1 << bitNumber)
	return b
}

// ClearBitInByte ...
func ClearBitInByte(b byte, bitNumber uint8) byte {
	var mask byte
	mask = ^(1 << bitNumber)
	b &= mask
	return b
}
