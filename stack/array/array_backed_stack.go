package array

// Stack is an interface for all stacks
type Stack interface {
	IsEmpty() bool
	Push(item string)
	Pop() string
}

// NewStack creates and returns new ArrayBackedStack instances
func NewStack() Stack {
	return &FixedSizeStack{}
}

// FixedSizeStack is an implementation of Stack interface
// backed by array storage
type FixedSizeStack struct {
	Strings [10]string
	Idx     int
}

func (s *FixedSizeStack) IsEmpty() bool {
	return s.Idx == 0
}

func (s *FixedSizeStack) Push(item string) {
	s.Strings[s.Idx] = item
	s.Idx++
}

func (s *FixedSizeStack) Pop() string {
	s.Idx--
	return s.Strings[s.Idx]
}
