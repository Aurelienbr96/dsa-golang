package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {

	t.Run("Return the index of the searched value", func(t *testing.T) {
		arr := []int{1, 3, 4, 5, 8, 10, 10, 11, 14, 15, 18, 20}
		val := BinarySearch(arr, 3)
		assert.Equal(t, 1, val)
	})

	t.Run("Return the index of the searched value", func(t *testing.T) {
		arr := []int{1, 3, 4, 5, 8, 10, 10, 11, 14, 15, 18, 20}
		val := BinarySearch(arr, 14)
		assert.Equal(t, 8, val)
	})

	t.Run("Return -1 if the value does not exist in the array", func(t *testing.T) {
		arr := []int{1, 3, 4, 5, 8, 10, 10, 11, 14, 15, 18, 20}
		val := BinarySearch(arr, 30)
		assert.Equal(t, -1, val)
	})

}
