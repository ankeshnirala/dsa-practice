package queue

import "fmt"

func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() interface{} {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue) Front() interface{} {
	if q.IsEmpty() {
		return nil
	}

	return q.items[0]
}

func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Display() {
	fmt.Println(q)
}
