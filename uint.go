package flatten

import (
	"encoding/binary"
)

func uint8ToBytes(u uint8) []byte {
	return []byte{u}
}

func uint16ToBytes(u uint16) []byte {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, u)

	if buf[0] != 0 {
		return buf
	}

	return buf[1:]
}

func uint32ToBytes(u uint32) []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, u)

	for i := 0; i < 3; i++ {
		if buf[i] != 0 {
			return buf[i:]
		}
	}

	return buf[3:]
}

func uint64ToBytes(u uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, u)

	for i := 0; i < 7; i++ {
		if buf[i] != 0 {
			return buf[i:]
		}
	}

	return buf[7:]
}

func uintToBytes(u uint) []byte {
	return uint32ToBytes(uint32(u))
}

func uintFromBytes(byt []byte) uint {
	if len(byt) != 4 {
		prefix := make([]byte, 4-len(byt))
		return uint(binary.BigEndian.Uint32(append(prefix, byt...)))
	}

	return uint(binary.BigEndian.Uint32(byt))
}

func uint8FromBytes(byt []byte) uint8 {
	return uint8(binary.BigEndian.Uint16(append([]byte{0}, byt...)))
}

func uint16FromBytes(byt []byte) uint16 {
	if len(byt) != 2 {
		prefix := make([]byte, 2-len(byt))
		return binary.BigEndian.Uint16(append(prefix, byt...))
	}

	return binary.BigEndian.Uint16(byt)
}

func uint32FromBytes(byt []byte) uint32 {
	if len(byt) != 4 {
		prefix := make([]byte, 4-len(byt))
		return binary.BigEndian.Uint32(append(prefix, byt...))
	}

	return binary.BigEndian.Uint32(byt)
}

func uint64FromBytes(byt []byte) uint64 {
	if len(byt) != 8 {
		prefix := make([]byte, 8-len(byt))
		return binary.BigEndian.Uint64(append(prefix, byt...))
	}

	return binary.BigEndian.Uint64(byt)
}
