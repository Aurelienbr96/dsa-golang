package bubble_sort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {

	t.Run("Should sort array", func(t *testing.T) {
		unsorted := []int{1, 5, 3, 8, 4, 10, 15, 20, 14, 10, 11, 18}
		Sort(unsorted)
		fmt.Printf("%v", unsorted)
		assert.Equal(t, unsorted, []int{1, 3, 4, 5, 8, 10, 10, 11, 14, 15, 18, 20})
	})

	t.Run("Should sort single-element array", func(t *testing.T) {
		unsorted := []int{1}
		Sort(unsorted)
		assert.Equal(t, unsorted, []int{1})
	})
}
