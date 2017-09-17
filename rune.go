package flatten

func bytesToRune(byt []byte) rune {
	return int32FromBytes(byt)
}

func runeToByts(r rune) []byte {
	return int32ToBytes(int32(r))
}
