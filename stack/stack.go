package stack

type Node[T any] struct {
	value T
	prev  *Node[T]
}

func newNode[T any](val T) *Node[T] {
	return &Node[T]{
		value: val,
		prev:  nil,
	}
}

type Stack[T any] struct {
	length int
	head   *Node[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		length: 0,
		head:   nil,
	}
}

func (s *Stack[T]) Push(val T) {
	s.length++
	node := newNode(val)
	if s.head == nil {
		s.head = node
		return
	}
	node.prev = s.head
	s.head = node
}

func (s *Stack[T]) Pop() *T {
	if s.length <= 0 || s.head == nil {
		return nil
	}
	s.length--
	currentHead := s.head
	if s.length == 0 {
		s.head = nil
		return &currentHead.value
	}

	s.head = currentHead.prev

	return &currentHead.value
}

func (s *Stack[T]) Peek() *T {
	if s.head == nil {
		return nil
	}
	return &s.head.value
}
