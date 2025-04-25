package tree

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestStack(t *testing.T) {

	t.Run("Test tree", func(t *testing.T) {
		node1 := NewNode(2, NewNode(3, NewNode(5, nil, nil), NewNode(15, nil, nil)), NewNode(20, nil, nil))

		result := PrintTree(node1)

		expect := []int{
			2,
			3,
			5,
			15,
			20,
		}

		assert.Equal(t, result, expect)
	})

}
