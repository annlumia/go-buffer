package buffer

import "testing"

func TestUint(t *testing.T) {
	buff := Buffer{0xfa, 0x74, 0x32, 0x22, 0x85, 0x88, 0x82, 0x98}

	if res := buff.ReadUint8(0); res != 250 {
		t.Errorf("buffer.ReadUint8 failed. Expect %v actually %v", 250, res)
	}

	if res := buff.ReadUint16LE(0); res != 29946 {
		t.Errorf("buffer.ReadUint16LE failed. Expect %v actually %v", 29946, res)
	}

	if res := buff.ReadUint16BE(0); res != 64116 {
		t.Errorf("buffer.ReadUint16BE failed. Expect %v actually %v", 64116, res)
	}

	if res := buff.ReadUint32LE(0); res != 573732090 {
		t.Errorf("buffer.ReadUint32LE failed. Expect %v actually %v", 573732090, res)
	}

	if res := buff.ReadUint32BE(0); res != 4201919010 {
		t.Errorf("buffer.ReadUint32BE failed. Expect %v actually %v", 4201919010, res)
	}

	if res := buff.ReadUint64LE(0); res != 10989496146123191546 {
		t.Errorf("buffer.ReadUint64LE failed. Expect %v actually %v", "10989496146123191546", res)
	}

	if res := buff.ReadUint64BE(0); res != 18047104730631013016 {
		t.Errorf("buffer.ReadUint64BE failed. Expect %v actually %v", "18047104730631013016", res)
	}
}
