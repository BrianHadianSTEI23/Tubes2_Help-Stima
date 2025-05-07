package model

type Queue struct {
	Val  Tree
	Next *Queue
}

func Enqueue(q *Queue, target *Tree) {
	for q.Next != nil {
		q = q.Next
	}
	// q at the end of the queue
	newQueue := Queue{
		Val:  *target,
		Next: nil,
	}
	q.Next = &newQueue
}

func Dequeue(q *Queue) *Queue {
	var temp *Queue = q
	q = q.Next
	return temp
}

func QueueLength(q *Queue) int64 {
	var count int64 = 0
	for q != nil {
		count++
		q = q.Next
	}
	return count
}
