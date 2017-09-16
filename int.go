package flatten

import "encoding/binary"

func intToBytes(i int) (byt []byte) {
	return uintToBytes(uint(i))
}

func int8ToBytes(i int8) (byt []byte) {
	return []byte{uint8(i)}
}

func int16ToBytes(u int16) []byte {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, uint16(u))

	return buf
}

func int32ToBytes(i int32) (byt []byte) {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(i))
	byt = buf

	return
}

func int64ToBytes(u int64) []byte {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(u))
	return buf
}

func int32FromBytes(byt []byte) int32 {
	return int32(binary.BigEndian.Uint32(byt))
}