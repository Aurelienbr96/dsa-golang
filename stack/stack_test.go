package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {

	type test struct {
		input  int
		output int
	}

	t.Run("Should Push an element to the Stack and increase its length", func(t *testing.T) {
		stack := NewStack[int]()
		stack.Push(3)

		assert.Equal(t, 1, stack.length)
		assert.Equal(t, 3, stack.head.value)
		assert.Nil(t, stack.head.prev)
	})

	t.Run("Should Push 2 elements to the Stack and increase its length", func(t *testing.T) {
		stack := NewStack[int]()
		stack.Push(3)
		stack.Push(4)

		assert.Equal(t, 2, stack.length)
		assert.Equal(t, 4, stack.head.value)
		assert.Equal(t, 3, stack.head.prev.value)
	})

	t.Run("Should Push 2 elements to the Stack and then pop and return its last element", func(t *testing.T) {
		stack := NewStack[int]()
		stack.Push(3)
		stack.Push(4)
		stack.Push(5)
		result := stack.Pop()

		assert.Equal(t, *result, 5)
		assert.Equal(t, stack.length, 2)
		assert.Equal(t, 4, stack.head.value)
		assert.Equal(t, 3, stack.head.prev.value)
	})

	t.Run("Should return nil if stack is empty when trying to pop", func(t *testing.T) {
		stack := NewStack[int]()
		result := stack.Pop()

		assert.Nil(t, result)
	})

	t.Run("Should return nil if stack is empty when trying to pop", func(t *testing.T) {
		stack := NewStack[int]()
		stack.Push(4)
		result := stack.Pop()

		assert.Equal(t, 4, *result)
		assert.Nil(t, stack.head)
		assert.Equal(t, stack.length, 0)
	})
}
