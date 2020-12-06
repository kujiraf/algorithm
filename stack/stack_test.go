package stack

import (
	"fmt"
	"testing"
)

func TestPushPop(t *testing.T) {
	stack := stack{}
	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.push(4)
	stack.push(5)
	e := "[1 2 3 4 5]"
	if a := fmt.Sprint(stack.data); a != e {
		t.Errorf("push assertion error: got %s, want %s", a, e)
	}

	p := stack.pop()
	if p.(int) != 5 {
		t.Errorf("pop assertion error: got %d, want %d", p, 5)
	}
	e = "[1 2 3 4]"
	if a := fmt.Sprint(stack.data); a != e {
		t.Errorf("pop assertion error: got %s, want %s", a, e)
	}
	stack.pop()
	p = stack.pop()
	if p.(int) != 3 {
		t.Errorf("pop assertion error: got %d, want %d", p, 3)
	}
	e = "[1 2]"
	if a := fmt.Sprint(stack.data); a != e {
		t.Errorf("pop assertion error: got %s, want %s", a, e)
	}
	stack.pop()
	stack.pop()
	e = "[]"
	if a := fmt.Sprint(stack.data); a != e {
		t.Errorf("pop assertion error: got %s, want %s", a, e)
	}
	stack.pop()
	stack.pop()
	p = stack.pop()
	if _, ok := p.(int); ok {
		t.Errorf("pop assertion error: got %d, want nil", p)
	}
	e = "[]"
	if a := fmt.Sprint(stack.data); a != e {
		t.Errorf("pop assertion error: got %s, want %s", a, e)
	}
}

func TestValidateBrackets(t *testing.T) {
	passData := "{[()()][]()[]()}"
	failedData := "{[()()][]()[](){"

	if !ValidateBrackets(passData) {
		t.Error("pass data is failed")
	}

	if ValidateBrackets(failedData) {
		t.Error("failed data is passed")
	}
}
