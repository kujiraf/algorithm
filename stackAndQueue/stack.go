package sq

import "fmt"

type stack struct {
	data []interface{}
}

func (s *stack) push(d interface{}) {
	s.data = append(s.data, d)
}

func (s *stack) pop() interface{} {
	l := len(s.data)
	if l == 0 {
		return nil
	}

	d := s.data[l-1]
	s.data = s.data[:l-1]
	fmt.Println("poped: ", s.data)
	return d
}

var brackets = map[rune]rune{'{': '}', '[': ']', '(': ')'}

func ValidateBrackets(s string) bool {
	if len(s) <= 1 {
		return false
	}

	stack := stack{}
	for _, r := range s {
		if bRight, ok := brackets[r]; ok {
			stack.push(bRight)
			continue
		}

		for _, v := range brackets {
			if r == v {
				if bRight := stack.pop(); bRight == r {
					break
				}
				return false
			}
		}
	}

	if len(stack.data) > 0 {
		return false
	}
	return true
}
