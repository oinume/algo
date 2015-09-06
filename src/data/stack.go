package data

type Stack struct {
	data []Object
}

func NewStack(capacity int) *Stack {
	data := make([]Object, 0, capacity)
	return &Stack{data: data}
}

func (s *Stack) Push(o Object) {
	s.data = append(s.data, o)
}

func (s *Stack) Pop() Object {
	o := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return o
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) Clear() {
}
