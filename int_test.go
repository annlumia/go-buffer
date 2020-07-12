package buffer

import "testing"

func TestInt(t *testing.T) {
	buff := Buffer{0x16, 0x16, 0xe9, 0xe9}

	if res := buff.ReadInt8(3); res != -23 {
		t.Errorf("buffer.ReadIntBE failed. Expect %v actually %v", -23, res)
	}

	if res := buff.ReadInt16LE(1); res != -5866 {
		t.Errorf("buffer.ReadIntBE failed. Expect %v actually %v", -5866, res)
	}

	if res := buff.ReadInt16BE(1); res != 5865 {
		t.Errorf("buffer.ReadIntBE failed. Expect %v actually %v", 5865, res)
	}
}
