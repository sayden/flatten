package flatten

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"math/rand"
	"testing"
	"time"
)

func Test_Fraud_BytesOps(t *testing.T) {

	t.Run("Int to bytes and to Int", func(t *testing.T) {

		input := getRandomSignedInt32()

		bytes := int32ToBytes(input)
		if input != int32FromBytes(bytes) {
			t.Errorf(fmt.Sprintf("Failed Conversion for int32: %v", input))

		}
	})

	t.Run("Float64 to bytes and to Float64", func(t *testing.T) {
		rand.Seed(time.Now().UTC().UnixNano())

		input := getRandomSignedFloat64()

		bytes := float64ToBytes(input)

		if input != float64FromBytes(bytes) {
			t.Errorf(fmt.Sprintf("Failed Conversion for float64: %v", input))
		}
	})

	t.Run("Check nil equality", func(t *testing.T) {

		if !equal(nil, nil) {
			t.Errorf("Failed equality for two nil bytes arrays")
		}

	})

	t.Run("Check nil and array equality", func(t *testing.T) {

		if equal(nil, int32ToBytes(getRandomSignedInt32())) {
			t.Errorf("Failed equality for nil array and array")
		}

	})

	t.Run("Check array and nil equality", func(t *testing.T) {

		if equal(int32ToBytes(getRandomSignedInt32()), nil) {
			t.Errorf("Failed equality for array and nil array")
		}

	})

	t.Run("Check equality", func(t *testing.T) {
		input := getRandomSignedInt32()

		bytes1 := int32ToBytes(input)
		bytes2 := int32ToBytes(input)

		if !equal(bytes1, bytes2) {
			t.Errorf(fmt.Sprintf("Failed equality between %v and %v", input, input))
		}

	})

	t.Run("Check inequality", func(t *testing.T) {
		input1 := int32(0)
		input2 := int32(1)

		bytes1 := int32ToBytes(input1)
		bytes2 := int32ToBytes(input2)

		if equal(bytes1, bytes2) {
			t.Errorf(fmt.Sprintf("Failed Inequality between %v and %v", input1, input2))
		}

	})

	t.Run("Reduce bytes array size", func(t *testing.T) {
		n := 5
		byt := intToBytes(n)
		assert.Equal(t, []byte{5}, byt)
		assert.Equal(t, n, intFromBytes(byt))
	})

	t.Run("Negative numbers and big ops", func(t *testing.T) {
		n := math.MinInt64
		byt := intToBytes(n)
		assert.Equal(t, 0, intFromBytes(byt))

		var n8 int8 = math.MinInt8
		byt = int8ToBytes(n8)
		assert.Equal(t, n8, int8FromBytes(byt))

		var n16 int16 = math.MinInt16
		byt = int16ToBytes(n16)
		assert.Equal(t, n16, int16FromBytes(byt))

		var n32 int32 = math.MinInt32
		byt = int32ToBytes(n32)
		assert.Equal(t, n32, int32FromBytes(byt))

		var n64 int64 = math.MinInt64
		byt = int64ToBytes(n64)
		assert.Equal(t, n64, int64FromBytes(byt))

		n64 = math.MaxInt64
		byt = int64ToBytes(n64)
		assert.Equal(t, n64, int64FromBytes(byt))
	})

}

func Test_IntToBytes(t *testing.T) {
	var n int = 33

	byt := intToBytes(n)
	assert.Equal(t, uint8(33), byt[len(byt)-1])
}

func getRandomSign() (sign int) {
	sign = 1
	rand.Seed(time.Now().UTC().UnixNano())
	if rand.Intn(2) == 1 {
		sign = -1
	}

	return
}

func getRandomSignedInt32() (value int32) {
	rand.Seed(time.Now().UTC().UnixNano())
	return int32(getRandomSign() * rand.Int())
}

func getRandomSignedFloat64() (value float64) {
	rand.Seed(time.Now().UTC().UnixNano())
	return float64(getRandomSign()) * rand.Float64()

}
