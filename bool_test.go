package buffer

import "testing"

func TestBool(t *testing.T) {
	buff := Buffer{0x0}
	bits := []bool{false, true, true, false, true, false, false, true}

	for i := 0; i < len(bits); i++ {
		buff.WriteBool(0, uint8(i), bits[i])
	}

	for i := 0; i < len(bits); i++ {
		bit := buff.ReadBool(0, uint8(i))
		if bit != bits[i] {
			t.Errorf("Unmatch in bit number %v. Expect %v but got %v", i, bits[i], bit)
		}
	}
}
