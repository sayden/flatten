package flatten

import (
	"encoding/binary"
)

func int8FromBytes(byt []byte) int8 {
	return int8(uint8FromBytes(byt))
}

func int16FromBytes(byt []byte) int16 {
	return int16(uint16FromBytes(byt))
}

func int32FromBytes(byt []byte) int32 {
	return int32(uint32FromBytes(byt))
}

func int64FromBytes(byt []byte) int64 {
	return int64(uint64FromBytes(byt))
}

func intFromBytes(byt []byte) int {
	if len(byt) != 4 {
		prefix := make([]byte, 4-len(byt))
		return int(binary.BigEndian.Uint32(append(prefix, byt...)))
	}

	return int(binary.BigEndian.Uint32(byt))
}

func intToBytes(i int) []byte {
	return uintToBytes(uint(i))
}

func int8ToBytes(i int8) (byt []byte) {
	return []byte{uint8(i)}
}

func int16ToBytes(u int16) []byte {
	return uint16ToBytes(uint16(u))
}

func int32ToBytes(i int32) (byt []byte) {
	return uint32ToBytes(uint32(i))
}

func int64ToBytes(i int64) []byte {
	return uint64ToBytes(uint64(i))
}
