package queue

type Queue struct {
	items    []int
	capacity int
}

func New(capacity int) Queue {
	return Queue{make([]int, 0, capacity), capacity}
}

func (q *Queue) Append(item int) bool {
	if len(q.items) == q.capacity {
		return false
	}

	q.items = append(q.items, item)
	return true
}

// returns both an int and a boolean
func (q *Queue) Next() (int, bool) {
	if len(q.items) > 0 {
		next := q.items[0]
		// reslice to remove 0th index
		q.items = q.items[1:]
		return next, true
	} else {
		return 0, false
	}
}
