package flatten

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Complex(t *testing.T) {
	t.Run("complex64", func(t *testing.T) {
		var c64 complex64 = 5 + 4i
		byt := complex64ToBytes(c64)
		assert.Equal(t, []byte{64, 160, 0, 0, 64, 128, 0, 0}, byt)
	})

	t.Run("complex128", func(t *testing.T) {
		var c64 complex128 = 5 + 4i
		byt := complex128ToBytes(c64)
		assert.Equal(t, []byte{64, 20, 0, 0, 0, 0, 0, 0, 64, 16, 0, 0, 0, 0, 0, 0}, byt)
	})
}
