package tree

import (
	"fmt"
	"math"
)

type Node[T any] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

func NewNode[T any](value T, left *Node[T], right *Node[T]) *Node[T] {
	return &Node[T]{
		Value: value,
		Left:  left,
		Right: right,
	}
}

func walk[T any](curr *Node[T], path *[]T) *[]T {
	if curr == nil {
		return path
	}

	*path = append(*path, curr.Value)
	walk(curr.Left, path)
	walk(curr.Right, path)

	return path
}

// depth first search, DFS
func PreOrderSearch[T any](head *Node[T]) *[]T {
	return walk(head, &[]T{})
}

func walkInOrderSearch[T any](curr *Node[T], path *[]T) *[]T {
	if curr == nil {
		return path
	}

	walkInOrderSearch(curr.Left, path)
	*path = append(*path, curr.Value)
	walkInOrderSearch(curr.Right, path)

	return path
}

func InOrderSearch[T any](head *Node[T]) *[]T {
	return walkInOrderSearch(head, &[]T{})
}

func walkPostOrderSearch[T any](curr *Node[T], path *[]T) *[]T {
	if curr == nil {
		return path
	}

	walkPostOrderSearch(curr.Left, path)
	walkPostOrderSearch(curr.Right, path)
	*path = append(*path, curr.Value)

	return path
}

func PostOrderSearch[T any](head *Node[T]) *[]T {
	return walkPostOrderSearch(head, &[]T{})
}

func TestLooseWeigh(weight float64, goalWeight float64, dept int32) string {
	if goalWeight > weight {
		return "Goal weigh too high"
	}
	if weight <= goalWeight {
		return "goal achieve"
	}
	weightW := weight * 0.005
	finalWeight := weight - weightW
	fmt.Printf("(week %d) This week you should loose %vkg and weight %v \n", dept, math.Round(weightW*10)/10, math.Round(finalWeight*10)/10)
	TestLooseWeigh(finalWeight, goalWeight, dept+1)
	return ""
}
