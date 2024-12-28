package linearsearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinearSearch(t *testing.T) {

	t.Run("Find the value", func(t *testing.T) {
		array := []int{1, 5, 3, 8, 4, 10, 15, 20, 14, 10, 11, 18}
		val := LinearSearch(array, 1)
		assert.Equal(t, val, 0)
	})

	t.Run("Should return -1 if the value we are looking for is not in the array", func(t *testing.T) {
		array := []int{1, 5, 6, 8}
		val := LinearSearch(array, 4)
		assert.Equal(t, val, -1)
	})
}
