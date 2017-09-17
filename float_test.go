package flatten

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"math"
)

func Test_Float(t *testing.T) {
	t.Run("Test bytes with array reduction", func(t *testing.T){
		var n float64 = math.MaxFloat64
		byt := float64ToBytes(n)
		assert.Equal(t, n, float64FromBytes(byt))

		var n32 float32 = math.MaxFloat32
		byt = float32ToBytes(n32)
		assert.Equal(t, n32, float32FromBytes(byt))

		n = math.SmallestNonzeroFloat64
		byt = float64ToBytes(n)
		assert.Equal(t, n, float64FromBytes(byt))

		n32 = math.SmallestNonzeroFloat32
		byt = float32ToBytes(n32)
		assert.Equal(t, n32, float32FromBytes(byt))
	})
}
