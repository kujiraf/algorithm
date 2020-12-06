package stack

type Stack struct {
	data []interface{}
}

func (s *Stack) Push(d interface{}) {
	s.data = append(s.data, d)
}

func (s *Stack) Pop() interface{} {
	l := len(s.data)
	if l == 0 {
		return nil
	}

	d := s.data[l-1]
	s.data = s.data[:l-1]
	return d
}
