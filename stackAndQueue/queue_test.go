package sq

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := queue{}
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.enqueue(4)
	q.enqueue(5)

	e := "[1 2 3 4 5]"
	if a := fmt.Sprint(q.data); a != e {
		t.Errorf("push assertion error: got %s, want %s", a, e)
	}

	d := q.dequeue()
	if d.(int) != 1 {
		t.Errorf("dequeue assertion error: got %d, want %d", d, 1)
	}
	e = "[2 3 4 5]"
	if a := fmt.Sprint(q.data); a != e {
		t.Errorf("dequeue assertion error: got %s, want %s", a, e)
	}
	q.dequeue()
	d = q.dequeue()
	if d.(int) != 3 {
		t.Errorf("dequeue assertion error: got %d, want %d", d, 3)
	}
	e = "[4 5]"
	if a := fmt.Sprint(q.data); a != e {
		t.Errorf("dequeue assertion error: got %s, want %s", a, e)
	}
	q.dequeue()
	q.dequeue()
	e = "[]"
	if a := fmt.Sprint(q.data); a != e {
		t.Errorf("dequeue assertion error: got %s, want %s", a, e)
	}
	q.dequeue()
	q.dequeue()
	d = q.dequeue()
	if _, ok := d.(int); ok {
		t.Errorf("dequeue assertion error: got %d, want nil", d)
	}
	e = "[]"
	if a := fmt.Sprint(q.data); a != e {
		t.Errorf("dequeue assertion error: got %s, want %s", a, e)
	}
}

func TestReverse(t *testing.T) {
	q := queue{}
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.enqueue(4)
	q.enqueue(5)
	e := "[1 2 3 4 5]"
	if a := fmt.Sprint(q.data); a != e {
		t.Errorf("push assertion error: got %s, want %s", a, e)
	}

	q.reverse()
	e = "[5 4 3 2 1]"
	if a := fmt.Sprint(q.data); a != e {
		t.Errorf("push assertion error: got %s, want %s", a, e)
	}

}
