package tree

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestStack(t *testing.T) {

	t.Run("PreOrderSearch", func(t *testing.T) {
		node1 := NewNode(2, NewNode(3, NewNode(5, nil, nil), NewNode(15, nil, nil)), NewNode(20, nil, nil))

		result := PreOrderSearch(node1)

		expect := []int{
			2,
			3,
			5,
			15,
			20,
		}

		assert.Equal(t, result, expect)
	})

	t.Run("InOrderSearch", func(t *testing.T) {
		node1 := NewNode(2, NewNode(3, NewNode(5, nil, nil), NewNode(15, nil, nil)), NewNode(20, nil, nil))

		result := InOrderSearch(node1)

		expect := []int{
			5,
			3,
			15,
			2,
			20,
		}

		assert.Equal(t, result, expect)
	})
	t.Run("PostOrderSearch", func(t *testing.T) {
		node1 := NewNode(2, NewNode(3, NewNode(5, nil, nil), NewNode(15, nil, nil)), NewNode(20, NewNode(1, nil, nil), NewNode(2, nil, nil)))

		result := PostOrderSearch(node1)

		expect := []int{
			5,
			15,
			3,
			1,
			2,
			20,
			2,
		}

		assert.Equal(t, result, expect)
	})
	t.Run("TestLooseWeigh", func(t *testing.T) {
		TestLooseWeigh(77, 70, 1)

	})

}
