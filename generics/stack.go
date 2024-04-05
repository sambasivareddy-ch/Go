package generics

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) StackType() T {
	var stkType T 
	return stkType
}

func (s *Stack[T]) Push(val T) {
	s.data = append(s.data, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) == 0 {
		var zeroedValue T 
		return zeroedValue, false
	}
	lastIndex := len(s.data)-1
	value := s.data[lastIndex]
	s.data = s.data[:lastIndex]
	return value, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func (s *Stack[T]) PrintStack() {
	PrintArray(s.data)
}