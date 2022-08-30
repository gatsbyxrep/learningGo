package myStack

type node struct {
	value interface{}
	next  *node
}

type Stack struct {
	head *node
	size int
}

func (s *Stack) Push(value interface{}) {
	newNode := new(node)
	newNode.value = value
	if s.head == nil {
		s.head = newNode
	} else {
		s.head.next = newNode
	}
	s.size++
}

func (s *Stack) Pop() interface{} {
	if s.head == nil {
		panic("no elements in Stack")
	}
	currentHead := s.head
	s.head = s.head.next
	s.size--
	return currentHead.value
}
