package twocrystalball

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoCrystalBall(t *testing.T) {

	t.Run("Find the value", func(t *testing.T) {
		array := []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, true, true, true}
		val := TwoCrystalBall(array)
		assert.Equal(t, 13, val)
	})

}
