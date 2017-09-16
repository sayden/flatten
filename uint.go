package flatten

import (
	"encoding/binary"
	"runtime"
)

func uint8ToBytes(u uint8) []byte {
	return []byte{u}
}

func uint16ToBytes(u uint16) []byte {
	var buf [2]byte
	binary.BigEndian.PutUint16(buf[:], u)

	return buf[:]
}

func uint32ToBytes(u uint32) []byte {
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:], u)
	return buf[:]
}

func uint64ToBytes(u uint64) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], u)
	return buf[:]
}

func uintToBytes(u uint) []byte {
	if runtime.GOARCH == "amd64" {
		var buf [8]byte
		binary.BigEndian.PutUint64(buf[:], uint64(u))
		return buf[:]
	}

	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:], uint32(u))
	return buf[:]
}
