package flatten

import (
	"github.com/cabify/fraud/bytes_ops"
	"time"
)

func getBytes(i interface{}) (byt []byte) {

	switch t := i.(type) {
	case uint8:
	case int:
		byt = bytes_ops.IntToBytes(t)
	case string:
		byt = []byte(t)
	case time.Time:
		byt = []byte(t.String())
	case float64:
		byt = bytes_ops.Float64ToBytes(t)
	case bool:
		if t {
			return []byte{1}
		} else {
			return []byte{0}
		}
	}

	return
}
