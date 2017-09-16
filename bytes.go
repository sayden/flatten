package flatten

func bytes(i interface{}) []byte {
	switch t := i.(type) {
	case float64:
		return float64ToBytes(t)

	case string:
		return stringToBytes(t)

	case bool:
		return boolToBytes(t)

	case uint:
		return uintToBytes(t)
	case uint8:
		return uint8ToBytes(t)
	case uint16:
		return uint16ToBytes(t)
	case uint32:
		return uint32ToBytes(t)
	case uint64:
		return uint64ToBytes(t)

	case int:
		return intToBytes(t)
	case int8:
		return int8ToBytes(t)
	case int16:
		return int16ToBytes(t)
	case int32:
		return int32ToBytes(t)
	case int64:
		return int64ToBytes(t)
	}

	return nil
}

func stringToBytes(s string) []byte {
	return []byte(s)
}

func boolToBytes(t bool) []byte {
	if t {
		return []byte{1}
	}
	return []byte{0}
}

func equal(b1, b2 []byte) (r bool) {
	if b1 == nil && b2 == nil {
		return true
	}

	if len(b1) != len(b2) {
		return
	}

	for i, v2 := range b2 {
		if v2 != b1[i] {
			return
		}
	}

	return true
}
