package sq

import "fmt"

type queue struct {
	data []interface{}
}

func (q *queue) enqueue(d interface{}) {
	q.data = append(q.data, d)
}

func (q *queue) dequeue() interface{} {
	l := len(q.data)
	if len(q.data) == 0 {
		return nil
	}

	d := q.data[0]
	if l > 1 {
		q.data = q.data[1:]
	} else {
		q.data = nil
	}
	fmt.Println("dequeued: ", q.data)
	return d
}

func (q *queue) reverse() {
	stk := stack{}

	l := len(q.data)
	for i := 0; i < l; i++ {
		stk.push(q.dequeue())
		defer func() { q.enqueue(stk.pop()) }()
	}

}
