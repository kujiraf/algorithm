package stack

import (
	"fmt"
	"testing"
)

func TestPushPop(t *testing.T) {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	e := "[1 2 3 4 5]"
	if a := fmt.Sprint(stack.data); a != e {
		t.Errorf("push assertion error: got %s, want %s", a, e)
	}

	p := stack.Pop()
	if p.(int) != 5 {
		t.Errorf("pop assertion error: got %d, want %d", p, 5)
	}
	e = "[1 2 3 4]"
	if a := fmt.Sprint(stack.data); a != e {
		t.Errorf("pop assertion error: got %s, want %s", a, e)
	}
	stack.Pop()
	p = stack.Pop()
	if p.(int) != 3 {
		t.Errorf("pop assertion error: got %d, want %d", p, 3)
	}
	e = "[1 2]"
	if a := fmt.Sprint(stack.data); a != e {
		t.Errorf("pop assertion error: got %s, want %s", a, e)
	}
	stack.Pop()
	stack.Pop()
	e = "[]"
	if a := fmt.Sprint(stack.data); a != e {
		t.Errorf("pop assertion error: got %s, want %s", a, e)
	}
	stack.Pop()
	stack.Pop()
	p = stack.Pop()
	if _, ok := p.(int); ok {
		t.Errorf("pop assertion error: got %d, want nil", p)
	}
	e = "[]"
	if a := fmt.Sprint(stack.data); a != e {
		t.Errorf("pop assertion error: got %s, want %s", a, e)
	}
}
