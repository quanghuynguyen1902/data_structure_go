package stack

// Stack is the structure which contains our elements
type Stack struct {
	size   int
	stack  []interface{}
	curInd int
}

// NewStack returns a pointer to a new stack having:
// - the array initialized with the provided size
// - the current index initialized with -1 (empty stack)
func NewStack(size int) *Stack {
	return &Stack{
		size:   size,
		stack:  make([]interface{}, size),
		curInd: -1,
	}
}

// IsEmpty returns true if the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.curInd < 0
}

func (s *Stack) Top() interface{} {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	return s.stack[s.curInd]
}

// Pop removes and returns the element at the top of the stack.
// This method panics if the stack is empty
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	el := s.stack[s.curInd]
	s.stack[s.curInd] = nil
	s.curInd--
	return el
}

// Push inserts the element at the top of the stack.
// This method panics if the stack is full
func (s *Stack) Push(el interface{}) {
	if s.IsFull() {
		panic("stack is full")
	}
	s.curInd++
	s.stack[s.curInd] = el
}

func (s *Stack) IsFull() bool {
	return s.curInd+1 == s.size
}
