package tree

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

func PrintTree[T any](head *Node[T]) *[]T {
	return walk(head, &[]T{})
}
