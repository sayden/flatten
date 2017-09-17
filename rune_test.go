package flatten

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Rune(t *testing.T) {
	var r rune = 'a'
	byt := runeToByts(r)
	assert.Equal(t, r, bytesToRune(byt))
}
